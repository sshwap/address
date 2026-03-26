package address

import (
	"crypto/sha256"
	"encoding/hex"
	"regexp"
	"strconv"
)

const rippleAlphabetChars = "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz"

var rippleAddressRegex = regexp.MustCompile(`^r[` + rippleAlphabetChars + `]{27,35}$`)
var rippleMemoRegex = regexp.MustCompile(`^[0-9]+$`)

// ValidateRipple validates a Ripple (XRP) address.
func ValidateRipple(address string) bool {
	// Check regex pattern
	if !rippleAddressRegex.MatchString(address) {
		return false
	}

	// Verify checksum
	return verifyRippleChecksum(address)
}

// verifyRippleChecksum verifies the checksum of a Ripple address.
func verifyRippleChecksum(address string) bool {
	decoded, err := Base58DecodeRipple(address)
	if err != nil || len(decoded) < 5 {
		return false
	}

	// Extract checksum (last 4 bytes)
	payload := decoded[:len(decoded)-4]
	checksum := decoded[len(decoded)-4:]

	// Compute expected checksum (double SHA256)
	hash1 := sha256.Sum256(payload)
	hash2 := sha256.Sum256(hash1[:])
	expectedChecksum := hash2[:4]

	// Compare checksums
	for i := 0; i < 4; i++ {
		if checksum[i] != expectedChecksum[i] {
			return false
		}
	}

	return true
}

// ValidateRippleMemo validates a Ripple destination tag (memo).
// Ripple memos are 32-bit unsigned integers (0 to 4,294,967,295).
func ValidateRippleMemo(memo string) bool {
	if memo == "" {
		return true // Memos are optional
	}

	// Check format
	if !rippleMemoRegex.MatchString(memo) {
		return false
	}

	// Parse and verify range
	val, err := strconv.ParseUint(memo, 10, 64)
	if err != nil {
		return false
	}

	// Must be 32-bit unsigned integer
	return val <= 4294967295
}

// rippleSha256Checksum computes the SHA256 double-hash checksum for Ripple.
func rippleSha256Checksum(payload []byte) string {
	hash1 := sha256.Sum256(payload)
	hash2 := sha256.Sum256(hash1[:])
	return hex.EncodeToString(hash2[:4])
}
