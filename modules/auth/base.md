# PLUGCHAIN SDK GO

## BASE MODULE

- [Query](#query)
    - [Block](#block) --Account Amount
    - [QueryTx](#quert_tx) --TotalSupply
    - [QueryAddress](#query_address) --QueryAddress
    - [QueryBlock](#query_blog) --QueryBlock
    - [BlockByHash](#block_hash) --BlockByHash

# realization

## Query<a name="query"></a><br/>

#### Block<a name="account"></a><br/>
>Query the latest block
```go
block, err := client.BaseClient.Block(context.Background(), nil)
```

#### QueryTx<a name="quert_tx"></a><br/>
>Query hash information
```go
tx, err := client.BaseClient.QueryTx("4CBE93F90230B6C1AF324D530858D2087E0D9A6F26DFDAC7842110284AF5728D")
```

#### QueryAddress<a name="query_address"></a><br/>
>display address
```go
address, err := client.BaseClient.QueryAddress("demo", "123123123")
```

#### QueryBlock<a name="query_blog"></a><br/>
>Query block information according to block height
```go
blog, err := client.BaseClient.QueryBlock(12)
hash:=blog.Block.Hash()
```

#### BlockByHash<a name="block_hash"></a><br/>
>Query block information according to block hash
```go
hash, err := types.HexBytesFrom("6488C66BEE3E246901ECAEAD01067B4B24F50682ADB7C39289A70B0857DE2308")
blogs, err := client.BaseClient.BlockByHash(context.Background(), hash)
```