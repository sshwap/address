package address

// ValidateSolana validates a Solana address.
// Solana addresses are base58 encoded 32-byte public keys.
func ValidateSolana(address string) bool {
	// Check length (43-44 characters when base58 encoded)
	if len(address) < 32 || len(address) > 44 {
		return false
	}

	// Decode and verify length
	decoded, err := Base58Decode(address)
	if err != nil {
		return false
	}

	// Solana addresses must be exactly 32 bytes
	return len(decoded) == 32
}
