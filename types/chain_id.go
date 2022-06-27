package types

import (
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

var (
	regexChainID         = `[a-z]{1,}`
	regexEIP155Separator = `_{1}`
	regexEIP155          = `[1-9][0-9]*`
	regexEpochSeparator  = `-{1}`
	regexEpoch           = `[1-9][0-9]*`
	ethermintChainID     = regexp.MustCompile(fmt.Sprintf(`^(%s)%s(%s)%s(%s)$`, regexChainID, regexEIP155Separator, regexEIP155, regexEpochSeparator, regexEpoch))
)

// ParseChainID parses a string chain identifier's epoch to an Ethereum-compatible
// chain-id in *big.Int format. The function returns an error if the chain-id has an invalid format
func ParseChainID(chainId string) (*big.Int, error) {
	chainId = strings.TrimSpace(chainId)
	if len(chainId) > 48 {
		return nil, errors.New(fmt.Sprintf("invalid chain ID. chain-id '%s' cannot exceed 48 chars", chainId))
	}
	matches := ethermintChainID.FindStringSubmatch(chainId)
	if matches == nil || len(matches) != 4 || matches[1] == "" {
		return nil, errors.New(fmt.Sprintf("invalid chain ID. %v:%v", chainId, matches))
	}

	chainIDInt, ok := new(big.Int).SetString(matches[2], 10)
	if !ok {
		return nil, errors.New(fmt.Sprintf("invalid chain ID. epoch %s must be base-10 integer format", matches[2]))
	}
	return chainIDInt, nil
}
