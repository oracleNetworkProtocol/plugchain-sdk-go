# PLUGCHAIN SDK GO

## KEYS MODULE

- [Show](#show) --Show
- [Add](#add) --Add
- [Recover](#recover) --Recover
- [Import](#import) --Import
- [Export](#export) --Export
- [Delete](#delete) --Delete
- [MnemonicImport](mnemonic) --MnemonicImport


# realization

#### Show<a name="show"></a><br/>
>Display address
**You need to import the private key before you can operate，Please see the key package for importing the private key**
```go
privKeyArmor, err := client.Key.Show("demo", "12312313")
fmt.Println(privKeyArmor)
fmt.Println(err)
```

#### Add<a name="add"></a><br/>
>Create new address
```go
address, mnemonic, err := client.Key.Add("demo", "12312313")
```

#### Recover<a name="recover"></a><br/>
>Restore address private key based on the help letter
```go
rs, err := client.Key.Recover("demo", "12312313", "camera torch fire elevator position good fringe turtle result subject language board angle agent mass mean measure lend yard north window mansion absurd exit")
```

#### Import<a name="import"></a><br/>
>Import address private key
```go
_, err = client.Key.Import("demo", "12312313", `-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: 6A2CF97123F50AA66EC7F0CF6B84F7B7
type: secp256k1

tfrFpzBp3UWGOdpq9KoZZ8KDVVHoN+6icHoCWUmUBaXA59AmjYymW1zZVkFsmDc3
iK4042axqRG8++28z5nXfSc4ZdDINWBbbe5xzX8=
=Bt49
-----END TENDERMINT PRIVATE KEY-----`)
```
OR
```go
_, err = client.Key.Import("demo", "12312313", string(getPrivKeyArmor()))
func getPrivKeyArmor() []byte {
    path, err := os.Getwd()
    if err != nil {
        panic(err)
    }
    path = filepath.Dir(path)
    path = filepath.Join(path, "github.com/oracleNetworkProtocol/plugchain-sdk-go/test/priv.info")
	bz, err := ioutil.ReadFile(path)
    if err != nil {
        panic(err)
    }
return bz
}
```

#### Export<a name="export"></a><br/>
>Export address private key
```go
privKeyArmor, err := client.Key.Export("demo", "12312313")
```

#### Delete<a name="delete"></a><br/>
>Delete address private key
```go
err = client.Key.Delete("demo", "12312313")
```

#### MnemonicImport<a name="mnemonic"></a><br/>
>Help note gain address
```go
mnemonic := "nerve leader thank marriage spice task van start piece crowd run hospital control outside cousin romance left choice poet wagon rude climb leisure spring"
km, err := crypto.NewMnemonicKeyManager(mnemonic, "secp256k1")
pubKey := km.ExportPubKey()
pubkeyBech32, err := types.Bech32ifyPubKey(types.Bech32PubKeyTypeAccPub, pubKey)
address := types.AccAddress(pubKey.Address()).String()
```