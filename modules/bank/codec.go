package bank

import (
	"plugchain-sdk-go/codec"
	"plugchain-sdk-go/codec/types"
	cryptocodec "plugchain-sdk-go/crypto/codec"
	"plugchain-sdk-go/modules/auth"
	sdk "plugchain-sdk-go/types"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgSend{},
		&MsgMultiSend{},
	)

	registry.RegisterImplementations(
		(*auth.Account)(nil),
		&auth.BaseAccount{},
		//&auth.EthAccount{},
	)
}
