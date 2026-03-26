package address

import (
	"encoding/hex"
	"regexp"
	"strings"

	"golang.org/x/crypto/sha3"
)

var ethAddressRegex = regexp.MustCompile(`^0x[0-9a-fA-F]{40}$`)
var ethLowerRegex = regexp.MustCompile(`^0x[0-9a-f]{40}$`)
var ethUpperRegex = regexp.MustCompile(`^0x[0-9A-F]{40}$`)

// ValidateEthereum validates an Ethereum address with EIP-55 checksum support.
func ValidateEthereum(address string) bool {
	// Check basic format
	if !ethAddressRegex.MatchString(address) {
		return false
	}

	// If all lowercase or all uppercase, valid (no checksum)
	if ethLowerRegex.MatchString(address) || ethUpperRegex.MatchString(address) {
		return true
	}

	// Verify EIP-55 checksum
	return verifyEthereumChecksum(address)
}

// verifyEthereumChecksum verifies the EIP-55 checksum of an Ethereum address.
func verifyEthereumChecksum(address string) bool {
	// Remove 0x prefix
	addr := strings.ToLower(address[2:])

	// Compute Keccak-256 hash of lowercase address
	hash := keccak256([]byte(addr))
	hashHex := hex.EncodeToString(hash)

	// Check each character
	for i := 0; i < 40; i++ {
		hashDigit, _ := hex.DecodeString(string(hashHex[i]) + "0")
		if len(hashDigit) == 0 {
			hashDigit = []byte{0}
		}
		hashVal := int(hashHex[i])
		if hashHex[i] >= 'a' && hashHex[i] <= 'f' {
			hashVal = int(hashHex[i] - 'a' + 10)
		} else if hashHex[i] >= '0' && hashHex[i] <= '9' {
			hashVal = int(hashHex[i] - '0')
		}

		addrChar := address[i+2]

		// If hash digit > 7, character should be uppercase
		// If hash digit <= 7, character should be lowercase
		if hashVal > 7 {
			if addrChar >= 'a' && addrChar <= 'f' {
				return false
			}
		} else {
			if addrChar >= 'A' && addrChar <= 'F' {
				return false
			}
		}
	}

	return true
}

// keccak256 computes the Keccak-256 hash of the input.
func keccak256(data []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	return hash.Sum(nil)
}

// keccak256Hex computes the Keccak-256 hash and returns it as a hex string.
func keccak256Hex(data []byte) string {
	return hex.EncodeToString(keccak256(data))
}

// keccak256Checksum returns the first 8 hex characters of a Keccak-256 hash.
func keccak256Checksum(data []byte) string {
	hash := keccak256(data)
	return hex.EncodeToString(hash[:4])
}
