# PLUGCHAIN SDK GO

## LIQUIDITY MODULE

- [Query](#query)
    - [QueryPool](#QueryPool) --QueryPool
    - [QueryAllPools](#QueryAllPools) --QueryAllPools
    - [QueryDepositFinish](#QueryDepositFinish) --QueryDepositFinish
    - [QueryDepositsFinish](#QueryDepositsFinish) --QueryDepositsFinish
    - [QuerySwapFinish](#QuerySwapFinish) --QuerySwapFinish
    - [QuerySwapsFinish](#QuerySwapsFinish) --QuerySwapsFinish
    - [WithdrawWithinBatch](#WithdrawWithinBatch) --WithdrawWithinBatch
    - [QueryWithdrawFinish](#QueryWithdrawFinish) --QueryWithdrawFinish
    - [QueryWithdrawsFinish](#QueryWithdrawsFinish) --QueryWithdrawsFinish
- [TX](#tx)
    - [AddLiquidity](#AddLiquidity) --AddLiquidity
    - [DepositWithinBatch](#DepositWithinBatch) --DepositWithinBatch
    - [WithdrawWithin](#WithdrawWithin) --WithdrawWithin
    - [SwapCoin](#SwapCoin) --SwapCoin

# realization

## Query<a name="query"></a><br/>

#### QueryPool<a name="QueryPool"></a><br/>

> Query current equity pool

```go
resp, err := client.Swap.QueryPool()
```

#### QueryAllPools<a name="QueryAllPools"></a><br/>

> List all coinswap transaction pairs

```go
resp, err := client.Swap.QueryAllPools(types.PageRequest{
Offset:     0,
Limit:      10,
CountTotal: false,
})
```

#### QueryDepositFinish<a name="QueryDepositFinish"></a><br/>

> Query all completed deposit message of liquidity pool by msg index

```go
resp, err := client.Swap.QueryDepositFinish( 1, 0)
```

#### QueryDepositsFinish<a name="QueryDepositsFinish"></a><br/>

> Query all completed deposit messages of liquidity pool

```go
resp, err := client.Swap.QueryDepositsFinish(1)
```

#### QuerySwapFinish<a name="QuerySwapFinish"></a><br/>

> Query all completed swap message of liquidity pool by msg index

```go
resp, err := client.Swap.QuerySwapFinish( 1, 1)
```

#### QuerySwapsFinish<a name="QuerySwapsFinish"></a><br/>

> Query all completed swap messages of liquidity pool

```go
resp, err := client.Swap.QuerySwapsFinish( 1)
```

#### WithdrawWithinBatch<a name="WithdrawWithinBatch"></a><br/>

> Query the withdraw messages in the liquidity pool batch

```go
request := coinswap.WithdrawWithinBatchRequest{
PoolId: 1,
}
res, err := client.Swap.WithdrawWithinBatch(request)
```

#### QueryWithdrawFinish<a name="QueryWithdrawFinish"></a><br/>

> Query all completed withdraw message of liquidity pool by msg index

```go
resp, err := client.Swap.QueryWithdrawFinish(1, 1)
```

#### QueryWithdrawsFinish<a name="QueryWithdrawsFinish"></a><br/>

> Query all completed withdraw messages of liquidity pool

```go
resp, err := client.Swap.QueryWithdrawsFinish(1)
```

## TX<a name="tx"></a><br/>

#### AddLiquidity<a name="AddLiquidity"></a><br/>

> Create liquidity pool and deposit coins
**You need to import the private key before you can operate，Please see the key package for importing the private key**

```go
baseTx := types.BaseTx{
From:     "demo", //Account name 
Password: "123123123",
Gas:      200000,
Mode:     types.Commit,
Memo:     "test",
}
baseTx.Fee, err = types.ParseDecCoins("2000plug") //Fee
request := coinswap.AddLiquidityRequest{
BaseToken: types.Coin{
Denom:  "plug",
Amount: types.NewInt(5000000),
},
Token: types.Coin{
Denom:  "tt3",
Amount: types.NewInt(1000000),
},
}
res, err := client.Swap.AddLiquidity(request, baseTx)
```

#### DepositWithinBatch<a name="DepositWithinBatch"></a><br/>

> Deposit coins to a liquidity pool
**You need to import the private key before you can operate，Please see the key package for importing the private key**

```go
baseTx := types.BaseTx{
From:     "demo", //Account name 
Password: "123123123",
Gas:      200000,
Mode:     types.Commit,
Memo:     "test",
}
baseTx.Fee, err = types.ParseDecCoins("2000plug") //Fee
request := coinswap.DepositWithinBatchRequest{
BaseToken: types.Coin{
Denom:  "plug",
Amount: types.NewInt(5000),
},
Token: types.Coin{
Denom:  "tt3",
Amount: types.NewInt(1000),
},
PoolId: 1,
}
res, err := client.Swap.DepositWithinBatch(request, baseTx)
```

#### WithdrawWithin<a name="WithdrawWithin"></a><br/>

> Withdraw pool coin from the specified liquidity pool
**You need to import the private key before you can operate，Please see the key package for importing the private key**

```go
baseTx := types.BaseTx{
From:     "demo", //Account name 
Password: "123123123",
Gas:      200000,
Mode:     types.Commit,
Memo:     "test",
}
baseTx.Fee, err = types.ParseDecCoins("2000plug") //Fee
request := coinswap.WithdrawWithinRequest{
PoolId:   1,
PoolCoin: types.Coin{"pool1F9D08EC2F71CFF36479EE3861FCFFB46F3CBB4F3514FF7D17985BE30A708FE5", types.NewInt(1000)},
}
res, err := client.Swap.WithdrawWithin(request, baseTx)
```

#### SwapCoin<a name="SwapCoin"></a><br/>

> Swap offer coin with demand coin from the liquidity pool with the given order price
**You need to import the private key before you can operate，Please see the key package for importing the private key**

```go
baseTx := types.BaseTx{
From:     "demo", //Account name 
Password: "123123123",
Gas:      200000,
Mode:     types.Commit,
Memo:     "test",
}
baseTx.Fee, err = types.ParseDecCoins("2000plug") //Fee
request := coinswap.SwapCoinRequest{
PoolId:          1,
OfferCoin:       types.Coin{"plug", types.NewInt(1000)},
DemandCoinDenom: "tts",
SwapFeeRate:     "0.003",
OrderPrice:      "0.019",
}
res, err := client.Swap.SwapCoin(request, baseTx)
```
