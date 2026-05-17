package address

import (
	"crypto/sha512"
	"encoding/base32"
	"strings"
)

// ValidateAlgorand validates an Algorand (ALGO) address.
// Algorand addresses are 58-character base32 strings encoding a 32-byte public key + 4-byte checksum.
func ValidateAlgorand(address string) bool {
	if address == "" || len(address) != 58 {
		return false
	}

	// Algorand uses base32 without padding
	padded := address
	for len(padded)%8 != 0 {
		padded += "="
	}

	decoded, err := base32.StdEncoding.DecodeString(strings.ToUpper(padded))
	if err != nil {
		return false
	}

	// Must decode to 36 bytes: 32-byte public key + 4-byte checksum
	if len(decoded) != 36 {
		return false
	}

	// Verify checksum: last 4 bytes of SHA-512/256 of the public key
	pubkey := decoded[:32]
	checksum := decoded[32:36]

	hash := sha512.Sum512_256(pubkey)
	expectedChecksum := hash[28:32]

	for i := 0; i < 4; i++ {
		if checksum[i] != expectedChecksum[i] {
			return false
		}
	}

	return true
}
