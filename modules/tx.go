package modules

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/modules/pvm"
	tmtypes "github.com/tendermint/tendermint/types"
	"math/big"
	"strings"
	"time"

	"github.com/gogo/protobuf/jsonpb"

	ctypes "github.com/tendermint/tendermint/rpc/core/types"

	clienttx "github.com/oracleNetworkProtocol/plugchain-sdk-go/client/tx"
	sdk "github.com/oracleNetworkProtocol/plugchain-sdk-go/types"
)

var bAttributeKeyEthereumBloom = []byte(sdk.AttributeKeyEthereumBloom)

// QueryTx returns the tx info
func (base baseClient) QueryTx(hash string) (sdk.ResultQueryTx, error) {
	tx, err := hex.DecodeString(hash)
	if err != nil {
		return sdk.ResultQueryTx{}, err
	}

	res, err := base.Tx(context.Background(), tx, true)
	if err != nil {
		return sdk.ResultQueryTx{}, err
	}

	resBlocks, err := base.getResultBlocks([]*ctypes.ResultTx{res})
	if err != nil {
		return sdk.ResultQueryTx{}, err
	}
	return base.parseTxResult(res, resBlocks[res.Height])
}

func (base baseClient) QueryTxs(builder *sdk.EventQueryBuilder, page, size *int) (sdk.ResultSearchTxs, error) {
	query := builder.Build()
	if len(query) == 0 {
		return sdk.ResultSearchTxs{}, errors.New("must declare at least one tag to search")
	}

	res, err := base.TxSearch(context.Background(), query, true, page, size, "asc")
	if err != nil {
		return sdk.ResultSearchTxs{}, err
	}

	resBlocks, err := base.getResultBlocks(res.Txs)
	if err != nil {
		return sdk.ResultSearchTxs{}, err
	}

	var txs []sdk.ResultQueryTx
	for i, tx := range res.Txs {
		txInfo, err := base.parseTxResult(tx, resBlocks[res.Txs[i].Height])
		if err != nil {
			return sdk.ResultSearchTxs{}, err
		}
		txs = append(txs, txInfo)
	}

	return sdk.ResultSearchTxs{
		Total: res.TotalCount,
		Txs:   txs,
	}, nil
}

func (base baseClient) TxSearchHandle(builder *sdk.EventQueryBuilder, page, size *int) (*ctypes.ResultTxSearch, error) {
	query := builder.Build()
	if len(query) == 0 {
		return nil, errors.New("must declare at least one tag to search")
	}

	return base.TxSearch(context.Background(), query, false, page, size, "asc")
}

func (base baseClient) QueryPvmTxs(res *ctypes.ResultTxSearch) ([]sdk.PvmResultQueryTx, error) {
	resBlocks, err := base.getResultBlocks(res.Txs)
	if err != nil {
		return []sdk.PvmResultQueryTx{}, err
	}

	if res.Txs == nil {
		return []sdk.PvmResultQueryTx{}, errors.New("Nonexistent Txs")
	}
	resBlock := resBlocks[res.Txs[0].Height]

	var txs []*pvm.MsgEthereumTx
	for _, tx := range resBlock.Block.Txs {
		tx, err := base.encodingConfig.TxConfig.TxDecoder()(tx)
		if err != nil {
			base.logger.Debug("failed to decode transaction in block", "height", resBlock.Block.Height, "error", err.Error())
			continue
		}

		for _, msg := range tx.GetMsgs() {
			ethMsg, ok := msg.(*pvm.MsgEthereumTx)
			if !ok {
				continue
			}

			hash := ethMsg.AsTransaction().Hash()
			builder := sdk.NewEventQueryBuilder()

			builder.AddCondition(sdk.NewCond(pvm.TypeMsgEthereumTx, pvm.AttributeKeyEthereumTxHash).EQ(hash))
			plugTx, err := base.TxSearchHandle(builder, nil, nil)

			if err != nil || (len(plugTx.Txs) > 1 && plugTx.Txs[0].Height != resBlock.Block.Height) {
				continue
			}

			txs = append(txs, ethMsg)
		}
	}

	//var txIndex uint64
	//
	//for i := range txs {
	//	if txs[i].Hash == res.Txs[0].Hash.String() {
	//		txIndex = uint64(i)
	//		break
	//	}
	//}
	//msg := txs[txIndex]
	//tx := msg.AsTransaction()
	//result, err := sdk.NewPVMTransaction(tx, common.BytesToHash(resBlock.Block.Hash()), uint64(resBlock.Block.Height), txIndex)
	var queryTxs []sdk.PvmResultQueryTx
	for txIndex, msg := range txs {
		tx := msg.AsTransaction()
		result, err := sdk.NewPVMTransaction(tx, common.BytesToHash(resBlock.Block.Hash()), uint64(resBlock.Block.Height), uint64(txIndex))
		if err != nil {
			continue
		}
		queryTxs = append(queryTxs, *result)
	}
	return queryTxs, nil
}

func (base baseClient) PvmBlockFromTendermint(block *tmtypes.Block,
	fullTx bool) (map[string]interface{}, error) {
	PVMTxs := []interface{}{}
	ctx := sdk.ContextWithHeight(block.Height)
	resBlockResult, err := base.TmClient.BlockResults(ctx, &block.Height)
	if err != nil {
		return nil, err
	}

	txResults := resBlockResult.TxsResults
	for i, txBz := range block.Txs {
		tx, err := base.encodingConfig.TxConfig.TxDecoder()(txBz)
		if err != nil {
			base.logger.Debug("failed to decode transaction in block", "height", block.Height, "error", err.Error())
			continue
		}
		for _, msg := range tx.GetMsgs() {
			ethMsg, ok := msg.(*pvm.MsgEthereumTx)
			if !ok {
				continue
			}
			tx := ethMsg.AsTransaction()
			if txResults[i].Code != 0 {
				base.logger.Debug("invalid tx result code", "hash", tx.Hash().Hex())
				continue
			}

			if !fullTx {
				hash := tx.Hash()
				PVMTxs = append(PVMTxs, hash)
				continue
			}
			pvmTx, err := sdk.NewPVMTransaction(tx, common.BytesToHash(block.Hash()), uint64(block.Height), uint64(i))
			if err != nil {
				base.logger.Debug("NewTransactionFromData for receipt failed", "hash", tx.Hash().Hex(), "error", err.Error())
				continue
			}
			PVMTxs = append(PVMTxs, pvmTx)
		}
	}

	bloom, err := base.BlockBloom(&block.Height)
	if err != nil {
		base.logger.Debug("failed to query BlockBloom", "height", block.Height, "error", err.Error())
	}

	req := &pvm.QueryValidatorAccountRequest{
		ConsAddress: sdk.ConsAddress(block.Header.ProposerAddress).String(),
	}

	conn, err := base.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return nil, err
	}

	res, err := pvm.NewQueryClient(conn).ValidatorAccount(context.Background(), req)
	if err != nil {
		return nil, err
	}
	addr, err := sdk.AccAddressFromBech32(res.AccountAddress)
	if err != nil {
		return nil, err
	}
	validatorAddr := common.BytesToAddress(addr)

	gasLimit, err := BlockMaxGasFromConsensusParams(ctx, base, block.Height)
	if err != nil {
		base.logger.Error("failed to query consensus params", "error", err.Error())
	}
	gasUsed := uint64(0)

	for _, txsResult := range txResults {
		gasUsed += uint64(txsResult.GetGasUsed())
	}

	formattenBlock := sdk.FormatBlock(block.Header, block.Size(), gasLimit, new(big.Int).SetUint64(gasUsed), PVMTxs, bloom, validatorAddr)
	return formattenBlock, nil
}

func (base baseClient) BlockBloom(height *int64) (ethtypes.Bloom, error) {
	result, err := base.BlockResults(context.Background(), height)
	if err != nil {
		return ethtypes.Bloom{}, err
	}
	for _, event := range result.EndBlockEvents {
		if event.Type != sdk.EventTypeBlockBloom {
			continue
		}

		for _, attr := range event.Attributes {
			if bytes.Equal(attr.Key, bAttributeKeyEthereumBloom) {
				return ethtypes.BytesToBloom(attr.Value), nil
			}
		}
	}
	return ethtypes.Bloom{}, errors.New("block bllom event is not fount")
}

func (base baseClient) QueryBlock(height int64) (sdk.BlockDetail, error) {
	block, err := base.Block(context.Background(), &height)
	if err != nil {
		return sdk.BlockDetail{}, err
	}

	blockResult, err := base.BlockResults(context.Background(), &height)
	if err != nil {
		return sdk.BlockDetail{}, err
	}

	return sdk.BlockDetail{
		BlockID:     block.BlockID,
		Block:       sdk.ParseBlock(base.encodingConfig.Amino, block.Block),
		BlockResult: sdk.ParseBlockResult(blockResult),
	}, nil
}

func (base baseClient) EstimateTxGas(txBytes []byte) (uint64, error) {
	res, err := base.ABCIQuery(context.Background(), "/app/simulate", txBytes)
	if err != nil {
		return 0, err
	}

	simRes, err := parseQueryResponse(res.Response.Value)
	if err != nil {
		return 0, err
	}

	adjusted := adjustGasEstimate(simRes.GasUsed, base.cfg.GasAdjustment)
	return adjusted, nil
}

func (base *baseClient) buildTx(msgs []sdk.Msg, baseTx sdk.BaseTx) ([]byte, *clienttx.Factory, sdk.Error) {
	builder, err := base.prepare(baseTx)
	if err != nil {
		return nil, builder, sdk.Wrap(err)
	}

	txByte, err := builder.BuildAndSign(baseTx.From, msgs, false)
	if err != nil {
		return nil, builder, sdk.Wrap(err)
	}

	base.Logger().Debug("sign transaction success")
	return txByte, builder, nil
}

func (base *baseClient) buildTxs(msgs []sdk.Msg, baseTx sdk.BaseTx, account map[string]string) ([]byte, *clienttx.Factory, sdk.Error) {
	builder, err := base.mutilprepare(baseTx, account)
	if err != nil {
		return nil, builder, sdk.Wrap(err)
	}

	txByte, err := builder.BuildAndSigns(msgs, false)
	if err != nil {
		return nil, builder, sdk.Wrap(err)
	}

	base.Logger().Debug("sign transaction success")
	return txByte, builder, nil
}

func (base *baseClient) buildTxWithAccount(addr string, accountNumber, sequence uint64, msgs []sdk.Msg, baseTx sdk.BaseTx) ([]byte, *clienttx.Factory, sdk.Error) {
	builder, err := base.prepareTemp(addr, accountNumber, sequence, baseTx)
	if err != nil {
		return nil, builder, sdk.Wrap(err)
	}

	txByte, err := builder.BuildAndSign(baseTx.From, msgs, false)
	if err != nil {
		return nil, builder, sdk.Wrap(err)
	}

	base.Logger().Debug("sign transaction success")
	return txByte, builder, nil
}

func (base baseClient) broadcastTx(txBytes []byte, mode sdk.BroadcastMode, simulate bool) (res sdk.ResultTx, err sdk.Error) {
	if simulate {
		estimateGas, err := base.EstimateTxGas(txBytes)
		if err != nil {
			return res, sdk.Wrap(err)
		}
		res.GasWanted = int64(estimateGas)
		return res, nil
	}

	switch mode {
	case sdk.Commit:
		res, err = base.broadcastTxCommit(txBytes)
	case sdk.Async:
		res, err = base.broadcastTxAsync(txBytes)
	case sdk.Sync:
		res, err = base.broadcastTxSync(txBytes)
	default:
		err = sdk.Wrapf("commit mode(%s) not supported", mode)
	}
	return
}

// broadcastTxCommit broadcasts transaction bytes to a Tendermint node
// and waits for a commit.
func (base baseClient) broadcastTxCommit(tx []byte) (sdk.ResultTx, sdk.Error) {
	res, err := base.BroadcastTxCommit(context.Background(), tx)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	if !res.CheckTx.IsOK() {
		return sdk.ResultTx{}, sdk.GetError(res.CheckTx.Codespace, res.CheckTx.Code, res.CheckTx.Log)
	}

	if !res.DeliverTx.IsOK() {
		return sdk.ResultTx{}, sdk.GetError(res.DeliverTx.Codespace, res.DeliverTx.Code, res.DeliverTx.Log)
	}

	return sdk.ResultTx{
		GasWanted: res.DeliverTx.GasWanted,
		GasUsed:   res.DeliverTx.GasUsed,
		Events:    sdk.StringifyEvents(res.DeliverTx.Events),
		Hash:      res.Hash.String(),
		Height:    res.Height,
	}, nil
}

// BroadcastTxSync broadcasts transaction bytes to a Tendermint node
// synchronously.
func (base baseClient) broadcastTxSync(tx []byte) (sdk.ResultTx, sdk.Error) {
	res, err := base.BroadcastTxSync(context.Background(), tx)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	if res.Code != 0 {
		return sdk.ResultTx{}, sdk.GetError(sdk.RootCodespace, res.Code, res.Log)
	}

	return sdk.ResultTx{Hash: res.Hash.String()}, nil
}

// BroadcastTxAsync broadcasts transaction bytes to a Tendermint node
// asynchronously.
func (base baseClient) broadcastTxAsync(tx []byte) (sdk.ResultTx, sdk.Error) {
	res, err := base.BroadcastTxAsync(context.Background(), tx)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	return sdk.ResultTx{Hash: res.Hash.String()}, nil
}

func (base baseClient) getResultBlocks(resTxs []*ctypes.ResultTx) (map[int64]*ctypes.ResultBlock, error) {
	resBlocks := make(map[int64]*ctypes.ResultBlock)
	for _, resTx := range resTxs {
		if _, ok := resBlocks[resTx.Height]; !ok {
			resBlock, err := base.Block(context.Background(), &resTx.Height)
			if err != nil {
				return nil, err
			}

			resBlocks[resTx.Height] = resBlock
		}
	}
	return resBlocks, nil
}

func (base baseClient) parseTxResult(res *ctypes.ResultTx, resBlock *ctypes.ResultBlock) (sdk.ResultQueryTx, error) {
	var tx sdk.Tx
	var err error

	if tx, err = base.encodingConfig.TxConfig.TxDecoder()(res.Tx); err != nil {
		return sdk.ResultQueryTx{}, err
	}

	return sdk.ResultQueryTx{
		Hash:   res.Hash.String(),
		Height: res.Height,
		Tx:     tx,
		Result: sdk.TxResult{
			Code:      res.TxResult.Code,
			Log:       res.TxResult.Log,
			GasWanted: res.TxResult.GasWanted,
			GasUsed:   res.TxResult.GasUsed,
			Events:    sdk.StringifyEvents(res.TxResult.Events),
		},
		Timestamp: resBlock.Block.Time.Format(time.RFC3339),
	}, nil
}

func adjustGasEstimate(estimate uint64, adjustment float64) uint64 {
	return uint64(adjustment * float64(estimate))
}

func parseQueryResponse(bz []byte) (sdk.SimulationResponse, error) {
	var simRes sdk.SimulationResponse
	if err := jsonpb.Unmarshal(strings.NewReader(string(bz)), &simRes); err != nil {
		return sdk.SimulationResponse{}, err
	}
	return simRes, nil
}

func BlockMaxGasFromConsensusParams(goCtx context.Context, clientCtx baseClient, blockHeight int64) (int64, error) {
	resConsParams, err := clientCtx.ConsensusParams(goCtx, &blockHeight)
	if err != nil {
		return int64(^uint32(0)), err
	}
	gasLimit := resConsParams.ConsensusParams.Block.MaxGas
	if gasLimit == -1 {
		gasLimit = int64(^uint32(0))
	}

	return gasLimit, nil
}
