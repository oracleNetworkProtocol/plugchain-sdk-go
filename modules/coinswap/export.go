package coinswap

import (
	sdk "github.com/oracleNetworkProtocol/plugchain-sdk-go/types"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/types/query"
)

// expose Record module api for user
type Client interface {
	sdk.Module
	AddLiquidity(request AddLiquidityRequest,
		baseTx sdk.BaseTx) (sdk.ResultTx, error)
	DepositWithinBatch(request DepositWithinBatchRequest,
		baseTx sdk.BaseTx) (sdk.ResultTx, error)
	WithdrawWithin(request WithdrawWithinRequest,
		baseTx sdk.BaseTx) (sdk.ResultTx, error)
	//RemoveLiquidity(request RemoveLiquidityRequest,
	//	baseTx sdk.BaseTx) (*RemoveLiquidityResponse, error)
	SwapCoin(request SwapCoinRequest,
		baseTx sdk.BaseTx) (sdk.ResultTx, error)

	WithdrawWithinBatch(request WithdrawWithinBatchRequest) (interface{}, error)
	//BuyTokenWithAutoEstimate(paidTokenDenom string, boughtCoin sdk.Coin,
	//	deadline int64,
	//	baseTx sdk.BaseTx,
	//) (res *SwapCoinResponse, err error)
	//SellTokenWithAutoEstimate(gotTokenDenom string, soldCoin sdk.Coin,
	//	deadline int64,
	//	baseTx sdk.BaseTx,
	//) (res *SwapCoinResponse, err error)

	PoolBatchDeposit(request PoolBatchDepositMsg) (interface{}, error)
	QueryPool(poolId uint64) (*QueryPoolResponse, error)
	QueryAllPools(pageReq sdk.PageRequest) (interface{}, error)
	QueryDepositFinish(poolId, msgIndex uint64) (interface{}, error)
	QueryDepositsFinish(poolId uint64) (interface{}, error)
	QuerySwapFinish(poolId, msgIndex uint64) (interface{}, error)
	QuerySwapsFinish(poolId uint64) (interface{}, error)
	QueryWithdrawFinish(poolId, msgIndex uint64) (interface{}, error)
	QueryWithdrawsFinish(poolId uint64) (interface{}, error)
	//EstimateTokenForSoldBase(tokenDenom string,
	//	soldBase sdk.Int,
	//) (sdk.Int, error)
	//EstimateBaseForSoldToken(soldToken sdk.Coin) (sdk.Int, error)
	//EstimateTokenForSoldToken(boughtTokenDenom string,
	//	soldToken sdk.Coin) (sdk.Int, error)
	//EstimateTokenForBoughtBase(soldTokenDenom string,
	//	boughtBase sdk.Int) (sdk.Int, error)
	//EstimateBaseForBoughtToken(boughtToken sdk.Coin) (sdk.Int, error)
	//EstimateTokenForBoughtToken(soldTokenDenom string,
	//	boughtToken sdk.Coin) (sdk.Int, error)
}

type AddLiquidityRequest struct {
	BaseToken sdk.Coin
	Token     sdk.Coin
}

type DepositWithinBatchRequest struct {
	PoolId    uint64
	BaseToken sdk.Coin
	Token     sdk.Coin
}

type WithdrawWithinRequest struct {
	PoolId   uint64
	PoolCoin sdk.Coin
}

type WithdrawWithinBatchRequest struct {
	PoolId uint64
}

type RemoveLiquidityRequest struct {
	MinTokenAmt sdk.Int
	MinBaseAmt  sdk.Int
	Liquidity   sdk.Coin
	Deadline    int64
}

type RemoveLiquidityResponse struct {
	TokenAmt  sdk.Int
	BaseAmt   sdk.Int
	Liquidity sdk.Coin
	TxHash    string
}

type SwapCoinRequest struct {
	PoolId          uint64
	OfferCoin       sdk.Coin
	DemandCoinDenom string
	SwapFeeRate     string
	OrderPrice      string
}

type PoolBatchDepositMsg struct {
	PoolId   uint64
	MsgIndex uint64
}

type SwapCoinResponse struct {
	InputAmt  sdk.Int
	OutputAmt sdk.Int
	TxHash    string
}

type QueryPoolResponse struct {
	Pool sdk.PoolInfo
}

type QueryAllPoolsResponse struct {
	Pools      []sdk.PoolInfo
	Pagination *query.PageResponse
}
