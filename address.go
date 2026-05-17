// Package address provides cryptocurrency address validation.
package address

import (
	"strings"
)

// Coin represents a cryptocurrency type.
type Coin string

const (
	BTC  Coin = "BTC"
	ETH  Coin = "ETH"
	USDT Coin = "USDT"
	USDC Coin = "USDC"
	XRP  Coin = "XRP"
	SOL  Coin = "SOL"
	BCH  Coin = "BCH"
	XMR  Coin = "XMR"
	LTC  Coin = "LTC"
	TRX  Coin = "TRX"
	FIRO Coin = "FIRO"
	ZANO Coin = "ZANO"
	BNB  Coin = "BNB"
	ADA  Coin = "ADA"
	DOGE Coin = "DOGE"
	LINK Coin = "LINK"
	TON  Coin = "TON"
	AVAX Coin = "AVAX"
	SUI  Coin = "SUI"
	DOT  Coin = "DOT"
	XLM  Coin = "XLM"
	HBAR Coin = "HBAR"
	NEAR Coin = "NEAR"
	ATOM Coin = "ATOM"
	ETC  Coin = "ETC"
	APT  Coin = "APT"
	KAS  Coin = "KAS"
	FIL  Coin = "FIL"
	VET  Coin = "VET"
	ARB  Coin = "ARB"
	OP   Coin = "OP"
	ALGO Coin = "ALGO"
	STX  Coin = "STX"
	INJ  Coin = "INJ"
	SEI  Coin = "SEI"
)

// Network represents a blockchain network type.
type Network string

const (
	Mainnet Network = "mainnet"
	Testnet Network = "testnet"
)

// Options contains validation options.
type Options struct {
	Network Network // mainnet or testnet (default: mainnet)
}

// DefaultOptions returns default validation options.
func DefaultOptions() Options {
	return Options{
		Network: Mainnet,
	}
}

// Validate checks if an address is valid for the given coin.
func Validate(coin Coin, address string, opts ...Options) bool {
	if address == "" {
		return false
	}

	opt := DefaultOptions()
	if len(opts) > 0 {
		opt = opts[0]
	}

	switch coin {
	case BTC:
		return ValidateBitcoin(address, opt.Network)
	case ETH, LINK, ETC, ARB, OP:
		return ValidateEthereum(address)
	case BNB:
		return ValidateEthereum(address) // BSC uses ETH format
	case USDT, USDC:
		return ValidateEthereum(address) || ValidateTron(address) || ValidateSolana(address)
	case XRP:
		return ValidateRipple(address)
	case SOL:
		return ValidateSolana(address)
	case BCH:
		return ValidateBitcoinCash(address, opt.Network)
	case XMR:
		return ValidateMonero(address, opt.Network)
	case LTC:
		return ValidateLitecoin(address, opt.Network)
	case TRX:
		return ValidateTron(address)
	case FIRO:
		return ValidateFiro(address, opt.Network)
	case ZANO:
		return ValidateZano(address)
	case ADA:
		return ValidateCardano(address, opt.Network)
	case DOGE:
		return ValidateDogecoin(address)
	case TON:
		return ValidateTON(address)
	case AVAX:
		return ValidateEthereum(address) // C-Chain uses ETH format
	case SUI:
		return ValidateSUI(address)
	case DOT:
		return ValidatePolkadot(address)
	case XLM:
		return ValidateStellar(address)
	case HBAR:
		return ValidateHedera(address)
	case NEAR:
		return ValidateNEAR(address)
	case ATOM:
		return ValidateCosmos(address)
	case APT:
		return ValidateAptos(address)
	case KAS:
		return ValidateKaspa(address)
	case FIL:
		return ValidateFilecoin(address)
	case VET:
		return ValidateVeChain(address)
	case ALGO:
		return ValidateAlgorand(address)
	case STX:
		return ValidateStacks(address, opt.Network)
	case INJ:
		return ValidateInjective(address)
	case SEI:
		return ValidateSei(address)
	default:
		return false
	}
}

// ValidateForNetwork validates an address for a specific coin and network string.
func ValidateForNetwork(coin Coin, network, address string) bool {
	if address == "" {
		return false
	}

	network = strings.ToUpper(network)

	// Multi-network stablecoins
	if coin == USDT || coin == USDC {
		switch network {
		case "ETH", "ERC20", "ERC-20", "ETHEREUM", "ARBITRUM", "OP", "BASE", "ZKSYNC", "MATIC", "AVAXC", "CELO", "MNT":
			return ValidateEthereum(address)
		case "TRX", "TRC20", "TRC-20", "TRON":
			return ValidateTron(address)
		case "SOL", "SOLANA", "SPL":
			return ValidateSolana(address)
		case "BSC", "BEP20", "BEP-20":
			return ValidateEthereum(address)
		case "ALGO":
			return ValidateAlgorand(address)
		case "APT":
			return ValidateAptos(address)
		case "SUI":
			return ValidateSUI(address)
		default:
			return ValidateEthereum(address) || ValidateTron(address) || ValidateSolana(address)
		}
	}

	// Coins available on BSC as wrapped tokens (ETH format)
	if network == "BSC" || network == "BEP20" || network == "BEP-20" {
		return ValidateEthereum(address)
	}

	// ETH L2 networks
	if network == "ARBITRUM" || network == "OP" || network == "BASE" || network == "ZKSYNC" || network == "MATIC" || network == "AVAXC" {
		return ValidateEthereum(address)
	}

	// BNB on OPBNB
	if coin == BNB && network == "OPBNB" {
		return ValidateEthereum(address)
	}

	// AVAX X-Chain (different format) - for now just accept C-Chain
	if coin == AVAX && network == "CCHAIN" {
		return ValidateEthereum(address)
	}

	// SEI EVM
	if coin == SEI && network == "SEIEVM" {
		return ValidateEthereum(address)
	}

	// DOT on Asset Hub
	if coin == DOT && network == "ASSETHUB" {
		return ValidatePolkadot(address)
	}

	// For other coins, use standard validation
	return Validate(coin, address)
}

// ValidateMemo validates a memo/tag for coins that support it.
func ValidateMemo(coin Coin, memo string) bool {
	switch coin {
	case XRP:
		return ValidateRippleMemo(memo)
	case XLM:
		return ValidateStellarMemo(memo)
	case HBAR:
		return ValidateHederaMemo(memo)
	case ATOM, INJ:
		return len(memo) <= 256 // Cosmos SDK memo limit
	case TON:
		return len(memo) <= 512
	case STX:
		return len(memo) <= 34
	default:
		return true
	}
}
