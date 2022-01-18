package store

import (
	"github.com/tendermint/tendermint/crypto"
)

// KeyType reflects a human-readable type for key listing.
type KeyType uint

// KeyInfo saves the basic information of the key
type KeyInfo struct {
	Name         string `json:"name"`
	PubKey       []byte `json:"pubkey"`
	PrivKeyArmor string `json:"priv_key_armor"`
	Algo         string `json:"algo"`
}

type KeyDAO interface {
	// Write will use user password to encrypt data and save to file, the file name is user name
	Write(name, password string, store KeyInfo) error

	// Read will read encrypted data from file and decrypt with user password
	Read(name, password string) (KeyInfo, error)

	// Delete will delete user data and use user password to verify permissions
	Delete(name, password string) error

	// Has returns whether the specified user name exists
	Has(name string) bool
}

type Crypto interface {
	Encrypt(data string, password string) (string, error)
	Decrypt(data string, password string) (string, error)
}

// Info is the publicly exposed information about a keypair
type Info interface {
	// Human-readable type for key listing
	GetType() KeyType
	// Name of the key
	GetName() string
	// Public key
	GetPubKey() crypto.PubKey
	// Bip44 Path
	//GetPath() (*hd.BIP44Params, error)
	// Algo
	//GetAlgo() hd.PubKeyType
}
