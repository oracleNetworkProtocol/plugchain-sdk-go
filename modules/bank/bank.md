# PLUGCHAIN SDK GO

## BANK MODULE

- [Query](#query)
  - [QueryAccount](#account) --Account Amount
  - [TotalSupply](#supply) --TotalSupply
- [TX](#tx)
  - [Send](#send) --Transfer
  - [MsgSend](#msgsend) --MsgSend
  - [MultiSendRequest](#multisendrequest) --MultiSendRequest
  - [MultiToMultiSend](#multitomultisend) --MultiToMultiSend

# realization

## Query<a name="query"></a><br/>

#### QueryAccount<a name="account"></a><br/>
>QueryAccount return account information specified address
```go
clientCtx := types.Context{
Client:            client.BaseClient,
InterfaceRegistry: client.EncodingConfig().InterfaceRegistry,
}
balance, err := client.Bank.QueryAccount("gx1h3vghuc356mah5swu9fy27fzrr20qe3pqfp6rw", clientCtx)
plug:=balance.Coins.AmountOf("uplugcn")
```

#### TotalSupply<a name="supply"></a><br/>
>TotalSupply queries the total supply of all coins.
```go
supply, err := client.Bank.TotalSupply()
```

## TX<a name="tx"></a><br/>

#### Send<a name="send"></a><br/>
>Send is responsible for transferring tokens from `From` to `to` account

**You need to import the private key before you can operate，Please see the key package for importing the private key**

```go
baseTx := types.BaseTx{
    From:     "demo", //Account name 
    Password: "123123123",
    Gas:      200000,
    Mode:     types.Commit,
    Memo:     "test",
}
baseTx.Fee, err = types.ParseDecCoins("2000uplugcn") //Fee
coins, err := types.ParseDecCoins("100000uplugcn")   //Transfer out quantity + currency name, for example:100000plug
to := "gx1akqhezuftdcc0eqzkq5peqpjlucgmyr7srx54j" //Receiving address
result, err := client.Bank.Send(to, coins, baseTx)
```


#### MsgSend<a name="msgsend"></a><br/>
>get TxHash before sending transactions

**You need to import the private key before you can operate，Please see the key package for importing the private key**
```go
baseTx := types.BaseTx{
From:     "demo", //Account name 
Password: "123123123",
Gas:      200000,
Mode:     types.Commit,
Memo:     "test",
}
coins, err := types.ParseCoins("100000uplugcn")
from := "gx1yhf7w0sq8yn6gqre2pulnqwyy30tjfc4v08f3x"
to := "gx1akqhezuftdcc0eqzkq5peqpjlucgmyr7srx54j"
msg := &bank.MsgSend{
FromAddress: from,
ToAddress:   to,
Amount:      coins,
}
txhash, err := client.BuildTxHash([]types.Msg{msg}, baseTx)
```


#### MultiSendRequest<a name="multisendrequest"></a><br/>
>One to many transaction

**You need to import the private key before you can operate，Please see the key package for importing the private key**
```go
baseTx := types.BaseTx{
    From:     "demo", //Account name 
    Password: "123123123",
    Gas:      200000,
    Mode:     types.Commit,
    Memo:     "test",
}
coins, err := types.ParseCoins("100000uplugcn")

coinss, err := types.ParseDecCoins("1000uplugcn")
mmm := bank.MultiSendRequest{[]bank.Receipt{
{Address: "gx1m6jpgfptpwsxmq7w88z30tsxp22ncwe4fakmnh", Amount:  coinss},
{Address: "gx1jfdv0jrmjr8pye40nwc7hd79k5j4xp2hrnfq34", Amount:  coinss},
}}
results, err := client.Bank.MultiSend(mmm, baseTx)
```


#### MultiToMultiSend<a name="multitomultisend"></a><br/>
>Many to many transaction

**You need to import the private key before you can operate，Please see the key package for importing the private key**
```go
    baseTx := types.BaseTx{
		Gas:  200000,
		Mode: types.Commit,
		Memo: "test",
	}
    //The first account number of the fee deduction address
	baseTx.Fee, err = types.ParseDecCoins("2000uplugcn")
	
    coins, err := types.ParseDecCoins("100uplugcn")
    param := bank.MultiToMultiSendRequest{Receipts: []bank.MutilReceipt{
        {From: "test1", Password: "123123123", To: "gx1m6jpgfptpwsxmq7w88z30tsxp22ncwe4fakmnh", Amount: coins},
        {From: "test2", Password: "123123123", To: "gx1jsw9trgnpdc3e3dwhe69fk4hqct9n8yq8jyzv9", Amount: coins},
        {From: "test3", Password: "123123123", To: "gx18y5dmarzq5wdc5y2h5ydkl9f8ar46gxz75null", Amount: coins},
    }}
    res, err := client.Bank.MultiToMultiSend(param, baseTx)
```