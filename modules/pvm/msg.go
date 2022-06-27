package pvm

import (
	"errors"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	codectypes "github.com/oracleNetworkProtocol/plugchain-sdk-go/codec/types"
	sdk "github.com/oracleNetworkProtocol/plugchain-sdk-go/types"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/types/tx"
	"math/big"
)

var (
	_ sdk.Msg = &MsgEthereumTx{}
	_ sdk.Tx  = &MsgEthereumTx{}

	_ codectypes.UnpackInterfacesMessage = &MsgEthereumTx{}
)

// AsTransaction creates an Ethereum Transaction type from the msg fields
func (msg MsgEthereumTx) AsTransaction() *ethtypes.Transaction {
	txData, err := UnpackTxData(msg.Data)
	if err != nil {
		return nil
	}

	return ethtypes.NewTx(txData.AsEthereumData())
}

// GetMsgs returns a single MsgEthereumTx as an sdk.Msg.
func (msg *MsgEthereumTx) GetMsgs() []sdk.Msg {
	return []sdk.Msg{msg}
}

func (msg MsgEthereumTx) GetGas() uint64 {
	txData, err := UnpackTxData(msg.Data)
	if err != nil {
		return 0
	}
	return txData.GetGas()
}

// UnpackInterfaces implements UnpackInterfacesMesssage.UnpackInterfaces
func (msg MsgEthereumTx) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	return unpacker.UnpackAny(msg.Data, new(TxData))
}

// fromEthereumTx populates the message fields from the given ethereum transaction
func (msg *MsgEthereumTx) FromEthereumTx(tx *ethtypes.Transaction) error {
	txData, err := NewTxDataFromTx(tx)
	if err != nil {
		return err
	}

	anyTxData, err := PackTxData(txData)
	if err != nil {
		return err
	}

	msg.Data = anyTxData
	msg.Size_ = float64(tx.Size())
	msg.Hash = tx.Hash().Hex()
	return nil
}

//---------------------------------------------

func (msg *MsgEthereumTx) BuildTx(b sdk.TxBuilder, evmDenom string) (sdk.Tx, error) {
	builder, ok := b.(tx.ExtensionOptionsTxBuilder)
	if !ok {
		return nil, errors.New("unsupported builder")
	}
	option, err := codectypes.NewAnyWithValue(&ExtensionOptionsEthereumTx{})
	if err != nil {
		return nil, err
	}

	txData, err := UnpackTxData(msg.Data)
	if err != nil {
		return nil, err
	}
	feeamt := GetFeeAmt(txData)
	fees := sdk.Coins{{Denom: evmDenom, Amount: sdk.NewIntFromBigInt(feeamt)}}

	builder.SetExtensionOptions(option)
	err = builder.SetMsgs(msg)
	if err != nil {
		return nil, err
	}
	builder.SetFeeAmount(fees)
	builder.SetGasLimit(msg.GetGas())
	_tx := builder.GetTx()
	return _tx, err
}

func GetFeeAmt(txData TxData) *big.Int {
	effectiveTip := txData.GetGasPrice()
	gasUsed := new(big.Int).SetUint64(txData.GetGas())
	feeInit := new(big.Int).Mul(gasUsed, effectiveTip)
	divisor := big.NewInt(1000)
	feeAmt := new(big.Int).Div(feeInit, divisor)

	feeAmtJudge := new(big.Int).Mul(feeAmt, divisor)
	if feeAmtJudge.Cmp(feeInit) != 0 {
		feeAmt = feeAmt.Add(feeAmt, big.NewInt(1))
	}
	return feeAmt
}
