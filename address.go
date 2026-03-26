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
	XRP  Coin = "XRP"
	SOL  Coin = "SOL"
	BCH  Coin = "BCH"
	XMR  Coin = "XMR"
	LTC  Coin = "LTC"
	TRX  Coin = "TRX"
	FIRO Coin = "FIRO"
	ZANO Coin = "ZANO"
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
	case ETH:
		return ValidateEthereum(address)
	case USDT:
		// USDT can be on multiple networks; check common ones
		return ValidateEthereum(address) || ValidateTron(address)
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
	default:
		return false
	}
}

// ValidateForNetwork validates an address for a specific coin and network string.
// The network parameter should match common network identifiers like "ETH", "TRX", "ERC20", "TRC20", etc.
func ValidateForNetwork(coin Coin, network, address string) bool {
	if address == "" {
		return false
	}

	network = strings.ToUpper(network)

	// Handle USDT with specific network
	if coin == USDT {
		switch network {
		case "ETH", "ERC20", "ERC-20", "ETHEREUM":
			return ValidateEthereum(address)
		case "TRX", "TRC20", "TRC-20", "TRON":
			return ValidateTron(address)
		case "SOL", "SOLANA", "SPL":
			return ValidateSolana(address)
		case "BSC", "BEP20", "BEP-20":
			return ValidateEthereum(address) // BSC uses same format as ETH
		default:
			// Try both ETH and TRX formats
			return ValidateEthereum(address) || ValidateTron(address)
		}
	}

	// For other coins, use standard validation
	return Validate(coin, address)
}

// ValidateMemo validates a memo/tag for coins that support it.
func ValidateMemo(coin Coin, memo string) bool {
	switch coin {
	case XRP:
		return ValidateRippleMemo(memo)
	default:
		// Most coins don't require memo validation
		return true
	}
}
