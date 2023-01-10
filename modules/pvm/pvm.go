package pvm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/codec"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/codec/types"
	sdk "github.com/oracleNetworkProtocol/plugchain-sdk-go/types"
	types2 "github.com/oracleNetworkProtocol/plugchain-sdk-go/types/feemarket/type"
	"math/big"
	"strings"
)

const (
	NAME        = "name()"
	SYMBOL      = "symbol()"
	DECIMALS    = "decimals()"
	TOTALSUPPLY = "totalSupply()"

	TypeMsgEthereumTx          = "ethereum_tx"
	AttributeKeyEthereumTxHash = "ethereumTxHash"
)

type pvmClient struct {
	sdk.BaseClient
	codec.Marshaler
}

func NewClient(bc sdk.BaseClient, cdc codec.Marshaler) Client {
	return pvmClient{
		BaseClient: bc,
		Marshaler:  cdc,
	}
}

func (p pvmClient) Name() string {
	return ModuleName
}

func (p pvmClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (p pvmClient) GetBalance(token, addr string) (*big.Int, error) {
	conn, err := p.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return big.NewInt(0), sdk.Wrap(err)
	}
	if err := sdk.ValidateAccAddressAll(token, addr); err != nil {
		return big.NewInt(0), sdk.Wrap(err)
	}
	hexAddr, err := sdk.AddressFromAccAddress(addr)
	if err != nil {
		return big.NewInt(0), sdk.Wrap(err)
	}
	bz, err := p.TransactionArgs(ArgsRequest{From: addr, Token: token, FunctionSelector: "balanceOf(address)", Args: []interface{}{hexAddr}})
	if err != nil {
		return big.NewInt(0), sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).EthCall(context.Background(), &EthCallRequest{
		Args: bz,
	})
	if err != nil {
		return big.NewInt(0), sdk.Wrap(err)
	}
	return common.BytesToHash(res.Ret).Big(), nil
}

func (p pvmClient) GetCall(token, _func string, parameter ...interface{}) ([]byte, error) {
	conn, err := p.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	bz, err := p.TransactionArgs(ArgsRequest{Token: token, FunctionSelector: _func, Args: parameter})
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).EthCall(context.Background(), &EthCallRequest{
		Args: bz,
	})
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	if res.VmError != "" {
		return nil, sdk.Wrap(errors.New(res.VmError))
	}
	return res.Ret, nil
}

// GasPrice returns the current gas price based on Ethermint's gas price oracle.
func (p pvmClient) GetGasPrice() (*big.Int, error) {
	conn, err := p.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	block, err := p.BaseClient.Block(context.Background(), nil)
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	tipcap := new(big.Int).SetInt64(sdk.DefaultGasPrice)
	baseFee, _ := p.BaseFee(block.Block.Height)
	if baseFee != nil {
		tipcap.Add(tipcap, baseFee)
	}
	return tipcap, nil
}

// EstimateGas returns an estimate of gas usage for the given smart contract call.
func (p pvmClient) EstimateGas(tran ArgsRequest) (uint64, error) {
	conn, err := p.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return 0, sdk.Wrap(err)
	}
	bz, err := p.TransactionArgs(tran)
	if err != nil {
		return 0, sdk.Wrap(err)
	}
	res, err := NewQueryClient(conn).EstimateGas(context.Background(), &EthCallRequest{
		Args:   bz,
		GasCap: 25000000,
	})
	if err != nil {
		return 0, sdk.Wrap(err)
	}
	return res.Gas, nil
}

//Query token basic information
//GetTokenInfo(token,NAME,SYMBOL,DECIMALS,TOTALSUPPLY)
func (p pvmClient) GetTokenInfo(token string, f ...string) (tir TokenInfoResponse, err error) {
	if err := sdk.ValidateAccAddress(token); err != nil {
		return tir, sdk.Wrap(err)
	}
	conn, err := p.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return tir, sdk.Wrap(err)
	}
	for _, _func := range f {
		bz, err := p.TransactionArgs(ArgsRequest{Token: token, FunctionSelector: _func})
		if err != nil {
			return tir, sdk.Wrap(err)
		}

		res, err := NewQueryClient(conn).EthCall(context.Background(), &EthCallRequest{
			Args: bz,
		})
		if err != nil || res.VmError != "" || res.Ret == nil {
			continue
		}
		switch _func {
		case NAME:
			name := res.Ret[64 : 64+res.Ret[63]]
			tir.Name = string(name)
		case SYMBOL:
			symbol := res.Ret[64 : 64+res.Ret[63]]
			tir.Symbol = string(symbol)
		case DECIMALS:
			tir.Decimals = common.BytesToHash(res.Ret).Big().Int64()
		case TOTALSUPPLY:
			tir.TotalSupply = common.BytesToHash(res.Ret).Big()
		}
	}
	return tir, nil
}

func (p pvmClient) GetTxByHash(hash string) (tx sdk.PvmResultQueryTx, err error) {
	builder := sdk.NewEventQueryBuilder()
	builder.AddCondition(sdk.NewCond(TypeMsgEthereumTx, AttributeKeyEthereumTxHash).EQ(hash))
	res, err := p.TxSearchHandle(builder, nil, nil)
	if err != nil {
		return tx, err
	}
	txs, err := p.BaseClient.QueryPvmTxs(res)
	if err != nil {
		return tx, err
	}
	for _, v := range txs {
		if v.Hash.String() == hash {
			tx = v
			break
		}
	}
	return tx, nil
}

func (p pvmClient) GetBlockByNumber(blockId int64, fullTx bool) (map[string]interface{}, error) {
	resBlock, err := p.Block(context.Background(), &blockId)
	if err != nil {
		return nil, err
	}
	if resBlock == nil || resBlock.Block == nil {
		return nil, err
	}
	res, err := p.BaseClient.PvmBlockFromTendermint(resBlock.Block, fullTx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p pvmClient) GetTransactionTxAndLogs(hash string) (tx PvmTxAndLogs, err error) {
	builder := sdk.NewEventQueryBuilder()
	builder.AddCondition(sdk.NewCond(TypeMsgEthereumTx, AttributeKeyEthereumTxHash).EQ(hash))
	res, err := p.TxSearchHandle(builder, nil, nil)
	if err != nil {
		return tx, err
	}
	txs, err := p.BaseClient.QueryPvmTxs(res)
	if err != nil {
		return tx, err
	}
	for _, v := range txs {
		if v.Hash.String() == hash {
			tx.PvmResultQueryTx = v
			break
		}
	}
	if len(res.Txs) == 0 {
		return tx, nil
	}
	logTx := res.Txs[0]
	if logTx == nil {
		return tx, nil
	}
	tx.Status = !strings.Contains(logTx.TxResult.GetLog(), sdk.AttributeKeyEthereumTxFailed)
	if !tx.Status {
		tx.Failed = getLogFailedInfo(logTx.TxResult.GetLog())
	}
	tx.PvmLogs, err = TxLogsFromEvents(logTx.TxResult.Events)
	return tx, nil
}

func (p pvmClient) GetTransactionLogs(hash string) ([]*PvmLog, error) {
	builder := sdk.NewEventQueryBuilder()
	builder.AddCondition(sdk.NewCond(TypeMsgEthereumTx, AttributeKeyEthereumTxHash).EQ(hash))
	res, err := p.TxSearchHandle(builder, nil, nil)
	if err != nil {
		return nil, err
	}
	if len(res.Txs) == 0 {
		return nil, nil
	}
	tx := res.Txs[0]
	if tx == nil {
		return nil, nil
	}
	return TxLogsFromEvents(tx.TxResult.Events)
}

//----------------------------- Transfer -----------------------------

func (p pvmClient) Sign(tran ArgsRequest, baseTx sdk.BaseTx) (*ethtypes.Transaction, error) {
	bz, err := p.PackData(tran.FunctionSelector, tran.Args...)
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	if tran.Sequence == 0 {
		from, _ := p.QueryAddress(baseTx.From, baseTx.Password)
		tran.From = from.String()
		account, err := p.QueryAccount(from.String())
		if err != nil {
			return nil, err
		}
		tran.Sequence = account.Sequence
	}
	if tran.GasPrice == 0 {
		gasPrice, _ := p.GetGasPrice()
		tran.GasPrice = gasPrice.Int64()
	}
	if tran.Gas == 0 {
		tran.Gas, _ = p.EstimateGas(tran)
	}
	legacyTx := &ethtypes.LegacyTx{
		Nonce:    tran.Sequence,
		Value:    &tran.Num,
		Gas:      tran.Gas,
		GasPrice: new(big.Int).SetInt64(tran.GasPrice),
		Data:     bz,
	}
	if tran.Token != "" {
		hexAddr, addErr := sdk.AddressFromAccAddress(tran.Token)
		if addErr != nil {
			return nil, sdk.Wrap(addErr)
		}
		legacyTx.To = &hexAddr
	}
	tx := ethtypes.NewTx(legacyTx)
	_privateKey, err := p.ExportEthsecp256k1(baseTx.From, baseTx.Password)
	privateKeyByte, err := hexutil.Decode("0x" + _privateKey)
	if err != nil {
		return nil, err
	}
	privateKey, err := crypto.ToECDSA(privateKeyByte)
	if err != nil {
		return nil, err
	}
	chainID, err := sdk.ParseChainID(p.BaseClient.QueryChainID())
	if err != nil {
		return nil, err
	}
	tx, err = ethtypes.SignTx(tx, ethtypes.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

//return pvmHash
func (p pvmClient) Send(tran ArgsRequest, baseTx sdk.BaseTx) (pvmHash string, cosmosHash string, err error) {
	tx, err := p.Sign(tran, baseTx)
	if err != nil {
		return "", "", err
	}
	ethereumTx := &MsgEthereumTx{}
	if err := ethereumTx.FromEthereumTx(tx); err != nil {
		return "", "", err
	}
	if err := ethereumTx.ValidateBasic(); err != nil {
		return "", "", err
	}
	txData, err := UnpackTxData(ethereumTx.Data)
	if err != nil {
		return "", "", err
	}
	feeamt := GetFeeAmt(txData)
	fees := sdk.NewDecCoin(sdk.BaseDenom, sdk.NewIntFromBigInt(feeamt))
	baseTx.Fee = sdk.DecCoins{fees}
	baseTx.Gas = txData.GetGas()
	baseTx.Mode = sdk.Sync

	ctx, err := p.BuildPvmAndSend(ethereumTx, baseTx)
	if err != nil {
		return "", "", err
	}
	pvmHash = ethereumTx.AsTransaction().Hash().String()
	cosmosHash = ctx.Hash
	return pvmHash, cosmosHash, nil
}

//Assembly data
func (p pvmClient) TransactionArgs(tran ArgsRequest) ([]byte, error) {
	args := TransactionArgs{}
	if tran.From != "" {
		hexAddr, addErr := sdk.AddressFromAccAddress(tran.From)
		if addErr != nil {
			return nil, sdk.Wrap(addErr)
		}
		args.From = &hexAddr
	}
	if tran.Token != "" {
		hexToken, addErr := sdk.AddressFromAccAddress(tran.Token)
		if addErr != nil {
			return nil, sdk.Wrap(addErr)
		}
		args.To = &hexToken
	}
	if tran.GasPrice > 0 {
		args.GasPrice = (*hexutil.Big)(new(big.Int).SetInt64(tran.GasPrice))
	}
	if tran.Gas > 0 {
		gas := hexutil.Uint64(tran.Gas)
		args.Gas = &gas
	}
	if tran.Num.Int64() > 0 {
		args.Value = (*hexutil.Big)(&tran.Num)
	}
	transferData, err := p.PackData(tran.FunctionSelector, tran.Args...)
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	args.Data = (*hexutil.Bytes)(&transferData)
	bz, err := json.Marshal(&args)
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	return bz, nil
}

// BaseFee returns the base fee tracked by the Fee Market module. If the base fee is not enabled,
// it returns the initial base fee amount. Return nil if London is not activated.
func (p pvmClient) BaseFee(height int64) (*big.Int, error) {
	conn, err := p.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	_params, _ := NewQueryClient(conn).Params(context.Background(), &QueryParamsRequest{})
	var _chainConfig params.ChainConfig
	_chainConfig.LondonBlock = _params.Params.ChainConfig.LondonBlock.BigInt()
	if !_chainConfig.IsLondon(new(big.Int).SetInt64(height)) {
		return nil, nil
	}
	resParams, err := types2.NewQueryClient(conn).Params(context.Background(), &types2.QueryParamsRequest{})
	if err != nil {
		return nil, err
	}
	if resParams.Params.NoBaseFee {
		return nil, nil
	}
	blockRes, err := p.BaseClient.BlockResults(context.Background(), &height)
	if err != nil {
		return nil, err
	}
	basseFee := sdk.BaseFeeFromEvents(blockRes.EndBlockEvents)
	if basseFee != nil {
		return basseFee, nil
	}
	// If we cannot find in events, we tried to get it from the state.
	// It will return feemarket.baseFee if london is activated but feemarket is not enable
	res, err := types2.NewQueryClient(conn).BaseFee(sdk.ContextWithHeight(height), &types2.QueryBaseFeeRequest{})
	if err == nil && res.BaseFee != nil {
		return res.BaseFee.BigInt(), nil
	}
	return nil, nil
}

//Splicing Data
func (p pvmClient) PackData(function_selector string, args ...interface{}) (data []byte, err error) {
	if function_selector == "" && args != nil {
		return args[0].([]byte), nil
	}
	function_selector = strings.Replace(function_selector, " ", "", -1)
	index := strings.Index(function_selector, "(")
	if index < 1 {
		return data, err
	}
	index += 1
	var _hexParamer []byte
	paramerStr := function_selector[index : len(function_selector)-1]
	if strings.Index(paramerStr, "tuple") >= 0 {
		arguments, new_function_selector, err := p.hexParamerTuple(paramerStr)
		if err != nil {
			return data, err
		}
		_hexParamer, err = arguments.Pack(args...)
		if err != nil {
			return data, err
		}
		function_selector = fmt.Sprintf("%v(%v)", function_selector[:index-1], new_function_selector)
	} else {
		arguments, err := p.hexParamer(paramerStr)
		if err != nil {
			return data, err
		}
		_hexParamer, err = arguments.Pack(args...)
		if err != nil {
			return data, err
		}
	}
	data = append(data, crypto.Keccak256([]byte(function_selector))[:4]...)
	data = append(data, _hexParamer...)
	return data, err
}

func (p pvmClient) UnPackData(function_selector string, data []byte) (args interface{}, err error) {
	if function_selector == "" {
		return nil, nil
	}
	function_selector = strings.Replace(function_selector, " ", "", -1)
	index := strings.Index(function_selector, "(")
	if index < 1 {
		return args, err
	}
	index += 1
	paramerStr := function_selector[index : len(function_selector)-1]
	if strings.Index(paramerStr, "tuple") >= 0 {
		arguments, _, err := p.hexParamerTuple(paramerStr)
		if err != nil {
			return args, err
		}
		args, err = arguments.Unpack(data)
		if err != nil {
			return args, err
		}
	} else {
		arguments, err := p.hexParamer(paramerStr)
		if err != nil {
			return args, err
		}
		args, err = arguments.Unpack(data)
		if err != nil {
			return args, err
		}
	}
	return args, err
}

func (p pvmClient) hexParamer(paramerStr string) (arguments abi.Arguments, err error) {
	paramer := strings.Split(paramerStr, ",")
	for _, v := range paramer {
		_type, e := abi.NewType(v, "", nil)
		if e != nil {
			return arguments, e
		}
		arguments = append(arguments, abi.Argument{
			Type: _type,
		})
	}
	return arguments, nil
}

func (p pvmClient) hexParamerTuple(paramerStr string) (arguments abi.Arguments, function_selectors string, err error) {
	paramer := strings.Split(paramerStr, ";")
	for _, v := range paramer {
		function_selector := ""
		var argumentMarshaling []abi.ArgumentMarshaling
		if len(v) > 5 && v[:5] == "tuple" {
			if v[len(v)-2:] == "[]" {
				v = v[:len(v)-2]
				argumentMarshaling, function_selector = cc(v)
				function_selector = fmt.Sprintf("%v[]", function_selector)
				v = "tuple[]"
			} else {
				argumentMarshaling, function_selector = cc(v)
				v = "tuple"
			}
		} else {
			function_selector = v
		}
		function_selectors = fmt.Sprintf("%v%v,", function_selectors, function_selector)
		_type, e := abi.NewType(v, "", argumentMarshaling)
		if e != nil {
			return arguments, function_selectors, e
		}
		arguments = append(arguments, abi.Argument{
			Type: _type,
		})
	}
	function_selectors = function_selectors[:len(function_selectors)-1]
	return arguments, function_selectors, nil
}

func cc(paramerStr string) ([]abi.ArgumentMarshaling, string) {
	var argumentMarshaling []abi.ArgumentMarshaling
	paramerStr = paramerStr[6 : len(paramerStr)-1]
	name := "("
PARAM:
	for {
		if paramerStr == "" {
			break PARAM
		}
		oneParamerStart := strings.Index(paramerStr, ",")
		if oneParamerStart < 0 {
			oneParamerStart = len(paramerStr)
			paramerStr = paramerStr + ","
		}
		paramer := paramerStr[:oneParamerStart]
		openParamerStart := strings.Index(paramer, "tuple(")
		if openParamerStart >= 0 {
			onParamerEnd := strings.LastIndex(paramerStr, ")")
			onParamerEnd = onParamerEnd + strings.Index(paramerStr[onParamerEnd:], ",")
			paramerStrSunType := strings.LastIndex(paramerStr[:onParamerEnd], ":")
			paramerStrSun := paramerStr[:paramerStrSunType]
			isArray := false
			_types := "tuple"
			if paramerStrSun[len(paramerStrSun)-2:] == "[]" {
				paramerStrSun = paramerStrSun[:len(paramerStrSun)-2]
				isArray = true
			}
			argumentMarshalingSun, sunName := cc(paramerStrSun)
			if isArray {
				sunName = fmt.Sprintf("%v[]", sunName)
				_types = "tuple[]"
			}
			name = fmt.Sprintf("%v%v,", name, sunName)
			argumentMarshaling = append(argumentMarshaling, abi.ArgumentMarshaling{
				Name: paramerStr[paramerStrSunType+1 : onParamerEnd], //paramerStrSun[paramerDataLen+1:],
				Type: _types,
				//InternalType: "struct main.AdditionalRecipient.F",
				Components: argumentMarshalingSun,
			})
			paramerStr = paramerStr[onParamerEnd+1:]
			continue PARAM
		}
		paramerStr = paramerStr[oneParamerStart+1:]
		openParamerEnd := strings.Index(paramer, ")")
		if openParamerEnd >= 0 {
			if openParamerEnd == len(paramer)-1 {
				paramerStr = ""
			}
			paramer = paramer[:openParamerEnd]
		}
		paramerData := strings.Split(paramer, ":")
		if len(paramerData) != 2 {
			return nil, ""
		}
		name = fmt.Sprintf("%v%v,", name, paramerData[0])
		argumentMarshaling = append(argumentMarshaling, abi.ArgumentMarshaling{
			Name: paramerData[1],
			//InternalType: "struct main.OrderComponents.F",
			Type: paramerData[0],
		})
	}
	name = fmt.Sprintf("%v)", name[:len(name)-1])
	return argumentMarshaling, name
}

func getLogFailedInfo(log string) (err string) {
	var logTxLog []LogTxLog
	_ = json.Unmarshal([]byte(log), &logTxLog)
	for _, v := range logTxLog[0].Events {
		if v.Type == TypeMsgEthereumTx {
			for _, val := range v.Attributes {
				if val.Key == sdk.AttributeKeyEthereumTxFailed {
					err = val.Value
				}
			}
		}
	}
	return err
}

type LogTxLog struct {
	Events []struct {
		Type       string `json:"type"`
		Attributes []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"attributes"`
	} `json:"events"`
}
