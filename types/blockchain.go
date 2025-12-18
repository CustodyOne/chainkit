package types

import (
	"fmt"
	"slices"
	"strings"
)

type SignatureType string

const (
	K256Keccak = SignatureType("k256-keccak")
	K256Sha256 = SignatureType("k256-sha256")
	Ed255      = SignatureType("ed255")
	Schnorr    = SignatureType("schnorr")
)

// Protocol is the type of a chain
type Protocol string

// List of supported Protocol
const (
	// ProtocolAptos       = Protocol("aptos")
	ProtocolBtc         = Protocol("btc")
	ProtocolBtcCash     = Protocol("btc-cash")
	ProtocolBtcLegacy   = Protocol("btc-legacy")
	ProtocolCosmos      = Protocol("cosmos")
	ProtocolCosmosEvmos = Protocol("evmos")
	ProtocolEVM         = Protocol("evm")
	ProtocolEVMLegacy   = Protocol("evm-legacy")
	// ProtocolSubstrate   = Protocol("substrate")
	ProtocolSolana = Protocol("solana")
	// ProtocolSui         = Blockchain("sui")
	ProtocolTron = Protocol("tron")
	ProtocolTon  = Protocol("ton")
	// Chainkit is a client-only blockchain
	ProtocolChainkit = Protocol("chainkit")
)

var SupportedProtocols = []Protocol{
	// ProtocolAptos,
	ProtocolBtc,
	ProtocolBtcCash,
	ProtocolBtcLegacy,
	ProtocolCosmos,
	ProtocolCosmosEvmos,
	ProtocolEVM,
	ProtocolEVMLegacy,
	// ProtocolSubstrate,
	ProtocolSolana,
	// ProtocolSui,
	ProtocolTron,
	ProtocolTon,
}

type StakingProvider string

const Kiln StakingProvider = "kiln"
const Figment StakingProvider = "figment"
const Twinstake StakingProvider = "twinstake"
const Native StakingProvider = "native"

var SupportedStakingProviders = []StakingProvider{
	Native,
	Kiln,
	Figment,
	Twinstake,
}

func (stakingProvider StakingProvider) Valid() bool {
	return slices.Contains(SupportedStakingProviders, stakingProvider)
}

type TxVariantInputType string

func NewStakingInputType(blockchain Protocol, variant string) TxVariantInputType {
	return TxVariantInputType(fmt.Sprintf("protocols/%s/staking/%s", blockchain, variant))
}

func NewUnstakingInputType(blockchain Protocol, variant string) TxVariantInputType {
	return TxVariantInputType(fmt.Sprintf("protocols/%s/unstaking/%s", blockchain, variant))
}

func NewWithdrawingInputType(blockchain Protocol, variant string) TxVariantInputType {
	return TxVariantInputType(fmt.Sprintf("protocols/%s/withdrawing/%s", blockchain, variant))
}

func (variant TxVariantInputType) Blockchain() Protocol {
	return Protocol(strings.Split(string(variant), "/")[1])
}
func (variant TxVariantInputType) Variant() string {
	return (strings.Split(string(variant), "/")[3])
}

func (variant TxVariantInputType) Validate() error {
	if len(strings.Split(string(variant), "/")) != 4 {
		return fmt.Errorf("invalid input variant type: %s", variant)
	}
	return nil
}

func (native NativeAsset) IsValid() bool {
	return NativeAsset(native).Protocol() != ""
}

func (native NativeAsset) Protocol() Protocol {
	switch native {
	case BTC:
		return ProtocolBtc
	case BCH:
		return ProtocolBtcCash
	case DOGE, LTC:
		return ProtocolBtcLegacy
	case AVAX, CELO, ETH, ETHW, MATIC, OptETH, ArbETH, BERA:
		return ProtocolEVM
	case BNB, FTM, ETC, EmROSE, AurETH, ACA, KAR, KLAY, OAS, CHZ, XDC, CHZ2:
		return ProtocolEVMLegacy
	// case APTOS:
	//		return ProtocolAptos
	case ATOM, XPLA, INJ, HASH, LUNC, LUNA, SEI, TIA:
		return ProtocolCosmos
	//case SUI:
	//	return ProtocolSui
	case SOL:
		return ProtocolSolana
	// case DOT, TAO, KSM:
	//	return ProtocolSubstrate
	case TRX:
		return ProtocolTron
	case TON:
		return ProtocolTon
	}
	return ""
}

func (protocol Protocol) SignatureAlgorithm() SignatureType {
	switch protocol {
	case ProtocolBtc, ProtocolBtcCash, ProtocolBtcLegacy:
		return K256Sha256
	case ProtocolEVM, ProtocolEVMLegacy, ProtocolCosmos, ProtocolCosmosEvmos, ProtocolTron:
		return K256Keccak
	case /*ProtocolAptos,*/ ProtocolSolana /*ProtocolSui,*/, ProtocolTon /*, ProtocolSubstrate*/ :
		return Ed255
	}
	return ""
}

type PublicKeyFormat string

var Raw PublicKeyFormat = "raw"
var Compressed PublicKeyFormat = "compressed"
var Uncompressed PublicKeyFormat = "uncompressed"

func (protocol Protocol) PublicKeyFormat() PublicKeyFormat {
	switch protocol {
	case ProtocolBtc, ProtocolBtcCash, ProtocolBtcLegacy:
		return Compressed
	case ProtocolCosmos, ProtocolCosmosEvmos:
		return Compressed
	case ProtocolEVM, ProtocolEVMLegacy, ProtocolTron:
		return Uncompressed
	case /*ProtocolAptos, */ ProtocolSolana /*ProtocolSui, */, ProtocolTon /*, ProtocolSubstrate*/ :
		return Raw
	}
	return ""
}
