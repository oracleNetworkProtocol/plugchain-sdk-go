# PLUGCHAIN SDK GO

## NFT MODULE

- [Query](#query)
    - [TransactionArgs](#TransactionArgs) --TransactionArgs
    - [PackData](#PackData) --PackData
    - [GetGasPrice](#GetGasPrice) --GetGasPrice
    - [EstimateGas](#EstimateGas) --EstimateGas
    - [GetBalance](#GetBalance) --GetBalance
    - [GetCall](#GetCall) --GetCall
    - [GetTokenInfo](#GetTokenInfo) --GetTokenInfo
    - [GetTxByHash](#GetTxByHash) --GetTxByHash
    - [GetBlockByNumber](#GetBlockByNumber) --GetBlockByNumber
    - [GetTransactionLogs](#GetTransactionLogs) --GetTransactionLogs
    - [GetTransactionTxAndLogs](#GetTransactionTxAndLogs) --GetTransactionTxAndLogs
- [TX](#tx)
    - [Sign](#sign) --Sign
    - [Send](#send) --Send

# prepare

## tran<a name="tran"></a>

> Transaction data set

```go
addr, _ := types.AddressFromAccAddress("gx19juplfvcw0la2av46f6nl9g3hfmwsw4l2vrnu3")
tran := pvm.ArgsRequest{
//Signature address
From:             "gx19juplfvcw0la2av46f6nl9g3hfmwsw4l2vrnu3",
//Contract address
Token:            "gx1eahrjsw0hu2f3ey4d796rvp9074z7qx9mv9r58",
//Contract execution method，Only the parameter type needs to be filled in
FunctionSelector: "transfer(address,uint256)",
//Call method incoming parameters
//The address type must be converted with the AddressFromAccAddress method，
//The uint parameter is converted to big.Int
Args:             []interface{}{addr, new(big.Int).SetInt64(5000000)},
}
```

# realization

## Query<a name="query"></a><br/>

#### TransactionArgs<a name="TransactionArgs"></a><br/>

> Assembly data

```go
//Please check tran data at the top
res, err := client.Pvm.TransactionArgs(tran)
```

#### PackData<a name="PackData"></a><br/>

> Pack performs the operation Go format -> Hexdata.

```go
addr, _ := types.AddressFromAccAddress("gx19juplfvcw0la2av46f6nl9g3hfmwsw4l2vrnu3")
res, err := client.Pvm.PackData("transfer(address,uint256)", addr, new(big.Int).SetInt64(1000000000))
```

#### GetGasPrice<a name="GetGasPrice"></a><br/>

> Get gas price

```go
gas, err := client.Pvm.GetGasPrice()
```

#### EstimateGas<a name="EstimateGas"></a><br/>

> Get gas limit

```go
//Please check tran data at the top
gasLimit, _ := client.Pvm.EstimateGas(tran)
```

#### GetBalance<a name="GetBalance"></a><br/>

> Get Address Token balance

```go
balance, err := client.Pvm.GetBalance("gx1eahrjsw0hu2f3ey4d796rvp9074z7qx9mv9r58", "gx19juplfvcw0la2av46f6nl9g3hfmwsw4l2vrnu3")
```

#### GetCall<a name="GetCall"></a><br/>

> Contract reading method，Parameter type view top tran parameter remarks

```go
addr, _ := types.AddressFromAccAddress("gx19juplfvcw0la2av46f6nl9g3hfmwsw4l2vrnu3")
res, err := client.Pvm.GetCall("gx1eahrjsw0hu2f3ey4d796rvp9074z7qx9mv9r58", "balanceOf(address)", addr)
```

#### GetTokenInfo<a name="GetTokenInfo"></a><br/>

> Obtain basic contract information

```go
//Get data optional
tokenInfo, err := client.Pvm.GetTokenInfo("gx1eahrjsw0hu2f3ey4d796rvp9074z7qx9mv9r58", pvm.NAME, pvm.SYMBOL, pvm.TOTALSUPPLY, pvm.DECIMALS)
```

#### GetTxByHash<a name="GetTxByHash"></a><br/>

> Obtain transaction information according to hash

```go
txData, _ := client.Pvm.GetTxByHash("0x8913ab96c12a52f13021003a2de683f6d6e5c6548c35d43eeef1c7335c72b46d")
```

#### GetBlockByNumber<a name="GetBlockByNumber"></a><br/>

> Go back to the transaction information according to the height of the block

```go
blockTxData, _ := client.Pvm.GetBlockByNumber(4255929, true)
```

#### GetTransactionLogs<a name="GetTransactionLogs"></a><br/>

> Obtain transaction log according to hash

```go
tx, _ := client.Pvm.GetTransactionLogs("0x8913ab96c12a52f13021003a2de683f6d6e5c6548c35d43eeef1c7335c72b46d")
```

#### GetTransactionTxAndLogs<a name="GetTransactionTxAndLogs"></a><br/>

> Obtain transaction information and log information according to hash

```go
tx, _ := client.Pvm.GetTransactionTxAndLogs("0x8913ab96c12a52f13021003a2de683f6d6e5c6548c35d43eeef1c7335c72b46d")
```

## TX<a name="tx"></a><br/>

#### Sign<a name="Sign"></a><br/>

> PVM transaction signature
**You need to import the private key before you can operate，Please see the key package for importing the private key**

```go
baseTx := types.BaseTx{
From:     "demo", //Account name 
Password: "123123123",
Mode:     types.Commit,
}
//Please check tran data at the top
tx, err := client.Pvm.Sign(tran, baseTxss)
```

#### Send<a name="Send"></a><br/>

> Send PVM transaction
**You need to import the private key before you can operate，Please see the key package for importing the private key**

```go
baseTx := types.BaseTx{
From:     "demo", //Account name 
Password: "123123123",
Mode:     types.Commit,
}
//Please check tran data at the top
txhash, cosmoshash, err := client.Pvm.Send(tran, baseTxss)
```
