package address

import "regexp"

// SUI uses 0x-prefixed 64-character hex addresses (32 bytes), like Aptos.
var suiRegex = regexp.MustCompile(`^0x[0-9a-fA-F]{64}$`)

// ValidateSUI validates a SUI address.
func ValidateSUI(address string) bool {
	return suiRegex.MatchString(address)
}
