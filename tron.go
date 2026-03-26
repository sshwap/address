package address

import (
	"crypto/sha256"
)

// ValidateTron validates a Tron address.
// Tron addresses are base58check encoded with a 0x41 prefix.
func ValidateTron(address string) bool {
	// Tron addresses are 34 characters
	if len(address) != 34 {
		return false
	}

	// Decode base58
	decoded, err := Base58Decode(address)
	if err != nil {
		return false
	}

	// Must be 25 bytes (21 byte address + 4 byte checksum)
	if len(decoded) != 25 {
		return false
	}

	// Verify checksum (last 4 bytes of double SHA256)
	payload := decoded[:21]
	checksum := decoded[21:]

	hash1 := sha256.Sum256(payload)
	hash2 := sha256.Sum256(hash1[:])
	expectedChecksum := hash2[:4]

	for i := 0; i < 4; i++ {
		if checksum[i] != expectedChecksum[i] {
			return false
		}
	}

	// First byte must be 0x41 (65 decimal) for mainnet
	return decoded[0] == 0x41
}
