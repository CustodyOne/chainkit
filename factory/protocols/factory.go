package protocols

import (
	"errors"
	"fmt"

	"github.com/CustodyOne/chainkit/blockchain/btc"
	btcclient "github.com/CustodyOne/chainkit/blockchain/btc/client"
	"github.com/CustodyOne/chainkit/blockchain/btc_cash"
	cosmosbuilder "github.com/CustodyOne/chainkit/blockchain/cosmos/builder"
	cosmosclient "github.com/CustodyOne/chainkit/blockchain/cosmos/client"

	evm_legacy "github.com/CustodyOne/chainkit/blockchain/evm_legacy"
	solanabuilder "github.com/CustodyOne/chainkit/blockchain/solana/builder"

	evmbuilder "github.com/CustodyOne/chainkit/blockchain/evm/builder"
	evmclient "github.com/CustodyOne/chainkit/blockchain/evm/client"
	solanaclient "github.com/CustodyOne/chainkit/blockchain/solana/client"
	tonclient "github.com/CustodyOne/chainkit/blockchain/ton/client"
	tronclient "github.com/CustodyOne/chainkit/blockchain/tron/client"
	xcbuilder "github.com/CustodyOne/chainkit/builder"

	btcaddress "github.com/CustodyOne/chainkit/blockchain/btc/address"
	cosmosaddress "github.com/CustodyOne/chainkit/blockchain/cosmos/address"
	evmaddress "github.com/CustodyOne/chainkit/blockchain/evm/address"
	solanaaddress "github.com/CustodyOne/chainkit/blockchain/solana/address"
	tonaddress "github.com/CustodyOne/chainkit/blockchain/ton/address"
	"github.com/CustodyOne/chainkit/factory/signer"

	// "github.com/CustodyOne/chainkit/blockchain/aptos"
	// "github.com/CustodyOne/chainkit/chain/evm_legacy"

	// "github.com/CustodyOne/chainkit/blockchain/evm_legacy"
	// "github.com/openweb-io/chainkit/blockchain/substrate"
	// "github.com/CustodyOne/chainkit/blockchain/sui"
	"github.com/CustodyOne/chainkit/blockchain/ton"
	"github.com/CustodyOne/chainkit/blockchain/tron"
	xc_client "github.com/CustodyOne/chainkit/client"
	xc "github.com/CustodyOne/chainkit/types"
)

type ClientCreator func(cfg *xc.ChainConfig) (xc_client.IClient, error)

var (
	creatorMap = make(map[xc.Protocol]ClientCreator)
)

func RegisterClient(cfg xc.Protocol, creator ClientCreator) {
	creatorMap[cfg] = creator
}

func init() {
	RegisterClient(xc.ProtocolBtc, func(cfg *xc.ChainConfig) (xc_client.IClient, error) {
		return btcclient.NewClient(cfg)
	})

	RegisterClient(xc.ProtocolBtcLegacy, func(cfg *xc.ChainConfig) (xc_client.IClient, error) {
		return btcclient.NewClient(cfg)
	})

	RegisterClient(xc.ProtocolBtcCash, func(cfg *xc.ChainConfig) (xc_client.IClient, error) {
		return btc_cash.NewClient(cfg)
	})

	RegisterClient(xc.ProtocolCosmos, func(cfg *xc.ChainConfig) (xc_client.IClient, error) {
		return cosmosclient.NewClient(cfg)
	})

	RegisterClient(xc.ProtocolCosmosEvmos, func(cfg *xc.ChainConfig) (xc_client.IClient, error) {
		return cosmosclient.NewClient(cfg)
	})

	RegisterClient(xc.ProtocolTon, func(cfg *xc.ChainConfig) (xc_client.IClient, error) {
		return tonclient.NewClient(cfg)
	})

	RegisterClient(xc.ProtocolTron, func(cfg *xc.ChainConfig) (xc_client.IClient, error) {
		return tronclient.NewClient(cfg)
	})

	RegisterClient(xc.ProtocolSolana, func(cfg *xc.ChainConfig) (xc_client.IClient, error) {
		return solanaclient.NewClient(cfg)
	})

	RegisterClient(xc.ProtocolEVM, func(cfg *xc.ChainConfig) (xc_client.IClient, error) {
		return evmclient.NewClient(cfg)
	})

	RegisterClient(xc.ProtocolEVMLegacy, func(cfg *xc.ChainConfig) (xc_client.IClient, error) {
		return evm_legacy.NewClient(cfg)
	})
}

func NewClient(cfg *xc.ChainConfig, blockchain xc.Protocol) (xc_client.IClient, error) {
	creator, ok := creatorMap[blockchain]
	if !ok {
		return nil, fmt.Errorf("creator %s not found", cfg.Protocol)
	}

	return creator(cfg)
}

func NewAddressBuilder(cfg *xc.ChainConfig) (xc.AddressBuilder, error) {
	switch xc.Protocol(cfg.Protocol) {
	case xc.ProtocolEVM:
		return evmaddress.NewAddressBuilder(cfg)
	//case types.ProtocolEVMLegacy:
	//	return evm_legacy.NewAddressBuilder(cfg)
	case xc.ProtocolCosmos, xc.ProtocolCosmosEvmos:
		return cosmosaddress.NewAddressBuilder(cfg)
	case xc.ProtocolSolana:
		return solanaaddress.NewAddressBuilder(cfg)
	//case types.BlockchainAptos:
	//	return aptos.NewAddressBuilder(cfg)
	case xc.ProtocolBtc, xc.ProtocolBtcLegacy:
		return btcaddress.NewAddressBuilder(cfg)
	case xc.ProtocolBtcCash:
		return btc_cash.NewAddressBuilder(cfg)
	// case types.BlockchainSui:
	// 	return sui.NewAddressBuilder(cfg)
	//case types.BlockchainSubstrate:
	//	return substrate.NewAddressBuilder(cfg)
	case xc.ProtocolTron:
		return tron.NewAddressBuilder(cfg)
	case xc.ProtocolTon:
		return tonaddress.NewAddressBuilder(cfg)
	}
	return nil, errors.New("no address builder defined for: " + string(cfg.ID()))
}

func NewSigner(cfg *xc.ChainConfig, secret string) (*signer.Signer, error) {
	return signer.New(cfg.Protocol, secret, cfg)
}

func NewTxBuilder(cfg *xc.ChainConfig) (xcbuilder.TxBuilder, error) {
	switch xc.Protocol(cfg.Protocol) {
	case xc.ProtocolEVM:
		return evmbuilder.NewTxBuilder(cfg)
	//case ProtocolEVMLegacy:
	//	return evm_legacy.NewTxBuilder(cfg)
	case xc.ProtocolCosmos, xc.ProtocolCosmosEvmos:
		return cosmosbuilder.NewTxBuilder(cfg)
	case xc.ProtocolSolana:
		return solanabuilder.NewTxBuilder(cfg)
	//case ProtocolAptos:
	//	return aptos.NewTxBuilder(cfg)
	//case BlockchainSui:
	//	return sui.NewTxBuilder(cfg)
	case xc.ProtocolBtc, xc.ProtocolBtcLegacy:
		return btc.NewTxBuilder(cfg)
	case xc.ProtocolBtcCash:
		return btc_cash.NewTxBuilder(cfg)
	// case BlockchainSubstrate:
	//	return substrate.NewTxBuilder(cfg)
	case xc.ProtocolTron:
		return tron.NewTxBuilder(cfg)
	case xc.ProtocolTon:
		return ton.NewTxBuilder(cfg)
	}
	return nil, errors.New("no tx-builder defined for: " + string(cfg.ID()))
}
