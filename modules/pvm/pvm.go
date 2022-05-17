package pvm

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/codec"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/codec/types"
	sdk "github.com/oracleNetworkProtocol/plugchain-sdk-go/types"
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

func NewClient(bc sdk.BaseClient, encodingConfig sdk.EncodingConfig, cdc codec.Marshaler) Client {
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
	bz, err := p.TransactionArgs(ArgsRequest{From: addr, Token: token, FunctionSelector: "balanceOf(address)", Args: []interface{}{addr}})
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

func (p pvmClient) GetTxByHash(hash string) (sdk.PvmResultQueryTx, error) {
	builder := sdk.NewEventQueryBuilder()
	builder.AddCondition(sdk.NewCond(TypeMsgEthereumTx, AttributeKeyEthereumTxHash).EQ(hash))
	s, err := p.BaseClient.QueryPvmTxs(builder, nil, nil)
	if err != nil {
		return s, err
	}
	return s, nil
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

//Splicing Data
func (p pvmClient) PackData(function_selector string, args ...interface{}) (data []byte, err error) {
	function_selector = strings.Replace(function_selector, " ", "", -1)
	index := strings.Index(function_selector, "(")
	if index < 1 {
		return data, err
	}
	data = append(data, crypto.Keccak256([]byte(function_selector))[:4]...)
	index += 1
	paramerStr := function_selector[index : len(function_selector)-1]
	paramer := strings.Split(paramerStr, ",")
	if len(paramer) != len(args) {
		return data, err
	}
	_hexParamer, err := p.hexParamer(paramer, false, args)
	if err != nil {
		return data, err
	}
	data = append(data, _hexParamer...)
	return data, err
}

func (p pvmClient) hexParamer(paramer []string, isArr bool, args []interface{}) (data []byte, err error) {
	h := len(args)
	for i := 0; i < h; i++ {
		arg := args[i]
		var v string
		if isArr {
			v = paramer[0]
		} else {
			v = paramer[i]
		}
		if strings.Count(v, "[") != 0 {
			_data, err := p.hexParamer([]string{v[2:]}, true, arg.([]interface{}))
			if err != nil {
				return data, err
			}
			data = append(data, _data...)
		} else {
			switch v {
			case "string":

			case "address":
				_arg := arg.(string)
				if err := sdk.ValidateAccAddress(_arg); err != nil {
					return data, err
				}
				hexAddr, err := sdk.AddressFromAccAddress(_arg)
				if err != nil {
					return data, err
				}
				_data := common.LeftPadBytes(hexAddr.Bytes(), 32)
				data = append(data, _data...)
			case "uint", "uint256", "uint8":

			case "int", "int256", "int8":

			case "bool":

			}
		}

	}
	return data, nil
}
