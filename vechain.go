package address

// ValidateVeChain validates a VeChain (VET) address.
// VeChain uses the same address format as Ethereum (0x + 40 hex).
func ValidateVeChain(address string) bool {
	return ValidateEthereum(address)
}
