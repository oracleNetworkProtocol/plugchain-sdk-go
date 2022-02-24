package types

import (
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/codec"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/codec/types"
)

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Marshaler         codec.Marshaler
	TxConfig          TxConfig
	Amino             *codec.LegacyAmino
}
