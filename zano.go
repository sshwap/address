package address

import (
	"encoding/hex"
	"regexp"
)

// ZANO uses CryptoNote-style addresses similar to Monero.
// Standard addresses are 97 characters, integrated addresses are 108 characters.
var zanoStandardRegex = regexp.MustCompile(`^[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{97}$`)
var zanoIntegratedRegex = regexp.MustCompile(`^[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{108}$`)

// ZANO address type prefixes (mainnet only for now)
// iZ prefix addresses use 0xc1 (193)
// aZ prefix addresses use 0xc2 (194)
var zanoAddressTypes = []int{0xc1, 0xc2}
var zanoIntegratedTypes = []int{0xc3} // Integrated payment ID addresses

// ValidateZano validates a ZANO address.
func ValidateZano(address string) bool {
	isStandard := zanoStandardRegex.MatchString(address)
	isIntegrated := zanoIntegratedRegex.MatchString(address)

	if !isStandard && !isIntegrated {
		return false
	}

	// Decode using CryptoNote base58
	decoded := cnBase58Decode(address)
	if decoded == "" {
		return false
	}

	// Validate address type
	addrType, err := hex.DecodeString(decoded[:2])
	if err != nil || len(addrType) == 0 {
		return false
	}

	var validTypes []int
	if isIntegrated {
		validTypes = zanoIntegratedTypes
	} else {
		validTypes = zanoAddressTypes
	}

	typeValid := false
	for _, t := range validTypes {
		if int(addrType[0]) == t {
			typeValid = true
			break
		}
	}
	if !typeValid {
		return false
	}

	// Verify checksum (last 8 hex chars = 4 bytes)
	if len(decoded) < 8 {
		return false
	}

	addrChecksum := decoded[len(decoded)-8:]
	payload := decoded[:len(decoded)-8]

	// Convert payload hex to bytes
	payloadBytes, err := hex.DecodeString(payload)
	if err != nil {
		return false
	}

	// Compute Keccak-256 checksum
	hashChecksum := keccak256Checksum(payloadBytes)

	return addrChecksum == hashChecksum
}
