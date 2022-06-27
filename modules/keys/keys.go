package keys

import (
	"github.com/ethereum/go-ethereum/crypto"
	crypto2 "github.com/oracleNetworkProtocol/plugchain-sdk-go/crypto"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/crypto/keys/ethsecp256k1"
	sdk "github.com/oracleNetworkProtocol/plugchain-sdk-go/types"
)

type keysClient struct {
	sdk.KeyManager
}

func NewClient(keyManager sdk.KeyManager) Client {
	return keysClient{keyManager}
}

func (k keysClient) Add(name, password string) (string, string, sdk.Error) {
	address, mnemonic, err := k.Insert(name, password)
	return address, mnemonic, sdk.Wrap(err)
}

func (k keysClient) Recover(name, password, mnemonic string) (string, sdk.Error) {
	address, err := k.KeyManager.Recover(name, password, mnemonic, "")
	return address, sdk.Wrap(err)
}

func (k keysClient) RecoverWithHDPath(name, password, mnemonic, hdPath string) (string, sdk.Error) {
	address, err := k.KeyManager.Recover(name, password, mnemonic, hdPath)
	return address, sdk.Wrap(err)
}

func (k keysClient) Import(name, password, privKeyArmor string) (string, sdk.Error) {
	address, err := k.KeyManager.Import(name, password, privKeyArmor)
	return address, sdk.Wrap(err)
}

func (k keysClient) Export(name, password string) (string, sdk.Error) {
	keystore, err := k.KeyManager.Export(name, password)
	return keystore, sdk.Wrap(err)
}

func (k keysClient) Delete(name, password string) sdk.Error {
	err := k.KeyManager.Delete(name, password)
	return sdk.Wrap(err)
}

func (k keysClient) Show(name, password string) (string, sdk.Error) {
	_, address, err := k.KeyManager.Find(name, password)
	if err != nil {
		return "", sdk.Wrap(err)
	}
	return address.String(), nil
}

func (k keysClient) Ethsecp256k1TOSecp256k1(keystr, password string) string {
	priv, err := crypto.HexToECDSA(keystr)
	if err != nil {
		return keystr
	}
	priKey := &ethsecp256k1.PrivKey{Key: crypto.FromECDSA(priv)}
	secp256k1Keystr := crypto2.EncryptArmorPrivKey(priKey, password, ethsecp256k1.Ethsecp256k1keyType)
	return secp256k1Keystr
}

func (k keysClient) ExportEthsecp256k1(name, password string) (string, error) {
	priv, err := k.KeyManager.ExportEthsecp256k1(name, password)
	if err != nil {
		return "", err
	}
	return priv, nil
}
