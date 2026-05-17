package address

import "strings"

// ValidateCosmos validates a Cosmos (ATOM) address.
// Cosmos uses bech32 encoding with "cosmos" HRP for mainnet.
// Addresses are 20 bytes (from a SHA-256 + RIPEMD-160 hash of the public key).
func ValidateCosmos(address string) bool {
	return validateBech32Address(address, "cosmos", 20)
}

// ValidateCosmosCompatible validates addresses for any Cosmos SDK chain
// by checking bech32 with a specific HRP and expected payload size.
func validateBech32Address(address, hrp string, expectedPayloadLen int) bool {
	if address == "" {
		return false
	}

	lowerAddr := strings.ToLower(address)
	if !strings.HasPrefix(lowerAddr, hrp+"1") {
		return false
	}

	decHrp, data, _, _ := Bech32Decode(address)
	if decHrp != hrp || len(data) < 1 {
		return false
	}

	// Convert from 5-bit to 8-bit
	decoded := convertBits(data, 5, 8, false)
	if decoded == nil || len(decoded) != expectedPayloadLen {
		return false
	}

	return true
}
