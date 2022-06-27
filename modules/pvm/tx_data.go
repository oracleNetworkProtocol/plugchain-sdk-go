package pvm

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/gogo/protobuf/proto"
	codectypes "github.com/oracleNetworkProtocol/plugchain-sdk-go/codec/types"
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

func (tx *AccessListTx) TxType() byte {
	return ethtypes.AccessListTxType
}

func (tx *AccessListTx) Copy() TxData {
	return &AccessListTx{
		ChainID:  tx.ChainID,
		Nonce:    tx.Nonce,
		GasPrice: tx.GasPrice,
		GasLimit: tx.GasLimit,
		To:       tx.To,
		Amount:   tx.Amount,
		Data:     common.CopyBytes(tx.Data),
		Accesses: tx.Accesses,
		V:        common.CopyBytes(tx.V),
		R:        common.CopyBytes(tx.R),
		S:        common.CopyBytes(tx.S),
	}
}

func (tx *AccessListTx) GetChainID() *big.Int {
	if tx.ChainID == nil {
		return nil
	}

	return tx.ChainID.BigInt()
}

func (tx *AccessListTx) GetAccessList() ethtypes.AccessList {
	if tx.Accesses == nil {
		return nil
	}
	return *tx.Accesses.ToEthAccessList()
}

func (tx *AccessListTx) GetData() []byte {
	return common.CopyBytes(tx.Data)
}

func (tx *AccessListTx) GetNonce() uint64 {
	return tx.Nonce
}

func (tx *AccessListTx) GetGas() uint64 {
	return tx.GasLimit
}

func (tx *AccessListTx) GetGasPrice() *big.Int {
	if tx.GasPrice == nil {
		return nil
	}
	return tx.GasPrice.BigInt()
}

func (tx *AccessListTx) GetGasTipCap() *big.Int {
	return tx.GetGasPrice()
}

func (tx *AccessListTx) GetGasFeeCap() *big.Int {
	return tx.GetGasPrice()
}

func (tx *AccessListTx) GetValue() *big.Int {
	if tx.Amount == nil {
		return nil
	}

	return tx.Amount.BigInt()
}

func (tx *AccessListTx) GetTo() *common.Address {
	if tx.To == "" {
		return nil
	}
	to := common.HexToAddress(tx.To)
	return &to
}

func (tx *AccessListTx) GetRawSignatureValues() (v, r, s *big.Int) {
	return rawSignatureValues(tx.V, tx.R, tx.S)
}

func (tx *AccessListTx) SetSignatureValues(chainID, v, r, s *big.Int) {
	if v != nil {
		tx.V = v.Bytes()
	}
	if r != nil {
		tx.R = r.Bytes()
	}
	if s != nil {
		tx.S = s.Bytes()
	}
	if chainID != nil {
		chainIDInt := types.NewIntFromBigInt(chainID)
		tx.ChainID = &chainIDInt
	}
}

func (tx *AccessListTx) AsEthereumData() ethtypes.TxData {
	v, r, s := tx.GetRawSignatureValues()
	return &ethtypes.AccessListTx{
		ChainID:    tx.GetChainID(),
		Nonce:      tx.GetNonce(),
		GasPrice:   tx.GetGasPrice(),
		Gas:        tx.GetGas(),
		To:         tx.GetTo(),
		Value:      tx.GetValue(),
		Data:       tx.GetData(),
		AccessList: tx.GetAccessList(),
		V:          v,
		R:          r,
		S:          s,
	}
}

func (tx *AccessListTx) Validate() error {
	gasPrice := tx.GetGasPrice()
	if gasPrice == nil {
		return sdkerrors.Wrap(errors.New("invalid gas price"), "cannot be nil")
	}
	if !IsValidInt256(gasPrice) {
		return sdkerrors.Wrap(errors.New("invalid gas price"), "out of bound")
	}

	if gasPrice.Sign() == -1 {
		return sdkerrors.Wrapf(errors.New("invalid gas price"), "gas price cannot be negative %s", gasPrice)
	}

	amount := tx.GetValue()
	// Amount can be 0
	if amount != nil && amount.Sign() == -1 {
		return sdkerrors.Wrapf(errors.New("invalid transaction amount"), "amount cannot be negative %s", amount)
	}
	if !IsValidInt256(amount) {
		return sdkerrors.Wrap(errors.New("invalid transaction amount"), "out of bound")
	}

	if !IsValidInt256(tx.Fee()) {
		return sdkerrors.Wrap(errors.New("invalid gas fee"), "out of bound")
	}

	if tx.To != "" {
		if err := types.ValidateAddress(tx.To); err != nil {
			return sdkerrors.Wrap(err, "invalid to address")
		}
	}

	if tx.GetChainID() == nil {
		return sdkerrors.Wrap(
			sdkerrors.ErrInvalidChainID,
			"chain ID must be present on AccessList txs",
		)
	}

	return nil
}

func (tx *AccessListTx) Fee() *big.Int {
	return fee(tx.GetGasPrice(), tx.GetGas())
}

func (tx *AccessListTx) Cost() *big.Int {
	return cost(tx.Fee(), tx.GetValue())
}

func (tx *DynamicFeeTx) TxType() byte {
	return ethtypes.DynamicFeeTxType
}

func (tx *DynamicFeeTx) Copy() TxData {
	return &DynamicFeeTx{
		ChainID:   tx.ChainID,
		Nonce:     tx.Nonce,
		GasTipCap: tx.GasTipCap,
		GasFeeCap: tx.GasFeeCap,
		GasLimit:  tx.GasLimit,
		To:        tx.To,
		Amount:    tx.Amount,
		Data:      common.CopyBytes(tx.Data),
		Accesses:  tx.Accesses,
		V:         common.CopyBytes(tx.V),
		R:         common.CopyBytes(tx.R),
		S:         common.CopyBytes(tx.S),
	}
}

func (tx *DynamicFeeTx) GetChainID() *big.Int {
	if tx.ChainID == nil {
		return nil
	}

	return tx.ChainID.BigInt()
}

func (tx *DynamicFeeTx) GetAccessList() ethtypes.AccessList {
	if tx.Accesses == nil {
		return nil
	}
	return *tx.Accesses.ToEthAccessList()
}

func (tx *DynamicFeeTx) GetData() []byte {
	return common.CopyBytes(tx.Data)
}

func (tx *DynamicFeeTx) GetNonce() uint64 {
	return tx.Nonce
}

func (tx *DynamicFeeTx) GetGas() uint64 {
	return tx.GasLimit
}

func (tx *DynamicFeeTx) GetGasPrice() *big.Int {
	return tx.GetGasFeeCap()
}

func (tx *DynamicFeeTx) GetGasTipCap() *big.Int {
	if tx.GasTipCap == nil {
		return nil
	}
	return tx.GasTipCap.BigInt()
}

func (tx *DynamicFeeTx) GetGasFeeCap() *big.Int {
	if tx.GasFeeCap == nil {
		return nil
	}
	return tx.GasFeeCap.BigInt()
}

func (tx *DynamicFeeTx) GetValue() *big.Int {
	if tx.Amount == nil {
		return nil
	}

	return tx.Amount.BigInt()
}

func (tx *DynamicFeeTx) GetTo() *common.Address {
	if tx.To == "" {
		return nil
	}
	to := common.HexToAddress(tx.To)
	return &to
}

func (tx *DynamicFeeTx) GetRawSignatureValues() (v, r, s *big.Int) {
	return rawSignatureValues(tx.V, tx.R, tx.S)
}

func (tx *DynamicFeeTx) SetSignatureValues(chainID, v, r, s *big.Int) {
	if v != nil {
		tx.V = v.Bytes()
	}
	if r != nil {
		tx.R = r.Bytes()
	}
	if s != nil {
		tx.S = s.Bytes()
	}
	if chainID != nil {
		chainIDInt := types.NewIntFromBigInt(chainID)
		tx.ChainID = &chainIDInt
	}
}

func (tx *DynamicFeeTx) AsEthereumData() ethtypes.TxData {
	v, r, s := tx.GetRawSignatureValues()
	return &ethtypes.DynamicFeeTx{
		ChainID:    tx.GetChainID(),
		Nonce:      tx.GetNonce(),
		GasTipCap:  tx.GetGasTipCap(),
		GasFeeCap:  tx.GetGasFeeCap(),
		Gas:        tx.GetGas(),
		To:         tx.GetTo(),
		Value:      tx.GetValue(),
		Data:       tx.GetData(),
		AccessList: tx.GetAccessList(),
		V:          v,
		R:          r,
		S:          s,
	}
}

func (tx *DynamicFeeTx) Validate() error {
	if tx.GasTipCap == nil {
		return sdkerrors.Wrap(errors.New("invalid gas cap"), "gas tip cap cannot nil")
	}

	if tx.GasFeeCap == nil {
		return sdkerrors.Wrap(errors.New("invalid gas cap"), "gas fee cap cannot nil")
	}

	if tx.GasTipCap.IsNegative() {
		return sdkerrors.Wrapf(errors.New("invalid gas cap"), "gas tip cap cannot be negative %s", tx.GasTipCap)
	}

	if tx.GasFeeCap.IsNegative() {
		return sdkerrors.Wrapf(errors.New("invalid gas cap"), "gas fee cap cannot be negative %s", tx.GasFeeCap)
	}

	if !IsValidInt256(tx.GetGasTipCap()) {
		return sdkerrors.Wrap(errors.New("invalid gas cap"), "out of bound")
	}

	if !IsValidInt256(tx.GetGasFeeCap()) {
		return sdkerrors.Wrap(errors.New("invalid gas cap"), "out of bound")
	}

	if tx.GasFeeCap.LT(*tx.GasTipCap) {
		return sdkerrors.Wrapf(
			errors.New("invalid gas cap"), "max priority fee per gas higher than max fee per gas (%s > %s)",
			tx.GasTipCap, tx.GasFeeCap,
		)
	}

	if !IsValidInt256(tx.Fee()) {
		return sdkerrors.Wrap(errors.New("invalid gas fee"), "out of bound")
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

	if tx.GetChainID() == nil {
		return sdkerrors.Wrap(
			sdkerrors.ErrInvalidChainID,
			"chain ID must be present on AccessList txs",
		)
	}

	return nil
}

func (tx *DynamicFeeTx) Fee() *big.Int {
	return fee(tx.GetGasPrice(), tx.GasLimit)
}

func (tx *DynamicFeeTx) Cost() *big.Int {
	return cost(tx.Fee(), tx.GetValue())
}

func NewTxDataFromTx(tx *ethtypes.Transaction) (TxData, error) {
	var txData TxData
	var err error
	switch tx.Type() {
	case ethtypes.DynamicFeeTxType:
		txData, err = newDynamicFeeTx(tx)
	case ethtypes.AccessListTxType:
		txData, err = newAccessListTx(tx)
	default:
		txData, err = newLegacyTx(tx)
	}
	if err != nil {
		return nil, err
	}

	return txData, nil
}

func PackTxData(txData TxData) (*codectypes.Any, error) {
	msg, ok := txData.(proto.Message)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrPackAny, "cannot proto marshal %T", txData)
	}

	anyTxData, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPackAny, err.Error())
	}

	return anyTxData, nil
}
