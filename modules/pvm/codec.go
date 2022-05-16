package pvm

import (
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/codec"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/codec/types"
	cryptocodec "github.com/oracleNetworkProtocol/plugchain-sdk-go/crypto/codec"
	sdk "github.com/oracleNetworkProtocol/plugchain-sdk-go/types"
	sdkerrors "github.com/oracleNetworkProtocol/plugchain-sdk-go/types/errors"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

type (
	ExtensionOptionsEthereumTxI interface{}
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgEthereumTx{},
	)
	registry.RegisterInterface("ethermint.evm.v1.ExtensionOptionsEthereumTx", (*ExtensionOptionsEthereumTxI)(nil), &ExtensionOptionsEthereumTx{})

	registry.RegisterInterface(
		"ethermint.evm.v1.TxData",
		(*TxData)(nil),
		&DynamicFeeTx{},
		&AccessListTx{},
		&LegacyTx{},
	)
	//registry.RegisterInterface("ethermint.evm.v1.TxData\"", (*TokenInterface)(nil), &Token{})
}

// UnpackTxData unpacks an Any into a TxData. It returns an error if the
// client state can't be unpacked into a TxData.
func UnpackTxData(any *types.Any) (TxData, error) {
	if any == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnpackAny, "protobuf Any message cannot be nil")
	}

	txData, ok := any.GetCachedValue().(TxData)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnpackAny, "cannot unpack Any into TxData %T", any)
	}

	return txData, nil
}
