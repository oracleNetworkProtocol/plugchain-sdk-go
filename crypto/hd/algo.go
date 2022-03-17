package hd

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	bip39 "github.com/cosmos/go-bip39"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/crypto/types"
	tmcrypto "github.com/tendermint/tendermint/crypto"
	tylerbip39 "github.com/tyler-smith/go-bip39"

	"github.com/oracleNetworkProtocol/plugchain-sdk-go/crypto/keys/ethsecp256k1"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/crypto/keys/secp256k1"
)

type SignatureAlgo interface {
	Name() PubKeyType
	Derive() DeriveFn
	Generate() GenerateFn
}

func NewSigningAlgoFromString(str string) (SignatureAlgo, error) {
	switch str {
	case string(Secp256k1.Name()):
		return Secp256k1, nil
	case string(EthSecp256k1.Name()):
		return EthSecp256k1, nil
	default:
		return nil, fmt.Errorf("provided algorithm `%s` is not supported", str)
	}
}

// PubKeyType defines an algorithm to derive key-pairs which can be used for cryptographic signing.
type PubKeyType string

const (
	// MultiType implies that a pubkey is a multisignature
	MultiType = PubKeyType("multi")
	// Secp256k1Type uses the Bitcoin secp256k1 ECDSA parameters.
	Secp256k1Type    = PubKeyType("secp256k1")
	EthSecp256k1Type = PubKeyType("eth_secp256k1")
	// Ed25519Type represents the Ed25519Type signature system.
	// It is currently not supported for end-user keys (wallets/ledgers).
	Ed25519Type = PubKeyType("ed25519")
	// Sr25519Type represents the Sr25519Type signature system.
	Sr25519Type = PubKeyType("sr25519")
)

var (
	// Secp256k1 uses the Bitcoin secp256k1 ECDSA parameters.
	Secp256k1    = secp256k1Algo{}
	EthSecp256k1 = ethsecp256k1Algo{}
)

type DeriveFn func(mnemonic string, bip39Passphrase, hdPath string) ([]byte, error)
type GenerateFn func(bz []byte) types.PrivKey

type WalletGenerator interface {
	Derive(mnemonic string, bip39Passphrase, hdPath string) ([]byte, error)
	Generate(bz []byte) tmcrypto.PrivKey
}

type secp256k1Algo struct {
}

func (s secp256k1Algo) Name() PubKeyType {
	return Secp256k1Type
}

// Derive derives and returns the secp256k1 private key for the given seed and HD path.
func (s secp256k1Algo) Derive() DeriveFn {
	return func(mnemonic string, bip39Passphrase, hdPath string) ([]byte, error) {
		seed, err := bip39.NewSeedWithErrorChecking(mnemonic, bip39Passphrase)
		if err != nil {
			return nil, err
		}

		masterPriv, ch := ComputeMastersFromSeed(seed)
		if len(hdPath) == 0 {
			return masterPriv[:], nil
		}
		derivedKey, err := DerivePrivateKeyForPath(masterPriv, ch, hdPath)

		return derivedKey, err
	}
}

// Generate generates a secp256k1 private key from the given bytes.
func (s secp256k1Algo) Generate() GenerateFn {
	return func(bz []byte) types.PrivKey {
		var bzArr = make([]byte, secp256k1.PrivKeySize)
		copy(bzArr, bz)

		return &secp256k1.PrivKey{Key: bzArr}
	}
}

type ethsecp256k1Algo struct {
}

func (s ethsecp256k1Algo) Name() PubKeyType {
	return EthSecp256k1Type
}

func (s ethsecp256k1Algo) Derive() DeriveFn {
	return func(mnemonic, bip39Passphrase, path string) ([]byte, error) {
		hdpath, err := accounts.ParseDerivationPath(path)
		if err != nil {
			return nil, err
		}

		seed, err := tylerbip39.NewSeedWithErrorChecking(mnemonic, bip39Passphrase)
		if err != nil {
			return nil, err
		}

		// create a BTC-utils hd-derivation key chain
		masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
		if err != nil {
			return nil, err
		}

		key := masterKey
		for _, n := range hdpath {
			key, err = key.Derive(n)
			if err != nil {
				return nil, err
			}
		}

		// btc-utils representation of a secp256k1 private key
		privateKey, err := key.ECPrivKey()
		if err != nil {
			return nil, err
		}

		// cast private key to a convertible form (single scalar field element of secp256k1)
		// and then load into ethcrypto private key format.
		// TODO: add links to godocs of the two methods or implementations of them, to compare equivalency
		privateKeyECDSA := privateKey.ToECDSA()
		derivedKey := crypto.FromECDSA(privateKeyECDSA)

		return derivedKey, nil
	}
}

// Generate generates a eth_secp256k1 private key from the given bytes.
func (s ethsecp256k1Algo) Generate() GenerateFn {
	return func(bz []byte) types.PrivKey {
		bzArr := make([]byte, ethsecp256k1.Ethsecp256k1PrivKeySize)
		copy(bzArr, bz)

		// TODO: modulo P
		return &ethsecp256k1.PrivKey{
			Key: bzArr,
		}
	}
}
