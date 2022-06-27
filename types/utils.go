package types

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	abci "github.com/tendermint/tendermint/abci/types"
	tmtypes "github.com/tendermint/tendermint/types"
	"math/big"
	"time"
)

// SortedJSON takes any JSON and returns it sorted by keys. Also, all white-spaces
// are removed.
// This method can be used to canonicalize JSON to be returned by GetSignBytes,
// e.g. for the ledger integration.
// If the passed JSON isn't valid it will return an error.
func SortJSON(toSortJSON []byte) ([]byte, error) {
	var c interface{}
	if err := json.Unmarshal(toSortJSON, &c); err != nil {
		return nil, err
	}
	return json.Marshal(c)
}

// MustSortJSON is like SortJSON but panic if an error occurs, e.g., if
// the passed JSON isn't valid.
func MustSortJSON(toSortJSON []byte) []byte {
	js, err := SortJSON(toSortJSON)
	if err != nil {
		panic(err)
	}
	return js
}

// Uint64ToBigEndian - marshals uint64 to a bigendian byte slice so it can be sorted
func Uint64ToBigEndian(i uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return b
}

// BigEndianToUint64 returns an uint64 from big endian encoded bytes. If encoding
// is empty, zero is returned.
func BigEndianToUint64(bz []byte) uint64 {
	if len(bz) == 0 {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

// Slight modification of the RFC3339Nano but it right pads all zeros and drops the time zone info
const SortableTimeFormat = "2006-01-02T15:04:05.000000000"

// Formats a time.Time into a []byte that can be sorted
func FormatTimeBytes(t time.Time) []byte {
	return []byte(t.UTC().Round(0).Format(SortableTimeFormat))
}

// Parses a []byte encoded using FormatTimeKey back into a time.Time
func ParseTimeBytes(bz []byte) (time.Time, error) {
	str := string(bz)
	t, err := time.Parse(SortableTimeFormat, str)
	if err != nil {
		return t, err
	}
	return t.UTC().Round(0), nil
}

// copy bytes
func CopyBytes(bz []byte) (ret []byte) {
	if bz == nil {
		return nil
	}
	ret = make([]byte, len(bz))
	copy(ret, bz)
	return ret
}

func GetOfferCoinFee(offerCoin Coin, swapFeeRate Dec) Coin {
	return NewCoin(offerCoin.Denom, offerCoin.Amount.ToDec().Mul(swapFeeRate.QuoInt64(2)).TruncateInt())
}

func NewPVMTransaction(tx *ethtypes.Transaction, blockHash common.Hash, blockNumber, index uint64) (*PvmResultQueryTx, error) {
	var signer types.Signer
	if tx.Protected() {
		signer = types.LatestSignerForChainID(tx.ChainId())
	} else {
		signer = types.HomesteadSigner{}
	}
	from, _ := types.Sender(signer, tx)
	v, r, s := tx.RawSignatureValues()
	al := tx.AccessList()
	result := &PvmResultQueryTx{
		BlockHash:        blockHash.String(),
		BlockNumber:      int64(blockNumber),
		From:             AccAddressFromHexAddress(from.String()),
		Gas:              tx.Gas(),
		GasPrice:         tx.GasPrice(),
		Hash:             tx.Hash(),
		Input:            tx.Data(),
		Nonce:            tx.Nonce(),
		TransactionIndex: index,
		Value:            tx.Value(),
		Type:             tx.Type(),
		Accesses:         &al,
		ContractAddress:  &AccAddress{},
		V:                v,
		R:                r,
		S:                s,
	}
	if tx.To() != nil {
		result.To = AccAddressFromHexAddress(tx.To().String())
	}
	if tx.To() == nil {
		ContractAddress := AccAddressFromHexAddress(crypto.CreateAddress(from, tx.Nonce()).String())
		result.ContractAddress = &ContractAddress
	}
	return result, nil
}

func FormatBlock(
	header tmtypes.Header, size int, gasLimit int64,
	gasUsed *big.Int, transactions []interface{}, bloom ethtypes.Bloom,
	validatorAddr common.Address,
) map[string]interface{} {
	var transactionsRoot common.Hash
	if len(transactions) == 0 {
		transactionsRoot = ethtypes.EmptyRootHash
	} else {
		transactionsRoot = common.BytesToHash(header.DataHash)
	}
	return map[string]interface{}{
		"number":           header.Height,
		"hash":             hexutil.Bytes(header.Hash()),
		"parentHash":       common.BytesToHash(header.LastBlockID.Hash.Bytes()),
		"nonce":            ethtypes.BlockNonce{},   // PoW specific
		"sha3Uncles":       ethtypes.EmptyUncleHash, // No uncles in Tendermint
		"logsBloom":        bloom,
		"stateRoot":        hexutil.Bytes(header.AppHash),
		"miner":            validatorAddr,
		"mixHash":          common.Hash{},
		"difficulty":       (*hexutil.Big)(big.NewInt(0)),
		"extraData":        "0x",
		"size":             hexutil.Uint64(size),
		"gasLimit":         hexutil.Uint64(gasLimit), // Static gas limit
		"gasUsed":          (*hexutil.Big)(gasUsed),
		"timestamp":        hexutil.Uint64(header.Time.Unix()),
		"transactionsRoot": transactionsRoot,
		"receiptsRoot":     ethtypes.EmptyRootHash,
		//"baseFeePerGas":    (*hexutil.Big)(baseFee),

		"uncles":          []common.Hash{},
		"transactions":    transactions,
		"totalDifficulty": (*hexutil.Big)(big.NewInt(0)),
	}
}

// SafeNewIntFromBigInt constructs Int from big.Int, return error if more than 256bits
func SafeNewIntFromBigInt(i *big.Int) (Int, error) {
	if !IsValidInt256(i) {
		return NewInt(0), fmt.Errorf("big int out of bound: %s", i)
	}
	return NewIntFromBigInt(i), nil
}

// IsValidInt256 check the bound of 256 bit number
func IsValidInt256(i *big.Int) bool {
	return i == nil || i.BitLen() <= maxBitLen
}

//BaseFeeFromEvents parses the feemarket basefee from cosmos events
func BaseFeeFromEvents(events []abci.Event) *big.Int {
	for _, event := range events {
		if event.Type != EventTypeFeeMarket {
			continue
		}

		for _, attr := range event.Attributes {
			if bytes.Equal(attr.Key, []byte(AttributeKeyBaseFee)) {
				result, success := new(big.Int).SetString(string(attr.Value), 10)
				if success {
					return result
				}
				return nil
			}
		}
	}
	return nil
}
