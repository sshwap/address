package address

// ValidateInjective validates an Injective (INJ) address.
// Injective is a Cosmos SDK chain with "inj" HRP and 20-byte payload.
func ValidateInjective(address string) bool {
	return validateBech32Address(address, "inj", 20)
}
