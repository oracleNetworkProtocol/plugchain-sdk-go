package pvm

import (
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	sdk "github.com/oracleNetworkProtocol/plugchain-sdk-go/types"
	"math/big"
)

type Client interface {
	sdk.Module
	TransactionArgs(tran ArgsRequest) ([]byte, error)
	PackData(function_selector string, args ...interface{}) ([]byte, error)

	GetGasPrice() (*big.Int, error)
	EstimateGas(tran ArgsRequest) (uint64, error)
	GetBalance(token, addr string) (*big.Int, error)
	GetCall(token, _func string, parameter ...interface{}) ([]byte, error)
	GetTokenInfo(token string, f ...string) (TokenInfoResponse, error)
	GetTxByHash(hash string) (sdk.PvmResultQueryTx, error)
	GetBlockByNumber(blockId int64, fullTx bool) (map[string]interface{}, error)
	GetTransactionLogs(hash string) ([]*PvmLog, error)
	GetTransactionTxAndLogs(hash string) (PvmTxAndLogs, error)

	Sign(tran ArgsRequest, baseTx sdk.BaseTx) (*ethtypes.Transaction, error)
	Send(tran ArgsRequest, baseTx sdk.BaseTx) (string, string, error)
}

type TokenInfoResponse struct {
	Name        string
	Symbol      string
	Decimals    int64
	TotalSupply *big.Int
}

type ArgsRequest struct {
	From             string
	Token            string
	Gas              uint64
	GasPrice         int64
	Sequence         uint64
	Num              big.Int
	Memo             string
	FunctionSelector string
	Args             []interface{}
}

type PvmTxAndLogs struct {
	PvmResultQueryTx sdk.PvmResultQueryTx
	PvmLogs          []*PvmLog
	Status           bool
	Failed           string
}
