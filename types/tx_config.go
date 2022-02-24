package types

import signingtypes "github.com/oracleNetworkProtocol/plugchain-sdk-go/types/tx/signing"

type (
	TxConfig interface {
		TxEncodingConfig

		NewTxBuilder() TxBuilder
		WrapTxBuilder(Tx) (TxBuilder, error)
		SignModeHandler() SignModeHandler
	}

	//Tx Builder The interface type that defines a specific transaction defined by an application must be implemented: it must be able to set messages, generate signatures, and provide specification bytes for signature
	TxBuilder interface {
		GetTx() Tx

		SetMsgs(msgs ...Msg) error
		SetSignatures(signatures ...signingtypes.SignatureV2) error
		SetMemo(memo string)
		SetFeeAmount(amount Coins)
		SetGasLimit(limit uint64)
		SetTimeoutHeight(height uint64)
	}

	//Encoders and decoders containing transactions
	TxEncodingConfig interface {
		TxEncoder() TxEncoder
		TxDecoder() TxDecoder
		TxJSONEncoder() TxEncoder
		TxJSONDecoder() TxDecoder
		MarshalSignatureJSON([]signingtypes.SignatureV2) ([]byte, error)
		UnmarshalSignatureJSON([]byte) ([]signingtypes.SignatureV2, error)
	}
)
