package address

import (
	"regexp"
	"strings"
)

// Kaspa addresses use a bech32-like format with "kaspa:" prefix.
// P2PKH: kaspa:q... (66 chars after prefix)
// P2SH: kaspa:p... (66 chars after prefix)
// All lowercase, bech32 charset.
var kaspaRegex = regexp.MustCompile(`^kaspa:[qp][` + regexp.QuoteMeta(bech32Charset) + `]{60,64}$`)

// ValidateKaspa validates a Kaspa (KAS) address.
func ValidateKaspa(address string) bool {
	if address == "" {
		return false
	}

	lowerAddr := strings.ToLower(address)
	return kaspaRegex.MatchString(lowerAddr)
}
