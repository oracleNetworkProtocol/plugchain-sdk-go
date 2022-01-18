package types

import (
	proto "github.com/gogo/protobuf/proto"
	tmcrypto "github.com/tendermint/tendermint/crypto"
)

//Public key
type PubKey interface {
	proto.Message
	tmcrypto.PubKey
}

//Private key
type PrivKey interface {
	proto.Message
	tmcrypto.PrivKey
}

type (
	Address = tmcrypto.Address
)

//Convert public key to tendermint's key
type IntoTmPubKey interface {
	AsTmPubKey() tmcrypto.PubKey
}
