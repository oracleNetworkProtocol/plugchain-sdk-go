package types

import (
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
	"google.golang.org/grpc"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
)

type TxManager interface {
	TmQuery
	BuildTxHash(msg []Msg, baseTx BaseTx) (string, Error)
	BuildAndSend(msg []Msg, baseTx BaseTx) (ResultTx, Error)
	BuildAndSign(msg []Msg, baseTx BaseTx) ([]byte, Error)
	SendBatch(msgs Msgs, baseTx BaseTx) ([]ResultTx, Error)
	BuildAndSendWithAccount(addr string, accountNumber, sequence uint64, msg []Msg, baseTx BaseTx) (ResultTx, Error)
	//BuildAndSends(txByte []byte, ctxMode BroadcastMode, simulate bool, addr []string) (ResultTx, Error)
	BuildAndSends(msg []Msg, baseTx BaseTx, account map[string]string, addr []string) (ResultTx, Error)
}

type Queries interface {
	StoreQuery
	AccountQuery
	TmQuery
}

type GRPCClient interface {
	GenConn() (*grpc.ClientConn, error)
	GenContext() (*grpc.ClientConn, error)
}

type ParamQuery interface {
	QueryParams(module string, res Response) Error
}

type StoreQuery interface {
	QueryWithResponse(path string, data interface{}, result Response) error
	Query(path string, data interface{}) ([]byte, error)
	QueryStore(key HexBytes, storeName string, height int64, prove bool) (abci.ResponseQuery, error)
}

type AccountQuery interface {
	QueryAccount(address string) (BaseAccount, Error)
	QueryAddress(name, password string) (AccAddress, Error)
}

type TmQuery interface {
	QueryTx(hash string) (ResultQueryTx, error)
	QueryTxs(builder *EventQueryBuilder, page, size *int) (ResultSearchTxs, error)
	QueryPvmTxs(res *ctypes.ResultTxSearch) (PvmResultQueryTx, error)
	PvmBlockFromTendermint(block *tmtypes.Block, fullTx bool) (map[string]interface{}, error)
	QueryBlock(height int64) (BlockDetail, error)
	TxSearchHandle(builder *EventQueryBuilder, page, size *int) (*ctypes.ResultTxSearch, error)
}

type TokenManager interface {
	QueryToken(denom string) (Token, error)
	SaveTokens(tokens ...Token)
}

type TokenConvert interface {
	ToMinCoin(coin ...DecCoin) (Coins, Error)
	ToMainCoin(coin ...Coin) (DecCoins, Error)
}

type Logger interface {
	Logger() log.Logger
	SetLogger(log.Logger)
}

type BaseClient interface {
	TxManager
	TokenManager
	KeyManager
	Queries
	TokenConvert
	TmClient
	Logger
	GRPCClient
}
