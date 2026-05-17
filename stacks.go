package address

import "regexp"

// Stacks uses c32check encoding (not standard base58).
// Mainnet addresses start with "SP", testnet with "ST".
// C32 alphabet: 0123456789ABCDEFGHJKMNPQRSTVWXYZ
// Addresses are 41-50 characters.

var stacksMainnetRegex = regexp.MustCompile(`^SP[0-9A-HJ-NP-TV-Z]{20,48}$`)
var stacksTestnetRegex = regexp.MustCompile(`^ST[0-9A-HJ-NP-TV-Z]{20,48}$`)

// ValidateStacks validates a Stacks (STX) address.
func ValidateStacks(address string, network Network) bool {
	if address == "" {
		return false
	}

	if network == Testnet {
		return stacksTestnetRegex.MatchString(address)
	}
	return stacksMainnetRegex.MatchString(address)
}
