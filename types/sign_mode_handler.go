package types

import (
	"plugchain-sdk-go/types/tx/signing"
)

//By generating sign bytes SignMode from Tx and  SignerData
type SignModeHandler interface {
	//Default mode
	DefaultMode() signing.SignMode

	// Process default list
	Modes() []signing.SignMode

	//Or signature data
	GetSignBytes(mode signing.SignMode, data SignerData, tx Tx) ([]byte, error)
}

// The specific information required to sign the transaction is not included in the transaction subject itself
type SignerData struct {
	// ChainID
	ChainID string

	// Signer's account number
	AccountNumber uint64

	// Sequence is the account sequence number of the signer that is used
	// for replay protection. This field is only useful for Legacy Amino signing,
	// since in SIGN_MODE_DIRECT the account sequence is already in the signer
	// info.
	Sequence uint64
}
