package plugchain_sdk

import (
	"fmt"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/codec"
	cdctypes "github.com/oracleNetworkProtocol/plugchain-sdk-go/codec/types"
	cryptocodec "github.com/oracleNetworkProtocol/plugchain-sdk-go/crypto/codec"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/modules"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/modules/bank"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/modules/coinswap"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/modules/gov"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/modules/keys"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/modules/nft"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/modules/staking"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/modules/token"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/types"
	txtypes "github.com/oracleNetworkProtocol/plugchain-sdk-go/types/tx"
	"github.com/tendermint/tendermint/libs/log"
)

type PLUGCHAINClient struct {
	logger         log.Logger
	moduleManager  map[string]types.Module
	encodingConfig types.EncodingConfig
	types.BaseClient
	Key     keys.Client
	Bank    bank.Client
	Token   token.Client
	Swap    coinswap.Client
	Nft     nft.Client
	Gov     gov.Client
	Staking staking.Client
}

func NewPLUGCHAINClient(cfg types.ClientConfig) PLUGCHAINClient {
	encodingConfig := makeEncodingConfig()

	//Create basic client instance
	baseClient := modules.NewBaseClient(cfg, encodingConfig, nil)
	keysClient := keys.NewClient(baseClient)

	bankClient := bank.NewClient(baseClient, encodingConfig.Marshaler)
	tokenClient := token.NewClient(baseClient, encodingConfig.Marshaler)
	swapClient := coinswap.NewClient(baseClient, encodingConfig.Marshaler, bankClient.TotalSupply)
	nftClient := nft.NewClient(baseClient, encodingConfig.Marshaler)
	govClient := gov.NewClient(baseClient, encodingConfig.Marshaler)
	stakingClient := staking.NewClient(baseClient, encodingConfig.Marshaler)

	client := &PLUGCHAINClient{
		logger:         baseClient.Logger(),
		moduleManager:  make(map[string]types.Module),
		encodingConfig: encodingConfig,
		BaseClient:     baseClient,
		Key:            keysClient,
		Bank:           bankClient,
		Token:          tokenClient,
		Swap:           swapClient,
		Nft:            nftClient,
		Gov:            govClient,
		Staking:        stakingClient,
	}

	client.RegisterModule(
		bankClient,
		tokenClient,
		stakingClient,
		govClient,
		nftClient,
		//randomClient,
		//oracleClient,
		//htlcClient,
		swapClient,
	)

	return *client
}

//Set log
func (client *PLUGCHAINClient) SetLogger(logger log.Logger) {
	client.BaseClient.SetLogger(logger)
}

func (client *PLUGCHAINClient) Codec() *codec.LegacyAmino {
	return client.encodingConfig.Amino
}

func (client *PLUGCHAINClient) AppCodec() codec.Marshaler {
	return client.encodingConfig.Marshaler
}

func (client *PLUGCHAINClient) EncodingConfig() types.EncodingConfig {
	return client.encodingConfig
}

func (client *PLUGCHAINClient) Manager() types.BaseClient {
	return client.BaseClient
}

func (client *PLUGCHAINClient) RegisterModule(ms ...types.Module) {
	for _, m := range ms {
		_, ok := client.moduleManager[m.Name()]
		if ok {
			panic(fmt.Sprintf("%s has register", m.Name()))
		}

		// m.RegisterCodec(client.encodingConfig.Amino)
		m.RegisterInterfaceTypes(client.encodingConfig.InterfaceRegistry)
		client.moduleManager[m.Name()] = m
	}
}

func (client *PLUGCHAINClient) Module(name string) types.Module {
	return client.moduleManager[name]
}

func makeEncodingConfig() types.EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := cdctypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := txtypes.NewTxConfig(marshaler, txtypes.DefaultSignModes)

	encodingConfig := types.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             amino,
	}
	RegisterLegacyAminoCodec(encodingConfig.Amino)
	RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}

// Register SDK message type
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*types.Msg)(nil), nil)
	cdc.RegisterInterface((*types.Tx)(nil), nil)
	cryptocodec.RegisterCrypto(cdc)
}

// Register interface register SDK message type
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterInterface("cosmos.v1beta1.Msg", (*types.Msg)(nil))

	txtypes.RegisterInterfaces(registry)
	cryptocodec.RegisterInterfaces(registry)
}
