package pvm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	sdk "github.com/oracleNetworkProtocol/plugchain-sdk-go/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"math/big"
)

const (
	ModuleName = "pvm"
)

// TransactionArgs represents the arguments to construct a new transaction
// or a message call using JSON-RPC.
// Duplicate struct definition since geth struct is in internal package
// Ref: https://github.com/ethereum/go-ethereum/blob/release/1.10.4/internal/ethapi/transaction_args.go#L36
type TransactionArgs struct {
	From                 *common.Address `json:"from"`
	To                   *common.Address `json:"to"`
	Gas                  *hexutil.Uint64 `json:"gas"`
	GasPrice             *hexutil.Big    `json:"gasPrice"`
	MaxFeePerGas         *hexutil.Big    `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *hexutil.Big    `json:"maxPriorityFeePerGas"`
	Value                *hexutil.Big    `json:"value"`
	Nonce                *hexutil.Uint64 `json:"nonce"`

	// We accept "data" and "input" for backwards-compatibility reasons.
	// "input" is the newer name and should be preferred by clients.
	// Issue detail: https://github.com/ethereum/go-ethereum/issues/15628
	Data  *hexutil.Bytes `json:"data"`
	Input *hexutil.Bytes `json:"input"`

	// Introduced by AccessListTxType transaction.
	AccessList *ethtypes.AccessList `json:"accessList,omitempty"`
	ChainID    *hexutil.Big         `json:"chainId,omitempty"`
}

var (
	_ sdk.Msg = &MsgEthereumTx{}
)

func (al AccessList) ToEthAccessList() *ethtypes.AccessList {
	var ethAccessList ethtypes.AccessList

	for _, tuple := range al {
		storageKeys := make([]common.Hash, len(tuple.StorageKeys))

		for i := range tuple.StorageKeys {
			storageKeys[i] = common.HexToHash(tuple.StorageKeys[i])
		}

		ethAccessList = append(ethAccessList, ethtypes.AccessTuple{
			Address:     common.HexToAddress(tuple.Address),
			StorageKeys: storageKeys,
		})
	}

	return &ethAccessList
}

func TxLogsFromEvents(events []abci.Event) ([]*PvmLog, error) {
	logs := make([]*Log, 0)
	for _, event := range events {
		if event.Type != sdk.EventTypeTxLog {
			continue
		}

		for _, attr := range event.Attributes {
			if !bytes.Equal(attr.Key, []byte(sdk.AttributeKeyTxLog)) {
				continue
			}

			var log Log
			if err := json.Unmarshal(attr.Value, &log); err != nil {
				return nil, err
			}

			logs = append(logs, &log)
		}
	}
	return LogsToPvmLog(logs), nil
}

func LogsToPvmLog(logs []*Log) []*PvmLog {
	var pvmLogs []*PvmLog // nolint: prealloc
	for i := range logs {
		pvmLogs = append(pvmLogs, logs[i].ToPvmLog())
	}
	return pvmLogs
}

func (log *Log) ToPvmLog() *PvmLog {
	var topics []common.Hash // nolint: prealloc
	for i := range log.Topics {
		topics = append(topics, common.HexToHash(log.Topics[i]))
	}

	return &PvmLog{
		Address:     sdk.AccAddressFromHexAddress(log.Address),
		Topics:      topics,
		Data:        log.Data,
		BlockNumber: log.BlockNumber,
		TxHash:      common.HexToHash(log.TxHash),
		TxIndex:     uint(log.TxIndex),
		Index:       uint(log.Index),
		BlockHash:   common.HexToHash(log.BlockHash),
		Removed:     log.Removed,
	}
}

// checkTxFee is an internal function used to check whether the fee of
// the given transaction is _reasonable_(under the cap).
func checkTxFee(gasPrice *big.Int, gas uint64, cap float64) error {
	// Short circuit if there is no cap for transaction fee at all.
	if cap == 0 {
		return nil
	}
	feePlug := new(big.Float).Quo(new(big.Float).SetInt(new(big.Int).Mul(gasPrice, new(big.Int).SetUint64(gas))), new(big.Float).SetInt(big.NewInt(sdk.UplugCn)))
	feeFloat, _ := feePlug.Float64()
	if feeFloat > cap {
		return fmt.Errorf("tx fee (%.2f uplugcn) exceeds the configured cap (%.2f uplugcn)", feeFloat, cap)
	}
	return nil
}
