package store

import (
	"fmt"
	"path/filepath"

	"github.com/tendermint/tendermint/libs/json"
	tmdb "github.com/tendermint/tm-db"
)

const (
	keyDBName  = "keys"
	infoSuffix = "info"
)

type LevelDBDAO struct {
	db tmdb.DB
	Crypto
}

func NewLevelDB(rootDir string, crypto Crypto) (KeyDAO, error) {
	db, err := tmdb.NewGoLevelDB(keyDBName, filepath.Join(rootDir, "keys"))
	if err != nil {
		return nil, err
	}

	if crypto == nil {
		crypto = AES{}
	}

	levelDB := LevelDBDAO{
		db:     db,
		Crypto: crypto,
	}
	return levelDB, nil
}

//Write a key message and store it locally
func (k LevelDBDAO) Write(name, password string, store KeyInfo) error {
	if k.Has(name) {
		return fmt.Errorf("name %s has exist", name)
	}

	privStr, err := k.Encrypt(store.PrivKeyArmor, password)
	if err != nil {
		return err
	}

	store.PrivKeyArmor = privStr

	bz, err := json.Marshal(store)
	if err != nil {
		return err
	}
	return k.db.SetSync(infoKey(name), bz)
}

//Read key information locally
func (k LevelDBDAO) Read(name, password string) (store KeyInfo, err error) {
	bz, err := k.db.Get(infoKey(name))
	if bz == nil || err != nil {
		return store, err
	}
	if err := json.Unmarshal(bz, &store); err != nil {
		return store, err
	}

	if len(password) > 0 {
		privStr, err := k.Decrypt(store.PrivKeyArmor, password)
		if err != nil {
			return store, err
		}
		store.PrivKeyArmor = privStr
	}
	return
}

//Delete a key message
func (k LevelDBDAO) Delete(name, password string) error {
	_, err := k.Read(name, password)
	if err != nil {
		return err
	}
	return k.db.DeleteSync(infoKey(name))
}

//Query whether key information exists locally
func (k LevelDBDAO) Has(name string) bool {
	existed, err := k.db.Has(infoKey(name))
	if err != nil {
		return false
	}
	return existed
}

func infoKey(name string) []byte {
	return []byte(fmt.Sprintf("%s.%s", name, infoSuffix))
}
