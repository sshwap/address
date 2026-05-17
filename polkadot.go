package address

import (
	"golang.org/x/crypto/blake2b"
)

// SS58 prefix for checksum
var ss58Prefix = []byte("SS58PRE")

// ValidatePolkadot validates a Polkadot (DOT) address.
// Polkadot uses SS58 encoding (modified base58) with network prefix 0.
func ValidatePolkadot(address string) bool {
	return validateSS58(address, []byte{0}) // Polkadot mainnet prefix
}

// validateSS58 validates an SS58-encoded address for given network prefix(es).
func validateSS58(address string, validPrefixes []byte) bool {
	if address == "" || len(address) < 10 {
		return false
	}

	decoded, err := Base58Decode(address)
	if err != nil {
		return false
	}

	// SS58 format: prefix(1-2 bytes) + pubkey(32 bytes) + checksum(2 bytes)
	// Simple accounts: 1 byte prefix + 32 bytes key + 2 bytes checksum = 35 bytes
	if len(decoded) != 35 {
		return false
	}

	prefix := decoded[0]
	validPrefix := false
	for _, vp := range validPrefixes {
		if prefix == vp {
			validPrefix = true
			break
		}
	}
	if !validPrefix {
		return false
	}

	// Verify checksum: blake2b-512 of (SS58PRE + prefix + pubkey), first 2 bytes
	payload := append(ss58Prefix, decoded[:33]...)
	hash := blake2b.Sum512(payload)
	return decoded[33] == hash[0] && decoded[34] == hash[1]
}
