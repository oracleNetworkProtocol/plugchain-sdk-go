package auth

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/crypto"

	"github.com/oracleNetworkProtocol/plugchain-sdk-go/codec"
	codectypes "github.com/oracleNetworkProtocol/plugchain-sdk-go/codec/types"

	sdk "github.com/oracleNetworkProtocol/plugchain-sdk-go/types"
)

var _ Account = (*EthAccount)(nil)

// GetAddress - Implements sdk.AccountI.
func (acc EthAccount) GetAddress() sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(acc.Address)
	return addr
}

// SetAddress - Implements sdk.AccountI.
func (acc *EthAccount) SetAddress(addr sdk.AccAddress) error {
	if len(acc.Address) != 0 {
		return errors.New("cannot override BaseAccount address")
	}

	acc.Address = addr.String()
	return nil
}

// GetPubKey - Implements sdk.AccountI.
func (acc EthAccount) GetPubKey() (pk crypto.PubKey) {
	if acc.PubKey == nil {
		return nil
	}
	content, ok := acc.PubKey.GetCachedValue().(crypto.PubKey)
	if !ok {
		return nil
	}
	return content
}

// SetPubKey - Implements sdk.AccountI.
func (acc *EthAccount) SetPubKey(pubKey crypto.PubKey) error {
	if pubKey == nil {
		acc.PubKey = nil
	} else {
		protoMsg, ok := pubKey.(proto.Message)
		if !ok {
			return sdk.Wrap(fmt.Errorf("err invalid key, can't proto encode %T", protoMsg))
		}

		any, err := codectypes.NewAnyWithValue(protoMsg)
		if err != nil {
			return err
		}

		acc.PubKey = any
	}

	return nil
}

// GetAccountNumber - Implements AccountI
func (acc EthAccount) GetAccountNumber() uint64 {
	return acc.AccountNumber
}

// SetAccountNumber - Implements AccountI
func (acc *EthAccount) SetAccountNumber(accNumber uint64) error {
	acc.AccountNumber = accNumber
	return nil
}

// GetSequence - Implements sdk.AccountI.
func (acc EthAccount) GetSequence() uint64 {
	return acc.Sequence
}

// SetSequence - Implements sdk.AccountI.
func (acc *EthAccount) SetSequence(seq uint64) error {
	acc.Sequence = seq
	return nil
}

func (acc EthAccount) String() string {
	out, _ := json.Marshal(acc)
	return string(out)
}

// Convert return a sdk.BaseAccount
func (acc *EthAccount) Convert() interface{} {
	// error don't use it
	return nil
}

// Convert return a sdk.BaseAccount
// in order to unpack pubKey so not use Convert()
func (acc *EthAccount) ConvertAccount(cdc codec.Marshaler) interface{} {
	account := sdk.BaseAccount{
		Address:       acc.Address,
		AccountNumber: acc.AccountNumber,
		Sequence:      acc.Sequence,
	}

	var pkStr string
	if acc.PubKey == nil {
		return account
	}

	var pk crypto.PubKey
	if err := cdc.UnpackAny(acc.PubKey, &pk); err != nil {
		return sdk.BaseAccount{}
	}

	pkStr, err := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeAccPub, pk)
	if err != nil {
		panic(err)
	}

	account.PubKey = pkStr
	return account
}
