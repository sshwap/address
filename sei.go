package address

// ValidateSei validates a Sei (SEI) address.
// Sei native chain uses Cosmos bech32 with "sei" HRP and 20 or 32-byte payload.
// Sei EVM uses Ethereum format.
func ValidateSei(address string) bool {
	return validateBech32Address(address, "sei", 20) ||
		validateBech32Address(address, "sei", 32) ||
		ValidateEthereum(address)
}
