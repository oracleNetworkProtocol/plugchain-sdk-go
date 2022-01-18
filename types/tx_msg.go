package types

import (
	"github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/crypto"
)

type (
	//The interface that the transaction message must satisfy
	Msg interface {
		proto.Message
		//Return message type
		Route() string
		//Return message string
		Type() string
		//Validation before query
		ValidateBasic() error
		//Get message - bytes
		GetSignBytes() []byte

		GetSigners() []AccAddress
	}

	//Return fee
	Fee interface {
		GetGas() uint64
		GetAmount() Coins
	}

	//autograph
	Signature interface {
		GetPubKey() crypto.PubKey
		GetSignature() []byte
	}

	//Interface that the transaction must complete
	Tx interface {
		// Get all transactions
		GetMsgs() []Msg

		//Validation before query
		ValidateBasic() error
	}

	FeeTx interface {
		Tx
		GetGas() uint64
		GetFee() Coins
		FeePayer() AccAddress
		FeeGranter() AccAddress
	}

	TxWithMemo interface {
		Tx
		GetMemo() string
	}

	TxWithTimeoutHeight interface {
		Tx

		GetTimeoutHeight() uint64
	}
)

type TxEncoder func(tx Tx) ([]byte, error)

type TxDecoder func(txBytes []byte) (Tx, error)
