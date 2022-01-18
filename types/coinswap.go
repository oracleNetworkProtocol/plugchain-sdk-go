package types

type PageRequest struct {
	// key is a value returned in PageResponse.next_key to begin
	// querying the next page most efficiently. Only one of offset or key
	// should be set.
	Key []byte `json:"key,omitempty"`
	// offset is a numeric offset that can be used when key is unavailable.
	// It is less efficient than using key. Only one of offset or key should
	// be set.
	Offset uint64 ` json:"offset,omitempty"`
	// limit is the total number of results to be returned in the result page.
	// If left empty it will default to a value to be set by each app.
	Limit uint64 ` json:"limit,omitempty"`
	// count_total is set to true  to indicate that the result set should include
	// a count of the total number of items available for pagination in UIs.
	// count_total is only respected when offset is used. It is ignored when key
	// is set.
	CountTotal bool ` json:"count_total,omitempty"`
}

type PoolInfo struct {
	// id of the pool
	Id uint64 `json:"id"`
	// id of the pool_type
	TypeId uint32 `json:"type_id,omitempty"`
	// denoms of reserve coin pair of the pool
	ReserveCoinDenoms []string `json:"reserve_coin_denoms,omitempty"`
	// reserve account address of the pool
	ReserveAccountAddress string `json:"reserve_account_address,omitempty"`
	// denom of pool coin of the pool
	PoolCoinDenom string `json:"pool_coin_denom,omitempty"`
}
