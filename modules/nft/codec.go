package nft

import (
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/codec"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/codec/types"
	cryptocodec "github.com/oracleNetworkProtocol/plugchain-sdk-go/crypto/codec"
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

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgIssueClass{},
		&MsgIssueNFT{},
		&MsgEditNFT{},
		&MsgTransferNFT{},
		&MsgBurnNFT{},
		&MsgTransferClass{},
	)
}
