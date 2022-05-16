package pvm

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/types"
	sdkerrors "github.com/oracleNetworkProtocol/plugchain-sdk-go/types/errors"
	"math/big"
)

var (
	_ TxData = &LegacyTx{}
	_ TxData = &AccessListTx{}
	_ TxData = &DynamicFeeTx{}
)

type TxData interface {
	// TODO: embed ethtypes.TxData. See https://github.com/ethereum/go-ethereum/issues/23154

	TxType() byte
	Copy() TxData
	GetChainID() *big.Int
	GetAccessList() ethtypes.AccessList
	GetData() []byte
	GetNonce() uint64
	GetGas() uint64
	GetGasPrice() *big.Int
	GetGasTipCap() *big.Int
	GetGasFeeCap() *big.Int
	GetValue() *big.Int
	GetTo() *common.Address

	GetRawSignatureValues() (v, r, s *big.Int)
	SetSignatureValues(chainID, v, r, s *big.Int)

	AsEthereumData() ethtypes.TxData
	Validate() error
	Fee() *big.Int
	Cost() *big.Int
}

func (tx *LegacyTx) TxType() byte {
	return ethtypes.LegacyTxType
}

func (tx *LegacyTx) Copy() TxData {
	return &LegacyTx{
		Nonce:    tx.Nonce,
		GasPrice: tx.GasPrice,
		GasLimit: tx.GasLimit,
		To:       tx.To,
		Amount:   tx.Amount,
		Data:     common.CopyBytes(tx.Data),
		V:        common.CopyBytes(tx.V),
		R:        common.CopyBytes(tx.R),
		S:        common.CopyBytes(tx.S),
	}
}

func (tx *LegacyTx) GetChainID() *big.Int {
	//TODO
	//v, _, _ := tx.GetRawSignatureValues()
	//return DeriveChainID(v)
	return big.NewInt(1)
}

func (tx *LegacyTx) GetAccessList() ethtypes.AccessList {
	return nil
}

func (tx *LegacyTx) GetData() []byte {
	return common.CopyBytes(tx.Data)
}

func (tx *LegacyTx) GetNonce() uint64 {
	return tx.Nonce
}

func (tx *LegacyTx) GetGas() uint64 {
	return tx.GasLimit
}

func (tx *LegacyTx) GetGasPrice() *big.Int {
	if tx.GasPrice == nil {
		return nil
	}
	return tx.GasPrice.BigInt()
}

func (tx *LegacyTx) GetGasTipCap() *big.Int {
	return tx.GetGasPrice()
}

func (tx *LegacyTx) GetGasFeeCap() *big.Int {
	return tx.GetGasPrice()
}

func (tx *LegacyTx) GetValue() *big.Int {
	if tx.Amount == nil {
		return nil
	}
	return tx.Amount.BigInt()
}

func (tx *LegacyTx) GetTo() *common.Address {
	if tx.To == "" {
		return nil
	}
	to := common.HexToAddress(tx.To)
	return &to
}

func (tx *LegacyTx) GetRawSignatureValues() (v, r, s *big.Int) {
	return rawSignatureValues(tx.V, tx.R, tx.S)
}

func rawSignatureValues(vBz, rBz, sBz []byte) (v, r, s *big.Int) {
	if len(vBz) > 0 {
		v = new(big.Int).SetBytes(vBz)
	}
	if len(rBz) > 0 {
		r = new(big.Int).SetBytes(rBz)
	}
	if len(sBz) > 0 {
		s = new(big.Int).SetBytes(sBz)
	}
	return v, r, s
}

func (tx *LegacyTx) SetSignatureValues(chainID, v, r, s *big.Int) {
	if v != nil {
		tx.V = v.Bytes()
	}
	if r != nil {
		tx.R = r.Bytes()
	}
	if s != nil {
		tx.S = s.Bytes()
	}
}

func (tx *LegacyTx) AsEthereumData() ethtypes.TxData {
	v, r, s := tx.GetRawSignatureValues()
	return &ethtypes.LegacyTx{
		Nonce:    tx.GetNonce(),
		GasPrice: tx.GetGasPrice(),
		Gas:      tx.GetGas(),
		To:       tx.GetTo(),
		Value:    tx.GetValue(),
		Data:     tx.GetData(),
		V:        v,
		R:        r,
		S:        s,
	}
}

func (tx *LegacyTx) Validate() error {
	gasPrice := tx.GetGasPrice()
	if gasPrice == nil {
		return sdkerrors.Wrap(errors.New("invalid transaction amount"), "gas price cannot be nil")
	}

	if gasPrice.Sign() == -1 {
		return sdkerrors.Wrapf(errors.New("invalid transaction amount"), "gas price cannot be negative %s", gasPrice)
	}
	if !IsValidInt256(gasPrice) {
		return sdkerrors.Wrap(errors.New("invalid transaction amount"), "out of bound")
	}
	if !IsValidInt256(tx.Fee()) {
		return sdkerrors.Wrap(errors.New("invalid transaction amount"), "out of bound")
	}

	amount := tx.GetValue()
	// Amount can be 0
	if amount != nil && amount.Sign() == -1 {
		return sdkerrors.Wrapf(errors.New("invalid transaction amount"), "amount cannot be negative %s", amount)
	}
	if !IsValidInt256(amount) {
		return sdkerrors.Wrap(errors.New("invalid transaction amount"), "out of bound")
	}

	if tx.To != "" {
		if err := types.ValidateAddress(tx.To); err != nil {
			return sdkerrors.Wrap(err, "invalid to address")
		}
	}

	return nil
}

func (tx *LegacyTx) Fee() *big.Int {
	return fee(tx.GetGasPrice(), tx.GetGas())
}

func (tx *LegacyTx) Cost() *big.Int {
	return cost(tx.Fee(), tx.GetValue())
}

func IsValidInt256(i *big.Int) bool {
	return i == nil || i.BitLen() <= 256
}

func fee(gasPrice *big.Int, gas uint64) *big.Int {
	gasLimit := new(big.Int).SetUint64(gas)
	return new(big.Int).Mul(gasPrice, gasLimit)
}

func cost(fee, value *big.Int) *big.Int {
	if value != nil {
		return new(big.Int).Add(fee, value)
	}
	return fee
}

func (m *AccessListTx) TxType() byte {
	//TODO implement me
	panic("implement me")
}

func (m *AccessListTx) Copy() TxData {
	//TODO implement me
	panic("implement me")
}

func (m *AccessListTx) GetChainID() *big.Int {
	//TODO implement me
	panic("implement me")
}

func (m *AccessListTx) GetAccessList() ethtypes.AccessList {
	//TODO implement me
	panic("implement me")
}

func (m *AccessListTx) GetData() []byte {
	//TODO implement me
	panic("implement me")
}

func (m *AccessListTx) GetNonce() uint64 {
	//TODO implement me
	panic("implement me")
}

func (m *AccessListTx) GetGas() uint64 {
	//TODO implement me
	panic("implement me")
}

func (m *AccessListTx) GetGasPrice() *big.Int {
	//TODO implement me
	panic("implement me")
}

func (m *AccessListTx) GetGasTipCap() *big.Int {
	//TODO implement me
	panic("implement me")
}

func (m *AccessListTx) GetGasFeeCap() *big.Int {
	//TODO implement me
	panic("implement me")
}

func (m *AccessListTx) GetValue() *big.Int {
	//TODO implement me
	panic("implement me")
}

func (m *AccessListTx) GetTo() *common.Address {
	//TODO implement me
	panic("implement me")
}

func (m *AccessListTx) GetRawSignatureValues() (v, r, s *big.Int) {
	//TODO implement me
	panic("implement me")
}

func (m *AccessListTx) SetSignatureValues(chainID, v, r, s *big.Int) {
	//TODO implement me
	panic("implement me")
}

func (m *AccessListTx) AsEthereumData() ethtypes.TxData {
	//TODO implement me
	panic("implement me")
}

func (m *AccessListTx) Validate() error {
	//TODO implement me
	panic("implement me")
}

func (m *AccessListTx) Fee() *big.Int {
	//TODO implement me
	panic("implement me")
}

func (m *AccessListTx) Cost() *big.Int {
	//TODO implement me
	panic("implement me")
}

func (m *DynamicFeeTx) TxType() byte {
	//TODO implement me
	panic("implement me")
}

func (m *DynamicFeeTx) Copy() TxData {
	//TODO implement me
	panic("implement me")
}

func (m *DynamicFeeTx) GetChainID() *big.Int {
	//TODO implement me
	panic("implement me")
}

func (m *DynamicFeeTx) GetAccessList() ethtypes.AccessList {
	//TODO implement me
	panic("implement me")
}

func (m *DynamicFeeTx) GetData() []byte {
	//TODO implement me
	panic("implement me")
}

func (m *DynamicFeeTx) GetNonce() uint64 {
	//TODO implement me
	panic("implement me")
}

func (m *DynamicFeeTx) GetGas() uint64 {
	//TODO implement me
	panic("implement me")
}

func (m *DynamicFeeTx) GetGasPrice() *big.Int {
	//TODO implement me
	panic("implement me")
}

func (m *DynamicFeeTx) GetGasTipCap() *big.Int {
	//TODO implement me
	panic("implement me")
}

func (m *DynamicFeeTx) GetGasFeeCap() *big.Int {
	//TODO implement me
	panic("implement me")
}

func (m *DynamicFeeTx) GetValue() *big.Int {
	//TODO implement me
	panic("implement me")
}

func (m *DynamicFeeTx) GetTo() *common.Address {
	//TODO implement me
	panic("implement me")
}

func (m *DynamicFeeTx) GetRawSignatureValues() (v, r, s *big.Int) {
	//TODO implement me
	panic("implement me")
}

func (m *DynamicFeeTx) SetSignatureValues(chainID, v, r, s *big.Int) {
	//TODO implement me
	panic("implement me")
}

func (m *DynamicFeeTx) AsEthereumData() ethtypes.TxData {
	//TODO implement me
	panic("implement me")
}

func (m *DynamicFeeTx) Validate() error {
	//TODO implement me
	panic("implement me")
}

func (m *DynamicFeeTx) Fee() *big.Int {
	//TODO implement me
	panic("implement me")
}

func (m *DynamicFeeTx) Cost() *big.Int {
	//TODO implement me
	panic("implement me")
}
