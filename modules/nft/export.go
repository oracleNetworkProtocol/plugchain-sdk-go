package nft

import sdk "plugchain-sdk-go/types"

// expose NFT module api for user
type Client interface {
	sdk.Module

	IssueDenom(request IssueDenomRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	MintNFT(request MintNFTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	EditNFT(request EditNFTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	TransferNFT(request TransferNFTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	TransferClass(request TransferClassRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	BurnNFT(request BurnNFTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)

	QuerySupply(denomID string) (uint64, sdk.Error)
	QueryOwner(creator, classId string, pageReq sdk.PageRequest) (QueryOwnerResp, sdk.Error)
	QueryCollection(denomID string, pageReq sdk.PageRequest) (QueryCollectionResp, sdk.Error)
	QueryDenom(denomID string) (QueryDenomResp, sdk.Error)
	QueryDenoms(pageReq sdk.PageRequest) ([]QueryDenomResp, sdk.Error)
	QueryNFT(denomID, ID string) (QueryNFTResp, sdk.Error)
}

type IssueDenomRequest struct {
	ID             string
	Name           string
	Schema         string
	Symbol         string
	MintRestricted bool
	EditRestricted bool
}

type MintNFTRequest struct {
	ID        string `json:"id"`
	ClassID   string `json:"class_id"`
	Name      string `json:"name"`
	URI       string `json:"uri"`
	Data      string `json:"data"`
	Recipient string `json:"recipient"`
}

type EditNFTRequest struct {
	ID      string `json:"id"`
	ClassID string `json:"class_id"`
	Name    string `json:"name"`
	URI     string `json:"uri"`
	Data    string `json:"data"`
}

type TransferClassRequest struct {
	ID        string `json:"id"`
	Recipient string `json:"recipient"`
}

type TransferNFTRequest struct {
	ID        string `json:"id"`
	ClassID   string `json:"class_id"`
	Recipient string `json:"recipient"`
}

type BurnNFTRequest struct {
	ClassID string `json:"class_id"`
	ID      string `json:"id"`
}

// IDC defines a set of nft ids that belong to a specific
type IDC struct {
	Denom    string   `json:"denom" yaml:"denom"`
	TokenIDs []string `json:"token_ids" yaml:"token_ids"`
}

type QueryOwnerResp struct {
	Address string `json:"address" yaml:"address"`
	IDCs    []IDC  `json:"idcs" yaml:"idcs"`
}

// BaseNFT non fungible token definition
type QueryNFTResp struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	URI     string `json:"uri"`
	Data    string `json:"data"`
	Creator string `json:"creator"`
}

type QueryDenomResp struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Schema         string `json:"schema"`
	Symbol         string `json:"symbol"`
	Owner          string `json:"owner"`
	MintRestricted bool   `json:"mint_restricted"`
	EditRestricted bool   `json:"edit_restricted"`
}

type QueryCollectionResp struct {
	Denom QueryDenomResp `json:"denom" yaml:"denom"`
	NFTs  []QueryNFTResp `json:"nfts" yaml:"nfts"`
}
