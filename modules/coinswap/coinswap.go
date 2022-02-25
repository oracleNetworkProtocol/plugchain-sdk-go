package coinswap

import (
	"context"
	"fmt"
	cdctypes "github.com/oracleNetworkProtocol/plugchain-sdk-go/codec/types"
	"strings"

	"github.com/oracleNetworkProtocol/plugchain-sdk-go/types/query"

	"github.com/oracleNetworkProtocol/plugchain-sdk-go/codec"
	sdk "github.com/oracleNetworkProtocol/plugchain-sdk-go/types"
)

type coinswapClient struct {
	sdk.BaseClient
	codec.Marshaler
	totalSupply
}

func NewClient(bc sdk.BaseClient, cdc codec.Marshaler, queryTotalSupply totalSupply) Client {
	return coinswapClient{
		BaseClient:  bc,
		Marshaler:   cdc,
		totalSupply: queryTotalSupply,
	}
}

func (swap coinswapClient) Name() string {
	return ModuleName
}

func (swap coinswapClient) RegisterInterfaceTypes(registry cdctypes.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (swap coinswapClient) AddLiquidity(request AddLiquidityRequest,
	baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	creator, err := swap.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := &MsgCreatePool{
		PoolCreatorAddress: creator.String(),
		PoolTypeId:         1,
		DepositCoins:       sdk.Coins{request.BaseToken, request.Token},
	}

	res, err := swap.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, err
	}
	return res, nil
}

func (swap coinswapClient) DepositWithinBatch(request DepositWithinBatchRequest,
	baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	creator, err := swap.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := &MsgDepositWithinBatch{
		DepositorAddress: creator.String(),
		PoolId:           request.PoolId,
		DepositCoins:     sdk.Coins{request.BaseToken, request.Token},
	}

	res, err := swap.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, err
	}
	return res, nil
}

//func (swap coinswapClient) RemoveLiquidity(request RemoveLiquidityRequest,
//	baseTx sdk.BaseTx) (*RemoveLiquidityResponse, error) {
//	creator, err := swap.QueryAddress(baseTx.From, baseTx.Password)
//	if err != nil {
//		return nil, sdk.Wrap(err)
//	}
//
//	msg := &MsgRemoveLiquidity{
//		WithdrawLiquidity: request.Liquidity,
//		MinToken:          ctypes.NewInt(request.MinTokenAmt.Int64()),
//		MinStandardAmt:    ctypes.NewInt(request.MinBaseAmt.Int64()),
//		Deadline:          request.Deadline,
//		Sender:            creator.String(),
//	}
//
//	res, err := swap.BuildAndSend([]sdk.Msg{msg}, baseTx)
//	if err != nil {
//		return nil, err
//	}
//
//	var totalCoins = sdk.NewCoins()
//	coinStrs := res.Events.GetValues(eventTypeTransfer, attributeKeyAmount)
//	for _, coinStr := range coinStrs {
//		coins, er := sdk.ParseCoins(coinStr)
//		if er != nil {
//			swap.Logger().Error("Parse coin str failed", "coin", coinStr)
//			continue
//		}
//		totalCoins = totalCoins.Add(coins...)
//	}
//
//	tokenDenom, er := GetTokenDenomFrom(request.Liquidity.Denom)
//	if er != nil {
//		return nil, er
//	}
//
//	response := &RemoveLiquidityResponse{
//		TokenAmt:  totalCoins.AmountOf(tokenDenom),
//		BaseAmt:   totalCoins.AmountOf(sdk.BaseDenom),
//		Liquidity: request.Liquidity,
//		TxHash:    res.Hash,
//	}
//	return response, nil
//}

func (swap coinswapClient) SwapCoin(request SwapCoinRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	creator, err := swap.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, err
	}

	orderPrice, e := sdk.NewDecFromStr(request.OrderPrice)
	if e != nil {
		return sdk.ResultTx{}, e
	}
	swapFeeRate, e := sdk.NewDecFromStr(request.SwapFeeRate)
	if e != nil {
		return sdk.ResultTx{}, e
	}

	msg := &MsgSwapWithinBatch{
		SwapRequesterAddress: creator.String(),
		PoolId:               request.PoolId,
		SwapTypeId:           1,
		OfferCoin:            request.OfferCoin,
		DemandCoinDenom:      request.DemandCoinDenom,
		OfferCoinFee:         sdk.GetOfferCoinFee(request.OfferCoin, swapFeeRate),
		OrderPrice:           orderPrice,
	}

	res, err := swap.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, err
	}
	return res, nil
}

func (swap coinswapClient) WithdrawWithin(request WithdrawWithinRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	creator, err := swap.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, err
	}

	msg := &MsgWithdrawWithinBatch{
		WithdrawerAddress: creator.String(),
		PoolId:            request.PoolId,
		PoolCoin:          request.PoolCoin,
	}
	res, err := swap.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return sdk.ResultTx{}, err
	}
	return res, nil
}

func (swap coinswapClient) WithdrawWithinBatch(request WithdrawWithinBatchRequest) (interface{}, error) {
	conn, err := swap.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	resp, err := NewQueryClient(conn).LiquidityPoolBatch(
		context.Background(),
		&QueryLiquidityPoolBatchRequest{PoolId: request.PoolId},
	)
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	return resp, err
}

//func (swap coinswapClient) BuyTokenWithAutoEstimate(paidTokenDenom string, boughtCoin sdk.Coin,
//	deadline int64,
//	baseTx sdk.BaseTx,
//) (res *SwapCoinResponse, err error) {
//	var amount = sdk.ZeroInt()
//	switch {
//	case paidTokenDenom == sdk.BaseDenom:
//		amount, err = swap.EstimateBaseForBoughtToken(boughtCoin)
//		break
//	case boughtCoin.Denom == sdk.BaseDenom:
//		amount, err = swap.EstimateTokenForBoughtBase(paidTokenDenom, boughtCoin.Amount)
//		break
//	default:
//		amount, err = swap.EstimateTokenForBoughtToken(paidTokenDenom, boughtCoin)
//		break
//	}
//
//	if err != nil {
//		return nil, err
//	}
//
//	req := SwapCoinRequest{
//		Input:      sdk.NewCoin(paidTokenDenom, amount),
//		Output:     boughtCoin,
//		Deadline:   deadline,
//		IsBuyOrder: true,
//	}
//	return swap.SwapCoin(req, baseTx)
//}

//func (swap coinswapClient) SellTokenWithAutoEstimate(gotTokenDenom string, soldCoin sdk.Coin,
//	deadline int64,
//	baseTx sdk.BaseTx,
//) (res *SwapCoinResponse, err error) {
//	var amount = sdk.ZeroInt()
//	switch {
//	case gotTokenDenom == sdk.BaseDenom:
//		amount, err = swap.EstimateBaseForSoldToken(soldCoin)
//		break
//	case soldCoin.Denom == sdk.BaseDenom:
//		amount, err = swap.EstimateTokenForSoldBase(gotTokenDenom, soldCoin.Amount)
//		break
//	default:
//		amount, err = swap.EstimateTokenForSoldToken(gotTokenDenom, soldCoin)
//		break
//	}
//
//	if err != nil {
//		return nil, err
//	}
//
//	req := SwapCoinRequest{
//		Input:      soldCoin,
//		Output:     sdk.NewCoin(gotTokenDenom, amount),
//		Deadline:   deadline,
//		IsBuyOrder: false,
//	}
//	return swap.SwapCoin(req, baseTx)
//}

func (swap coinswapClient) PoolBatchDeposit(request PoolBatchDepositMsg) (interface{}, error) {
	conn, err := swap.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	resp, err := NewQueryClient(conn).PoolBatchDepositMsg(
		context.Background(),
		&QueryPoolBatchDepositMsgRequest{
			PoolId:   request.PoolId,
			MsgIndex: request.MsgIndex,
		},
	)
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	return resp, nil
}

func (swap coinswapClient) QueryPool(PoolId uint64) (*QueryPoolResponse, error) {
	conn, err := swap.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	resp, err := NewQueryClient(conn).LiquidityPool(
		context.Background(),
		&QueryLiquidityPoolRequest{PoolId: PoolId},
	)
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	return resp.Convert().(*QueryPoolResponse), err
}

func (swap coinswapClient) QueryAllPools(req sdk.PageRequest) (interface{}, error) {
	conn, err := swap.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	resp, err := NewQueryClient(conn).LiquidityPools(
		context.Background(),
		&QueryLiquidityPoolsRequest{
			Pagination: &query.PageRequest{
				Key:        req.Key,
				Offset:     req.Offset,
				Limit:      req.Limit,
				CountTotal: req.CountTotal,
			},
		},
	)
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	return resp, err
}

func (swap coinswapClient) QueryDepositFinish(PoolId, MsgIndex uint64) (interface{}, error) {
	conn, err := swap.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	resp, err := NewQueryClient(conn).PoolDepositSuccessMsg(
		context.Background(),
		&QueryPoolDepositSuccessMsgRequest{
			PoolId:   PoolId,
			MsgIndex: MsgIndex,
		},
	)
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	return resp, err
}

func (swap coinswapClient) QueryDepositsFinish(PoolId uint64) (interface{}, error) {
	conn, err := swap.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	resp, err := NewQueryClient(conn).PoolDepositSuccessMsgs(
		context.Background(),
		&QueryPoolDepositSuccessMsgsRequest{
			PoolId: PoolId,
		},
	)
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	return resp, err
}

func (swap coinswapClient) QuerySwapFinish(PoolId, MsgIndex uint64) (interface{}, error) {
	conn, err := swap.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	resp, err := NewQueryClient(conn).PoolSwapSuccessMsg(
		context.Background(),
		&QueryPoolSwapSuccessMsgRequest{
			PoolId:   PoolId,
			MsgIndex: MsgIndex,
		},
	)
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	return resp, err
}

func (swap coinswapClient) QuerySwapsFinish(PoolId uint64) (interface{}, error) {
	conn, err := swap.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	resp, err := NewQueryClient(conn).PoolSwapSuccessMsgs(
		context.Background(),
		&QueryPoolSwapSuccessMsgsRequest{
			PoolId: PoolId,
		},
	)
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	return resp, err
}

func (swap coinswapClient) QueryWithdrawFinish(PoolId, MsgIndex uint64) (interface{}, error) {
	conn, err := swap.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	resp, err := NewQueryClient(conn).PoolWithdrawSuccessMsg(
		context.Background(),
		&QueryPoolWithdrawSuccessMsgRequest{
			PoolId:   PoolId,
			MsgIndex: MsgIndex,
		},
	)
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	return resp, err
}

func (swap coinswapClient) QueryWithdrawsFinish(PoolId uint64) (interface{}, error) {
	conn, err := swap.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	resp, err := NewQueryClient(conn).PoolWithdrawSuccessMsgs(
		context.Background(),
		&QueryPoolWithdrawSuccessMsgsRequest{
			PoolId: PoolId,
		},
	)
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	return resp, err
}

//func (swap coinswapClient) EstimateTokenForSoldBase(tokenDenom string,
//	soldBaseAmt sdk.Int,
//) (sdk.Int, error) {
//	result, err := swap.QueryPool(tokenDenom)
//	if err != nil {
//		return sdk.ZeroInt(), err
//	}
//	fee := sdk.MustNewDecFromStr(result.Pool.Fee)
//	amount := getInputPrice(soldBaseAmt,
//		result.Pool.Standard.Amount, result.Pool.Token.Amount, fee)
//	return amount, nil
//}

//func (swap coinswapClient) EstimateBaseForSoldToken(soldToken sdk.Coin) (sdk.Int, error) {
//	result, err := swap.QueryPool(soldToken.Denom)
//	if err != nil {
//		return sdk.ZeroInt(), err
//	}
//	fee := sdk.MustNewDecFromStr(result.Pool.Fee)
//	amount := getInputPrice(soldToken.Amount,
//		result.Pool.Token.Amount, result.Pool.Standard.Amount, fee)
//	return amount, nil
//}

//func (swap coinswapClient) EstimateTokenForSoldToken(boughtTokenDenom string,
//	soldToken sdk.Coin) (sdk.Int, error) {
//	if boughtTokenDenom == soldToken.Denom {
//		return sdk.ZeroInt(), errors.New("invalid trade")
//	}
//
//	boughtBaseAmt, err := swap.EstimateBaseForSoldToken(soldToken)
//	if err != nil {
//		return sdk.ZeroInt(), err
//	}
//	return swap.EstimateTokenForSoldBase(boughtTokenDenom, boughtBaseAmt)
//}

//func (swap coinswapClient) EstimateTokenForBoughtBase(soldTokenDenom string,
//	exactBoughtBaseAmt sdk.Int) (sdk.Int, error) {
//	result, err := swap.QueryPool(soldTokenDenom)
//	if err != nil {
//		return sdk.ZeroInt(), err
//	}
//	fee := sdk.MustNewDecFromStr(result.Pool.Fee)
//	amount := getOutputPrice(exactBoughtBaseAmt,
//		result.Pool.Token.Amount, result.Pool.Standard.Amount, fee)
//	return amount, nil
//}

//func (swap coinswapClient) EstimateBaseForBoughtToken(boughtToken sdk.Coin) (sdk.Int, error) {
//	result, err := swap.QueryPool(boughtToken.Denom)
//	if err != nil {
//		return sdk.ZeroInt(), err
//	}
//	fee := sdk.MustNewDecFromStr(result.Pool.Fee)
//	amount := getOutputPrice(boughtToken.Amount,
//		result.Pool.Standard.Amount, result.Pool.Token.Amount, fee)
//	return amount, nil
//}

//func (swap coinswapClient) EstimateTokenForBoughtToken(soldTokenDenom string,
//	boughtToken sdk.Coin) (sdk.Int, error) {
//	if soldTokenDenom == boughtToken.Denom {
//		return sdk.ZeroInt(), errors.New("invalid trade")
//	}
//
//	soldBaseAmt, err := swap.EstimateBaseForBoughtToken(boughtToken)
//	if err != nil {
//		return sdk.ZeroInt(), err
//	}
//	return swap.EstimateTokenForBoughtBase(soldTokenDenom, soldBaseAmt)
//}

func GetLiquidityDenomFrom(denom string) (string, error) {
	if denom == sdk.BaseDenom {
		return "", sdk.Wrapf("should not be base denom : %s", denom)
	}
	return fmt.Sprintf("swap%s", denom), nil
}

func GetTokenDenomFrom(liquidityDenom string) (string, error) {
	if !strings.HasPrefix(liquidityDenom, "swap") {
		return "", sdk.Wrapf("wrong liquidity denom : %s", liquidityDenom)
	}
	return strings.TrimPrefix(liquidityDenom, "swap"), nil
}

// getInputPrice returns the amount of coins bought (calculated) given the input amount being sold (exact)
// The fee is included in the input coins being bought
// https://github.com/runtimeverification/verified-smart-contracts/blob/uniswap/uniswap/x-y-k.pdf
func getInputPrice(inputAmt, inputReserve, outputReserve sdk.Int, fee sdk.Dec) sdk.Int {
	deltaFee := sdk.OneDec().Sub(fee)
	inputAmtWithFee := inputAmt.Mul(sdk.NewIntFromBigInt(deltaFee.BigInt()))
	numerator := inputAmtWithFee.Mul(outputReserve)
	denominator := inputReserve.Mul(sdk.NewIntWithDecimal(1, sdk.Precision)).Add(inputAmtWithFee)
	return numerator.Quo(denominator)
}

// getOutputPrice returns the amount of coins sold (calculated) given the output amount being bought (exact)
// The fee is included in the output coins being bought
func getOutputPrice(outputAmt, inputReserve, outputReserve sdk.Int, fee sdk.Dec) sdk.Int {
	deltaFee := sdk.OneDec().Sub(fee)
	numerator := inputReserve.Mul(outputAmt).Mul(sdk.NewIntWithDecimal(1, sdk.Precision))
	denominator := (outputReserve.Sub(outputAmt)).Mul(sdk.NewIntFromBigInt(deltaFee.BigInt()))
	return numerator.Quo(denominator).Add(sdk.OneInt())
}
