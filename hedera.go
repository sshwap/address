package address

import (
	"regexp"
	"strconv"
	"strings"
)

// hederaRegex matches Hedera account IDs: shard.realm.account (e.g., 0.0.12345)
var hederaRegex = regexp.MustCompile(`^\d+\.\d+\.\d+$`)

// ValidateHedera validates a Hedera Hashgraph (HBAR) address.
// Hedera uses account IDs in the format: shard.realm.account (e.g., 0.0.12345).
func ValidateHedera(address string) bool {
	if address == "" {
		return false
	}

	if !hederaRegex.MatchString(address) {
		return false
	}

	parts := strings.Split(address, ".")
	if len(parts) != 3 {
		return false
	}

	// Validate each part is a valid non-negative integer
	for _, part := range parts {
		n, err := strconv.ParseUint(part, 10, 64)
		if err != nil {
			return false
		}
		// Shard and realm are typically 0, account must be positive
		if part == parts[2] && n == 0 {
			return false
		}
		_ = n
	}

	return true
}

// ValidateHederaMemo validates a Hedera memo.
func ValidateHederaMemo(memo string) bool {
	if memo == "" {
		return true
	}
	return len(memo) <= 100
}
