package signing

import "github.com/oracleNetworkProtocol/plugchain-sdk-go/crypto/types"

//Signature data set, single signature data, multi signature data
type SignatureData interface {
	isSignatureData()
}

type SingleSignatureData struct {
	// SignMode represents the SignMode of the signature
	SignMode SignMode

	// SignMode represents the SignMode of the signature
	Signature []byte
}

// MultiSignatureData represents the nested SignatureData of a multisig signature
type MultiSignatureData struct {
	// BitArray is a compact way of indicating which signers from the multisig key
	// have signed
	BitArray *types.CompactBitArray

	// Signatures is the nested SignatureData's for each signer
	Signatures []SignatureData
}

var _, _ SignatureData = &SingleSignatureData{}, &MultiSignatureData{}

func (m *SingleSignatureData) isSignatureData() {}
func (m *MultiSignatureData) isSignatureData()  {}
