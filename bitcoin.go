package address

import (
	"crypto/sha256"
	"encoding/hex"
)

// Bitcoin address configuration
type bitcoinConfig struct {
	addressTypes []string // Hex-encoded address type bytes
	bech32Hrps   []string // Bech32 HRPs
	segwitVers   []int    // Allowed SegWit versions
}

var bitcoinMainnet = bitcoinConfig{
	addressTypes: []string{"00", "05"},
	bech32Hrps:   []string{"bc"},
	segwitVers:   []int{0, 1}, // v0 (Native SegWit) and v1 (Taproot)
}

var bitcoinTestnet = bitcoinConfig{
	addressTypes: []string{"6f", "c4", "3c", "26"},
	bech32Hrps:   []string{"tb"},
	segwitVers:   []int{0, 1},
}

// ValidateBitcoin validates a Bitcoin address.
func ValidateBitcoin(address string, network Network) bool {
	config := bitcoinMainnet
	if network == Testnet {
		config = bitcoinTestnet
	}

	// Try base58 first, then SegWit
	return validateBase58Address(address, config.addressTypes) ||
		ValidateSegwitAddress(address, config.bech32Hrps, config.segwitVers)
}

// validateBase58Address validates a base58check encoded address.
func validateBase58Address(address string, validTypes []string) bool {
	decoded, err := Base58Decode(address)
	if err != nil || len(decoded) != 25 {
		return false
	}

	// Extract checksum (last 4 bytes)
	payload := decoded[:21]
	checksum := decoded[21:]

	// Calculate expected checksum
	hash1 := sha256.Sum256(payload)
	hash2 := sha256.Sum256(hash1[:])
	expectedChecksum := hash2[:4]

	// Verify checksum
	for i := 0; i < 4; i++ {
		if checksum[i] != expectedChecksum[i] {
			return false
		}
	}

	// Verify address type
	addrType := hex.EncodeToString(decoded[:1])
	for _, validType := range validTypes {
		if addrType == validType {
			return true
		}
	}

	return false
}

// sha256Checksum computes the SHA256 double-hash checksum (first 8 hex chars).
func sha256Checksum(hexPayload string) string {
	payload, err := hex.DecodeString(hexPayload)
	if err != nil {
		return ""
	}
	hash1 := sha256.Sum256(payload)
	hash2 := sha256.Sum256(hash1[:])
	return hex.EncodeToString(hash2[:4])
}
