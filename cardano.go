package address

import (
	"regexp"
	"strings"
)

// Cardano Shelley address: "addr1" + bech32 data (typically 98 chars total = 98-115 chars)
var cardanoShelleyRegex = regexp.MustCompile(`^addr1[` + bech32Charset + `]{50,110}$`)
var cardanoTestRegex = regexp.MustCompile(`^addr_test1[` + bech32Charset + `]{50,110}$`)

// ValidateCardano validates a Cardano (ADA) address.
// Shelley-era: bech32 with "addr1" prefix (mainnet) or "addr_test1" (testnet).
// Byron-era: base58-encoded, starts with "Ae2" or "DdzFF".
func ValidateCardano(address string, network Network) bool {
	if address == "" {
		return false
	}

	lowerAddr := strings.ToLower(address)

	if network == Testnet {
		if cardanoTestRegex.MatchString(lowerAddr) {
			return true
		}
	} else {
		if cardanoShelleyRegex.MatchString(lowerAddr) {
			return true
		}
	}

	// Byron-era base58 addresses
	if strings.HasPrefix(address, "Ae2") || strings.HasPrefix(address, "DdzFF") {
		decoded, err := Base58Decode(address)
		if err != nil {
			return false
		}
		return len(decoded) > 20
	}

	return false
}
