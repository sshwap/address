package address

import "regexp"

// Aptos uses 0x-prefixed 64-character hex addresses (32 bytes).
// Some addresses may have leading zeros stripped (shorter than 64 hex chars).
var aptosRegex = regexp.MustCompile(`^0x[0-9a-fA-F]{1,64}$`)

// ValidateAptos validates an Aptos (APT) address.
func ValidateAptos(address string) bool {
	return aptosRegex.MatchString(address)
}
