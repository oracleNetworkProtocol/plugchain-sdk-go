package coinswap

import (
	"plugchain-sdk-go/codec"
	"plugchain-sdk-go/codec/types"
	cryptocodec "plugchain-sdk-go/crypto/codec"
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
		&MsgCreatePool{},
		&MsgDepositWithinBatch{},
		&MsgSwapWithinBatch{},
		&MsgWithdrawWithinBatch{},
	)
}
