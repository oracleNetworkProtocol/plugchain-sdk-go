package types

import (
	"fmt"
	"os"

	"github.com/oracleNetworkProtocol/plugchain-sdk-go/types/store"
)

const (
	defaultGas           = 200000
	defaultTimeout       = 5
	defaultPath          = "$HOME/plugchain-sdk-go/leveldb"
	defaultLogLevel      = "info"
	defaultAlgo          = "eth_secp256k1"
	defaultFees          = "200uplugcn"
	defaultMode          = Sync
	defaultGasAdjustment = 1.0
)

type ClientConfig struct {
	//plugChain node rpc address
	NodeURL string

	//plugChain grpc address
	GRPCAddr string

	//plugChain chain-id
	ChainId string

	//mas gas limit
	Gas uint64

	//Fee
	Fee DecCoins

	//PrivKeyArmor DAO Implements
	KeyDAO store.KeyDAO

	// Private key generation algorithm(sm2,secp256k1)
	Algo string

	//Broadcast transaction mode
	Mode BroadcastMode

	//Transaction broadcast timeout(seconds)
	Timeout uint

	//log level(trace|debug|info|warn|error|fatal|panic)
	LogLevel string

	//adjustment factor to be multiplied against the estimate returned by the tx simulation;
	GasAdjustment float64

	//whether to enable caching
	Cached bool
}

func NewClientConfig(url, grpcAddr, chainId string, options ...Option) (ClientConfig, error) {
	cfg := ClientConfig{
		NodeURL:  url,
		GRPCAddr: grpcAddr,
		ChainId:  chainId,
	}
	for _, optionFn := range options {
		if err := optionFn(&cfg); err != nil {
			return ClientConfig{}, err
		}
	}
	if err := cfg.checkAndSetDefault(); err != nil {
		return ClientConfig{}, err
	}
	return cfg, nil
}

func (cfg *ClientConfig) checkAndSetDefault() error {
	if len(cfg.NodeURL) == 0 {
		return fmt.Errorf("nodeURI is required")
	}

	if len(cfg.ChainId) == 0 {
		return fmt.Errorf("chainID is required")
	}

	if err := GasOption(cfg.Gas)(cfg); err != nil {
		return err
	}

	if err := FeeOption(cfg.Fee)(cfg); err != nil {
		return err
	}

	if err := AlgoOption(cfg.Algo)(cfg); err != nil {
		return err
	}

	if err := KeyDAOOption(cfg.KeyDAO)(cfg); err != nil {
		return err
	}

	if err := ModeOption(cfg.Mode)(cfg); err != nil {
		return err
	}

	if err := TimeoutOption(cfg.Timeout)(cfg); err != nil {
		return err
	}

	if err := LogLevelOption(cfg.LogLevel)(cfg); err != nil {
		return err
	}

	return GasAdjustmentOption(cfg.GasAdjustment)(cfg)
}

func AlgoOption(algo string) Option {
	return func(cfg *ClientConfig) error {
		if algo == "" {
			algo = defaultAlgo
		}
		cfg.Algo = algo
		return nil
	}
}

func ModeOption(mode BroadcastMode) Option {
	return func(cfg *ClientConfig) error {
		if mode == "" {
			mode = defaultMode
		}
		cfg.Mode = mode
		return nil
	}
}

func FeeOption(fee DecCoins) Option {
	return func(cfg *ClientConfig) error {
		if fee == nil || fee.Empty() || !fee.IsValid() {
			fees, _ := ParseDecCoins(defaultFees)
			fee = fees
		}
		cfg.Fee = fee
		return nil
	}
}

func GasAdjustmentOption(gasAdjustment float64) Option {
	return func(cfg *ClientConfig) error {
		if gasAdjustment <= 0 {
			gasAdjustment = defaultGasAdjustment
		}
		cfg.GasAdjustment = gasAdjustment
		return nil
	}
}

type Option func(cfg *ClientConfig) error

func GasOption(gas uint64) Option {
	return func(cfg *ClientConfig) error {
		if gas <= 0 {
			gas = defaultGas
		}
		cfg.Gas = gas
		return nil
	}
}

func KeyDAOOption(dao store.KeyDAO) Option {
	return func(cfg *ClientConfig) error {
		if dao == nil {
			defaultPath := os.ExpandEnv(defaultPath)
			levelDB, err := store.NewLevelDB(defaultPath, nil)
			if err != nil {
				return err
			}
			dao = levelDB
		}
		cfg.KeyDAO = dao
		return nil
	}
}

func TimeoutOption(timeout uint) Option {
	return func(cfg *ClientConfig) error {
		if timeout <= 0 {
			timeout = defaultTimeout
		}
		cfg.Timeout = timeout
		return nil
	}
}

func LogLevelOption(level string) Option {
	return func(cfg *ClientConfig) error {
		if level == "" {
			level = defaultLogLevel
		}
		cfg.LogLevel = level
		return nil
	}
}

func CachedOption(enabled bool) Option {
	return func(cfg *ClientConfig) error {
		cfg.Cached = enabled
		return nil
	}
}
