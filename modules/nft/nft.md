# PLUGCHAIN SDK GO

## NFT MODULE

- [Query](#query)
    - [QuerySupply](#supply) --QuerySupply
    - [QueryOwner](#owner) --QueryOwner
    - [QueryCollection](#collection) --QueryCollection
    - [QueryDenom](#denom) --QueryDenom
    - [QueryDenoms](#denoms) --QueryDenoms
    - [QueryNFT](#nft) --QueryNFT
- [TX](#tx)
    - [IssueDenom](#issue) --IssueDenom
    - [MintNFT](#mint) --MintNFT
    - [EditNFT](#edit) --EditNFT
    - [TransferNFT](#transfer_nft) --TransferNFT
    - [TransferClass](#transfer_class) --TransferClass
    - [BurnNFT](#burn) --BurnNFT

# realization

## Query<a name="query"></a><br/>

#### QuerySupply<a name="supply"></a><br/>
>total supply of a collection or owner of NFTs
```go
res, err := client.Nft.QuerySupply("a123")
```

#### QueryOwner<a name="owner"></a><br/>
>Get the NFTs owned by an account addr.
```go
res, err := client.Nft.QueryOwner("gx1akqhezuftdcc0eqzkq5peqpjlucgmyr7srx54j", "a123", types.PageRequest{
	Offset:     0,
	Limit:      10,
	CountTotal: false,
})
```

#### QueryCollection<a name="collection"></a><br/>
>Get all the NFTs from a given collection.
```go
res, err := client.Nft.QueryCollection("a123", types.PageRequest{
	Offset:     0,
	Limit:      10,
	CountTotal: false,
})
```


#### QueryDenom<a name="denom"></a><br/>
>Query the class by the specified class id.
```go
res, err := client.Nft.QueryDenom("a123")
```

#### QueryDenoms<a name="denoms"></a><br/>
>Query all denominations of all collections of NFTs.
```go
res, err := client.Nft.QueryDenoms(types.PageRequest{
	Offset:     0,
	Limit:      10,
	CountTotal: false,
})
```

#### QueryNFT<a name="nft"></a><br/>
>Query a single NFT from a collection
```go
res, err := client.Nft.QueryNFT("a123", "a1232")
```



## TX<a name="tx"></a><br/>

#### IssueDenom<a name="issue"></a><br/>
>Issue a new class.
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
request := nft.IssueDenomRequest{
	ID:             "a123",//The name of the collection
	Name:           "aad",//The name of the class
	Schema:         "test",
	Symbol:         "aads",
	MintRestricted: false,//MintRestricted is true means that only class owners can issue NFTs under this category, false means anyone can
	EditRestricted: false,//EditRestricted is true means that no one in this category can edit the NFT, false means that only the owner of this NFT can edit
}
res, err := client.Nft.IssueDenom(request, baseTx)
```

#### MintNFT<a name="mint"></a><br/>
>Issue a new nft.
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
request := nft.MintNFTRequest{
    ID:        "a1231",//The id of the nft
    ClassID:   "a123",//The name of the collection
    Name:      "aad",//The name of nft
    URI:       "https://www.baidu.com",//URI of off-chain NFT data
    Data:      string(getArmor()),//The data of the nft data 
    Recipient: "",
}
res, err := client.Nft.MintNFT(request, baseTx)

func getArmor() []byte {
    path, err := os.Getwd()
    if err != nil {
        panic(err)
    }
    path = filepath.Dir(path)
    path = filepath.Join(path, "github.com/oracleNetworkProtocol/plugchain-sdk-go/test/aad.info")
    bz, err := ioutil.ReadFile(path)
    if err != nil {
        panic(err)
    }
    return bz
}
```

#### EditNFT<a name="edit"></a><br/>
>edit a nft.
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
request := nft.EditNFTRequest{
    ID:        "a1231",//The id of the nft
    ClassID:   "a123",//The name of the collection
    Name:    "aads",
    URI:     "https://www.google.com",
    Data:    "this is a tree",
}
res, err := client.Nft.EditNFT(request, baseTx)
```

#### TransferNFT<a name="transfer_nft"></a><br/>
>transfer an NFT to a recipient.
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
request := nft.TransferNFTRequest{
    ID:        "a1231",//The id of the nft
    ClassID:   "a123",//The name of the collection
    Recipient: "gx1akqhezuftdcc0eqzkq5peqpjlucgmyr7srx54j",
}
res, err := client.Nft.TransferNFT(request, baseTx)
```

#### TransferClass<a name="transfer_class"></a><br/>
>transfer an Class to a recipient.
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
request := nft.TransferClassRequest{
    ID:        "a123",//The name of the collection
    Recipient: "gx1akqhezuftdcc0eqzkq5peqpjlucgmyr7srx54j",
}
res, err := client.Nft.TransferClass(request, baseTx)
```

#### BurnNFT<a name="burn"></a><br/>
>burn a nft.
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
request := nft.BurnNFTRequest{
    ID:        "a1231",//The id of the nft
    ClassID:   "a123",//The name of the collection
}
res, err := client.Nft.BurnNFT(request, baseTx)
```
