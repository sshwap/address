package address

import (
	"regexp"
)

// NEAR supports two address formats:
// 1. Named accounts: alice.near, bob.testnet (lowercase alphanumeric, hyphens, dots, 2-64 chars)
// 2. Implicit accounts: 64-character hex string (ed25519 public key)

var nearNamedRegex = regexp.MustCompile(`^[a-z0-9]([a-z0-9._-]*[a-z0-9])?$`)
var nearImplicitRegex = regexp.MustCompile(`^[0-9a-f]{64}$`)

// ValidateNEAR validates a NEAR Protocol address.
func ValidateNEAR(address string) bool {
	if address == "" || len(address) < 2 || len(address) > 64 {
		return false
	}

	// Implicit account (64 hex chars)
	if len(address) == 64 && nearImplicitRegex.MatchString(address) {
		return true
	}

	// Named account
	if !nearNamedRegex.MatchString(address) {
		return false
	}

	// No consecutive dots/hyphens
	for i := 1; i < len(address); i++ {
		if (address[i] == '.' || address[i] == '-') && (address[i-1] == '.' || address[i-1] == '-') {
			return false
		}
	}

	return true
}
