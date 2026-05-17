package address

import (
	"regexp"
	"strconv"
)

// Filecoin address formats:
// f0... - ID address (numeric)
// f1... - secp256k1 (base32-lower, 41 chars total)
// f2... - actor (base32-lower, 41 chars total)
// f3... - BLS (base32-lower, 86 chars total)
// f4... - delegated (variable, for FVM/EVM)

var filF0Regex = regexp.MustCompile(`^f0[1-9]\d*$`)
var filF1F2Regex = regexp.MustCompile(`^f[12][a-z2-7]{38,39}$`)
var filF3Regex = regexp.MustCompile(`^f3[a-z2-7]{83,84}$`)
var filF4Regex = regexp.MustCompile(`^f4[1-9]\d*f[a-z2-7]+$`)

// ValidateFilecoin validates a Filecoin (FIL) address.
func ValidateFilecoin(address string) bool {
	if address == "" || len(address) < 3 {
		return false
	}

	if filF0Regex.MatchString(address) {
		id := address[2:]
		_, err := strconv.ParseUint(id, 10, 64)
		return err == nil
	}

	if filF1F2Regex.MatchString(address) {
		return true
	}

	if filF3Regex.MatchString(address) {
		return true
	}

	if filF4Regex.MatchString(address) {
		return true
	}

	return false
}
