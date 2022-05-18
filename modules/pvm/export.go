package pvm

import (
	sdk "github.com/oracleNetworkProtocol/plugchain-sdk-go/types"
	"math/big"
)

type Client interface {
	sdk.Module
	TransactionArgs(tran ArgsRequest) ([]byte, error)
	PackData(function_selector string, args ...interface{}) ([]byte, error)

	GetBalance(token, addr string) (*big.Int, error)
	GetTokenInfo(token string, f ...string) (TokenInfoResponse, error)
	GetTxByHash(hash string) (sdk.PvmResultQueryTx, error)
	GetBlockByNumber(blockId int64, fullTx bool) (map[string]interface{}, error)
	GetTransactionLogs(hash string) ([]*PvmLog, error)
	GetTransactionTxAndLogs(hash string) (PvmTxAndLogs, error)
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
	Gas              int
	GasPrice         int
	Sequence         int
	Num              int
	Memo             string
	FunctionSelector string
	Args             []interface{}
}

type PvmTxAndLogs struct {
	PvmResultQueryTx sdk.PvmResultQueryTx
	PvmLogs          []*PvmLog
	Status           bool
}
