package bank

import (
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/codec"
	types2 "github.com/oracleNetworkProtocol/plugchain-sdk-go/codec/types"
	cryptocodec "github.com/oracleNetworkProtocol/plugchain-sdk-go/crypto/codec"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/modules/auth"
	sdk "github.com/oracleNetworkProtocol/plugchain-sdk-go/types"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

func RegisterInterfaces(registry types2.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgSend{},
		&MsgMultiSend{},
	)

	registry.RegisterImplementations(
		(*auth.Account)(nil),
		&auth.BaseAccount{},
		&auth.EthAccount{},
	)
	//registry.RegisterImplementations(
	//	(*auth.EthAccount)(nil),
	//	&auth.EthAccount{},
	//)
}
