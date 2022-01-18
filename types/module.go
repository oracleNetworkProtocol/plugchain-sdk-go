package types

import (
	"github.com/tendermint/tendermint/crypto"
	"plugchain-sdk-go/codec/types"
)

type Module interface {
	Name() string
	RegisterInterfaceTypes(registry types.InterfaceRegistry)
}

//System type message return
type Response interface {
	Convert() interface{}
}

type SplitAble interface {
	Len() int
	Sub(begin, end int) SplitAble
}

type KeyManager interface {
	Sign(name, password string, data []byte) ([]byte, crypto.PubKey, error)
	Insert(name, password string) (string, string, error)
	Recover(name, password, mnemonic, hdPath string) (string, error)
	Import(name, password string, privKeyArmor string) (address string, err error)
	Export(name, password string) (privKeyArmor string, err error)
	Delete(name, password string) error
	Find(name, password string) (crypto.PubKey, AccAddress, error)
}
