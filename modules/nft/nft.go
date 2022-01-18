package nft

import (
	"context"
	"plugchain-sdk-go/codec"
	"plugchain-sdk-go/codec/types"
	"plugchain-sdk-go/types/query"

	sdk "plugchain-sdk-go/types"
)

type nftClient struct {
	sdk.BaseClient
	codec.Marshaler
}

func NewClient(bc sdk.BaseClient, cdc codec.Marshaler) Client {
	return nftClient{
		BaseClient: bc,
		Marshaler:  cdc,
	}
}

func (nc nftClient) Name() string {
	return ModuleName
}

func (nc nftClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (nc nftClient) IssueDenom(request IssueDenomRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	sender, err := nc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := &MsgIssueClass{
		ID:             request.ID,
		Name:           request.Name,
		Schema:         request.Schema,
		Owner:          sender.String(),
		Symbol:         request.Symbol,
		MintRestricted: request.MintRestricted,
		EditRestricted: request.EditRestricted,
	}
	return nc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (nc nftClient) MintNFT(request MintNFTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	sender, err := nc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	var recipient = sender.String()
	if len(request.Recipient) > 0 {
		if err := sdk.ValidateAccAddress(request.Recipient); err != nil {
			return sdk.ResultTx{}, sdk.Wrap(err)
		}
		recipient = request.Recipient
	}

	msg := &MsgIssueNFT{
		ID:        request.ID,
		ClassID:   request.ClassID,
		Name:      request.Name,
		URI:       request.URI,
		Data:      request.Data,
		Owner:     sender.String(),
		Recipient: recipient,
	}
	return nc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (nc nftClient) EditNFT(request EditNFTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	sender, err := nc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := &MsgEditNFT{
		ID:      request.ID,
		Name:    request.Name,
		ClassID: request.ClassID,
		URI:     request.URI,
		Data:    request.Data,
		Owner:   sender.String(),
	}
	return nc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (nc nftClient) TransferNFT(request TransferNFTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	sender, err := nc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	if err := sdk.ValidateAccAddress(request.Recipient); err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := &MsgTransferNFT{
		ID:        request.ID,
		ClassID:   request.ClassID,
		Recipient: request.Recipient,
		Owner:     sender.String(),
	}
	return nc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (nc nftClient) TransferClass(request TransferClassRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	sender, err := nc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	if err := sdk.ValidateAccAddress(request.Recipient); err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := &MsgTransferClass{
		ID:        request.ID,
		Recipient: request.Recipient,
		Owner:     sender.String(),
	}
	return nc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (nc nftClient) BurnNFT(request BurnNFTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	sender, err := nc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := &MsgBurnNFT{
		Owner:   sender.String(),
		ID:      request.ID,
		ClassID: request.ClassID,
	}
	return nc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (nc nftClient) QuerySupply(denom string) (uint64, sdk.Error) {
	if len(denom) == 0 {
		return 0, sdk.Wrapf("denom is required")
	}

	conn, err := nc.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return 0, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).Supply(
		context.Background(),
		&QuerySupplyRequest{
			ClassId: denom,
		},
	)
	if err != nil {
		return 0, sdk.Wrap(err)
	}

	return res.Amount, nil
}

func (nc nftClient) QueryOwner(creator, denom string, req sdk.PageRequest) (QueryOwnerResp, sdk.Error) {
	if len(denom) == 0 {
		return QueryOwnerResp{}, sdk.Wrapf("denom is required")
	}

	if err := sdk.ValidateAccAddress(creator); err != nil {
		return QueryOwnerResp{}, sdk.Wrap(err)
	}

	conn, err := nc.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return QueryOwnerResp{}, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).Owner(
		context.Background(),
		&QueryOwnerRequest{
			Address: creator,
			ClassId: denom,
			Pagination: &query.PageRequest{
				Key:        req.Key,
				Offset:     req.Offset,
				Limit:      req.Limit,
				CountTotal: req.CountTotal,
			},
		},
	)
	if err != nil {
		return QueryOwnerResp{}, sdk.Wrap(err)
	}

	return res.Owner.Convert().(QueryOwnerResp), nil
}

func (nc nftClient) QueryCollection(denom string, req sdk.PageRequest) (QueryCollectionResp, sdk.Error) {
	if len(denom) == 0 {
		return QueryCollectionResp{}, sdk.Wrapf("denom is required")
	}

	conn, err := nc.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return QueryCollectionResp{}, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).Collection(
		context.Background(),
		&QueryCollectionRequest{
			ClassId: denom,
			Pagination: &query.PageRequest{
				Key:        req.Key,
				Offset:     req.Offset,
				Limit:      req.Limit,
				CountTotal: req.CountTotal,
			}},
	)
	if err != nil {
		return QueryCollectionResp{}, sdk.Wrap(err)
	}

	return res.Collection.Convert().(QueryCollectionResp), nil
}

func (nc nftClient) QueryDenoms(req sdk.PageRequest) ([]QueryDenomResp, sdk.Error) {
	conn, err := nc.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).Classes(
		context.Background(),
		&QueryClassesRequest{
			Pagination: &query.PageRequest{
				Key:        req.Key,
				Offset:     req.Offset,
				Limit:      req.Limit,
				CountTotal: req.CountTotal,
			}},
	)
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	return denoms(res.Classes).Convert().([]QueryDenomResp), nil
}

func (nc nftClient) QueryDenom(denom string) (QueryDenomResp, sdk.Error) {
	conn, err := nc.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return QueryDenomResp{}, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).Class(
		context.Background(),
		&QueryClassRequest{ClassId: denom},
	)
	if err != nil {
		return QueryDenomResp{}, sdk.Wrap(err)
	}

	return res.Class.Convert().(QueryDenomResp), nil
}

func (nc nftClient) QueryNFT(denom, ID string) (QueryNFTResp, sdk.Error) {
	if len(denom) == 0 {
		return QueryNFTResp{}, sdk.Wrapf("denom is required")
	}

	if len(ID) == 0 {
		return QueryNFTResp{}, sdk.Wrapf("ID is required")
	}

	conn, err := nc.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return QueryNFTResp{}, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).NFT(
		context.Background(),
		&QueryNFTRequest{
			ClassId: denom,
			NftId:   ID,
		},
	)
	if err != nil {
		return QueryNFTResp{}, sdk.Wrap(err)
	}

	return res.Nft.Convert().(QueryNFTResp), nil
}
