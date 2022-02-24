package token

import (
	sdk "github.com/oracleNetworkProtocol/plugchain-sdk-go/types"
)

type Client interface {
	sdk.Module

	IssueToken(req IssueTokenRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	EditToken(req EditTokenRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	TransferToken(to string, symbol string, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	MintToken(symbol string, amount uint64, to string, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)

	QueryToken(symbol string) (sdk.Token, error)
	QueryTokens(owner string) (sdk.Tokens, error)
	QueryFees(symbol string) (QueryFeesResp, error)
	QueryParams() (QueryParamsResp, error)
}

type IssueTokenRequest struct {
	Symbol        string `json:"symbol"`
	Name          string `json:"name"`
	Scale         uint32 `json:"scale"`
	MinUnit       string `json:"min_unit"`
	InitialSupply uint64 `json:"initial_supply"`
	MaxSupply     uint64 `json:"max_supply"`
	Mintable      bool   `json:"mintable"`
}

type EditTokenRequest struct {
	Symbol    string `json:"symbol"`
	Name      string `json:"name"`
	MaxSupply uint64 `json:"max_supply"`
}

// QueryFeesResp is for the token fees query output
type QueryFeesResp struct {
	Exist    bool     `json:"exist"`     // indicate if the token has existed
	IssueFee sdk.Coin `json:"issue_fee"` // issue fee
	MintFee  sdk.Coin `json:"mint_fee"`  // mint fee
}

// token params
type QueryParamsResp struct {
	IssueTokenBaseFee    string `json:"issue_token_base_fee"`    // e.g., 300000*10^18
	OperateTokenFeeRatio string `json:"operate_token_fee_ratio"` // e.g., 10%
}
