package pvm

import (
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/types"
)

func newDynamicFeeTx(tx *ethtypes.Transaction) (*DynamicFeeTx, error) {
	txData := &DynamicFeeTx{
		Nonce:    tx.Nonce(),
		Data:     tx.Data(),
		GasLimit: tx.Gas(),
	}
	v, r, s := tx.RawSignatureValues()
	if tx.To() != nil {
		txData.To = tx.To().Hex()
	}

	if tx.Value() != nil {
		amountInt, err := types.SafeNewIntFromBigInt(tx.Value())
		if err != nil {
			return nil, err
		}
		txData.Amount = &amountInt
	}

	if tx.GasFeeCap() != nil {
		gasFeeCapInt, err := types.SafeNewIntFromBigInt(tx.GasFeeCap())
		if err != nil {
			return nil, err
		}
		txData.GasFeeCap = &gasFeeCapInt
	}

	if tx.GasTipCap() != nil {
		gasTipCapInt, err := types.SafeNewIntFromBigInt(tx.GasTipCap())
		if err != nil {
			return nil, err
		}
		txData.GasTipCap = &gasTipCapInt
	}

	if tx.AccessList() != nil {
		al := tx.AccessList()
		txData.Accesses = NewAccessList(&al)
	}
	txData.SetSignatureValues(tx.ChainId(), v, r, s)
	return txData, nil
}

// NewAccessList creates a new protobuf-compatible AccessList from an ethereum
// core AccessList type
func NewAccessList(ethAccessList *ethtypes.AccessList) AccessList {
	if ethAccessList == nil {
		return nil
	}

	al := AccessList{}
	for _, tuple := range *ethAccessList {
		storageKeys := make([]string, len(tuple.StorageKeys))

		for i := range tuple.StorageKeys {
			storageKeys[i] = tuple.StorageKeys[i].String()
		}

		al = append(al, AccessTuple{
			Address:     tuple.Address.String(),
			StorageKeys: storageKeys,
		})
	}

	return al
}

func newAccessListTx(tx *ethtypes.Transaction) (*AccessListTx, error) {
	txData := &AccessListTx{
		Nonce:    tx.Nonce(),
		Data:     tx.Data(),
		GasLimit: tx.Gas(),
	}

	v, r, s := tx.RawSignatureValues()
	if tx.To() != nil {
		txData.To = tx.To().Hex()
	}

	if tx.Value() != nil {
		amountInt, err := types.SafeNewIntFromBigInt(tx.Value())
		if err != nil {
			return nil, err
		}
		txData.Amount = &amountInt
	}

	if tx.GasPrice() != nil {
		gasPriceInt, err := types.SafeNewIntFromBigInt(tx.GasPrice())
		if err != nil {
			return nil, err
		}
		txData.GasPrice = &gasPriceInt
	}

	if tx.AccessList() != nil {
		al := tx.AccessList()
		txData.Accesses = NewAccessList(&al)
	}

	txData.SetSignatureValues(tx.ChainId(), v, r, s)
	return txData, nil
}

func newLegacyTx(tx *ethtypes.Transaction) (*LegacyTx, error) {
	txData := &LegacyTx{
		Nonce:    tx.Nonce(),
		Data:     tx.Data(),
		GasLimit: tx.Gas(),
	}

	v, r, s := tx.RawSignatureValues()
	if tx.To() != nil {
		txData.To = tx.To().Hex()
	}

	if tx.Value() != nil {
		amountInt, err := types.SafeNewIntFromBigInt(tx.Value())
		if err != nil {
			return nil, err
		}
		txData.Amount = &amountInt
	}

	if tx.GasPrice() != nil {
		gasPriceInt, err := types.SafeNewIntFromBigInt(tx.GasPrice())
		if err != nil {
			return nil, err
		}
		txData.GasPrice = &gasPriceInt
	}

	txData.SetSignatureValues(tx.ChainId(), v, r, s)
	return txData, nil
}
