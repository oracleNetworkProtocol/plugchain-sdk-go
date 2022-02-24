package coinswap

import (
	sdk "github.com/oracleNetworkProtocol/plugchain-sdk-go/types"
)

const (
	ModuleName = "coinswap"

	eventTypeTransfer = "transfer"
	eventTypeSwap     = "swap"

	attributeKeyAmount = "amount"
)

var (
	_ sdk.Msg = &MsgCreatePool{}
	_ sdk.Msg = &MsgDepositWithinBatch{}
	_ sdk.Msg = &MsgSwapWithinBatch{}
	_ sdk.Msg = &MsgWithdrawWithinBatch{}
)

type totalSupply = func() (sdk.Coins, sdk.Error)

// Route implements Msg.
func (msg MsgCreatePool) Route() string { return ModuleName }

// Type implements Msg.
func (msg MsgCreatePool) Type() string { return "add_liquidity" }

// GetSignBytes implements Msg.
func (msg MsgCreatePool) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// ValidateBasic implements Msg.
func (msg MsgCreatePool) ValidateBasic() error {
	if err := msg.DepositCoins.Validate(); err != nil {
		return sdk.Wrapf("invalid DepositCoins: %s", msg.DepositCoins.String())
	}
	if msg.DepositCoins[0] == msg.DepositCoins[1] {
		return sdk.Wrapf("identical MaxToken: %s", msg.DepositCoins.String())
	}
	if _, err := sdk.AccAddressFromBech32(msg.PoolCreatorAddress); err != nil {
		return sdk.Wrap(err)
	}
	return nil
}

// GetSigners implements Msg.
func (msg MsgCreatePool) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.PoolCreatorAddress)}
}

// Route implements Msg.
func (msg MsgDepositWithinBatch) Route() string { return ModuleName }

// Type implements Msg.
func (msg MsgDepositWithinBatch) Type() string { return "Within_liquidity" }

// GetSignBytes implements Msg.
func (msg MsgDepositWithinBatch) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// ValidateBasic implements Msg.
func (msg MsgDepositWithinBatch) ValidateBasic() error {
	if err := msg.DepositCoins.Validate(); err != nil {
		return sdk.Wrapf("invalid DepositCoins: %s", msg.DepositCoins.String())
	}
	if msg.DepositCoins[0] == msg.DepositCoins[1] {
		return sdk.Wrapf("identical MaxToken: %s", msg.DepositCoins.String())
	}
	if !msg.DepositCoins.IsAllPositive() {
		return sdk.Wrapf("IsNotPositive DepositCoins: %s", msg.DepositCoins.String())
	}
	if _, err := sdk.AccAddressFromBech32(msg.DepositorAddress); err != nil {
		return sdk.Wrap(err)
	}
	return nil
}

// GetSigners implements Msg.
func (msg MsgDepositWithinBatch) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.DepositorAddress)}
}

//// Route implements Msg.
//func (msg MsgRemoveLiquidity) Route() string { return ModuleName }
//
//// Type implements Msg.
//func (msg MsgRemoveLiquidity) Type() string { return "remove_liquidity" }
//
//// GetSignBytes implements Msg.
//func (msg MsgRemoveLiquidity) GetSignBytes() []byte {
//	b, err := ModuleCdc.MarshalJSON(&msg)
//	if err != nil {
//		panic(err)
//	}
//	return sdk.MustSortJSON(b)
//}
//
//// ValidateBasic implements Msg.
//func (msg MsgRemoveLiquidity) ValidateBasic() error {
//	if msg.MinToken.IsNegative() {
//		return sdk.Wrapf("minimum token amount can not be negative")
//	}
//	if !msg.WithdrawLiquidity.IsValid() || !msg.WithdrawLiquidity.IsPositive() {
//		return sdk.Wrapf("invalid withdrawLiquidity (%s)", msg.WithdrawLiquidity.String())
//	}
//	if msg.MinStandardAmt.IsNegative() {
//		return sdk.Wrapf("minimum standard token amount %s can not be negative", msg.MinStandardAmt.String())
//	}
//	if msg.Deadline <= 0 {
//		return sdk.Wrapf("deadline %d must be greater than 0", msg.Deadline)
//	}
//	return nil
//}
//
//// GetSigners implements Msg.
//func (msg MsgRemoveLiquidity) GetSigners() []sdk.AccAddress {
//	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Sender)}
//}

// Route implements Msg.
func (msg MsgSwapWithinBatch) Route() string { return ModuleName }

// Type implements Msg.
func (msg MsgSwapWithinBatch) Type() string { return "swap_order" }

// GetSignBytes implements Msg.
func (msg MsgSwapWithinBatch) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// ValidateBasic implements Msg.
func (msg MsgSwapWithinBatch) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.SwapRequesterAddress); err != nil {
		return sdk.Wrapf("invalid pool swap requester address ,err:%s", err)
	}
	if err := msg.OfferCoin.Validate(); err != nil {
		return err
	}
	if !msg.OfferCoin.IsPositive() {
		return sdk.Wrapf("invalid offer coin amount")
	}
	if !msg.OrderPrice.IsPositive() {
		return sdk.Wrapf("invalid order price")
	}
	if !msg.OfferCoin.Amount.GTE(sdk.NewInt(100)) {
		return sdk.Wrapf("offer amount should be over 100 micro")
	}
	return nil
}

// GetSigners implements Msg.
func (msg MsgSwapWithinBatch) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.SwapRequesterAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// Route implements Msg.
func (msg MsgWithdrawWithinBatch) Route() string { return ModuleName }

// Type implements Msg.
func (msg MsgWithdrawWithinBatch) Type() string { return "withdraw_within" }

// GetSignBytes implements Msg.
func (msg MsgWithdrawWithinBatch) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// ValidateBasic implements Msg.
func (msg MsgWithdrawWithinBatch) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.WithdrawerAddress); err != nil {
		return sdk.Wrapf("invalid pool withdrawer address")
	}
	if err := msg.PoolCoin.Validate(); err != nil {
		return err
	}
	if !msg.PoolCoin.IsPositive() {
		return sdk.Wrapf("invalid pool coin amount")
	}
	return nil
}

// GetSigners implements Msg.
func (msg MsgWithdrawWithinBatch) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.WithdrawerAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (m QueryLiquidityPoolResponse) Convert() interface{} {
	return &QueryPoolResponse{
		Pool: _loadPoolInfo(m.Pool),
	}
}

func (m QueryLiquidityPoolsResponse) Convert() interface{} {

	return &QueryAllPoolsResponse{
		Pagination: m.Pagination,
		Pools:      _loadPools(m.Pools),
	}
}

func _loadPoolInfo(info Pool) sdk.PoolInfo {
	return sdk.PoolInfo{
		Id:                    info.Id,
		TypeId:                info.TypeId,
		ReserveCoinDenoms:     info.ReserveCoinDenoms,
		ReserveAccountAddress: info.ReserveAccountAddress,
		PoolCoinDenom:         info.PoolCoinDenom,
	}
}
func _loadPools(pools []Pool) (ret []sdk.PoolInfo) {
	for _, pool := range pools {
		ret = append(ret, _loadPoolInfo(pool))
	}
	return
}
