package address

import (
	"regexp"
	"strings"
)

// Bitcoin Cash CashAddr format
var bchCashAddrRegex = regexp.MustCompile(`^[qQpP][0-9a-zA-Z]{41}$`)

// Base32 alphabet for CashAddr (different from Bech32)
const cashAddrCharset = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

// ValidateBitcoinCash validates a Bitcoin Cash address.
// Supports both CashAddr format and legacy Bitcoin format.
func ValidateBitcoinCash(address string, network Network) bool {
	// Try CashAddr format first
	if validateCashAddr(address, network) {
		return true
	}

	// Fall back to legacy Bitcoin format
	config := bitcoinMainnet
	if network == Testnet {
		config = bitcoinTestnet
	}

	return validateBase58Address(address, config.addressTypes)
}

// validateCashAddr validates a Bitcoin Cash CashAddr format address.
func validateCashAddr(address string, network Network) bool {
	var rawAddr string
	var prefix string

	// Check for prefix
	parts := strings.Split(address, ":")
	if len(parts) == 1 {
		rawAddr = address
		if network == Mainnet {
			prefix = "bitcoincash"
		} else {
			prefix = "bchtest"
		}
	} else if len(parts) == 2 {
		if network == Mainnet && parts[0] != "bitcoincash" {
			return false
		}
		if network == Testnet && parts[0] != "bchtest" {
			return false
		}
		prefix = parts[0]
		rawAddr = parts[1]
	} else {
		return false
	}

	// Check regex pattern
	if !bchCashAddrRegex.MatchString(rawAddr) {
		return false
	}

	// Check for mixed case
	if strings.ToLower(rawAddr) != rawAddr && strings.ToUpper(rawAddr) != rawAddr {
		return false
	}

	rawAddr = strings.ToLower(rawAddr)

	// Decode and verify checksum
	data := make([]int, len(rawAddr))
	for i, c := range rawAddr {
		idx := strings.IndexByte(cashAddrCharset, byte(c))
		if idx == -1 {
			return false
		}
		data[i] = idx
	}

	// Verify polymod checksum
	return verifyCashAddrChecksum(prefix, data)
}

// verifyCashAddrChecksum verifies the CashAddr polymod checksum.
func verifyCashAddrChecksum(prefix string, data []int) bool {
	// Expand prefix
	prefixData := make([]int, len(prefix)+1)
	for i, c := range prefix {
		prefixData[i] = int(c) & 0x1f
	}
	prefixData[len(prefix)] = 0

	// Combine with data
	values := append(prefixData, data...)

	// Compute polymod
	return cashAddrPolymod(values) == 0
}

// cashAddrPolymod computes the CashAddr polymod.
func cashAddrPolymod(values []int) uint64 {
	c := uint64(1)
	for _, d := range values {
		c0 := c >> 35
		c = ((c & 0x07ffffffff) << 5) ^ uint64(d)
		if c0&0x01 != 0 {
			c ^= 0x98f2bc8e61
		}
		if c0&0x02 != 0 {
			c ^= 0x79b76d99e2
		}
		if c0&0x04 != 0 {
			c ^= 0xf33e5fb3c4
		}
		if c0&0x08 != 0 {
			c ^= 0xae2eabe2a8
		}
		if c0&0x10 != 0 {
			c ^= 0x1e4f43e470
		}
	}
	return c ^ 1
}
