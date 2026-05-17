package address

import (
	"encoding/base64"
	"strings"
)

// ValidateTON validates a TON (The Open Network) address.
// TON uses two formats:
// - Raw: workchain:hex_hash (e.g., 0:abc...def) — 66 chars
// - User-friendly: base64url encoded, 48 chars, starting with EQ/UQ (bounceable/non-bounceable)
func ValidateTON(address string) bool {
	if address == "" {
		return false
	}

	// Raw format: 0:64_hex_chars or -1:64_hex_chars
	if strings.Contains(address, ":") {
		return validateTONRaw(address)
	}

	// User-friendly base64 format
	return validateTONUserFriendly(address)
}

func validateTONRaw(address string) bool {
	parts := strings.SplitN(address, ":", 2)
	if len(parts) != 2 {
		return false
	}

	// Workchain must be "0" or "-1"
	if parts[0] != "0" && parts[0] != "-1" {
		return false
	}

	// Hash must be exactly 64 hex characters
	hash := parts[1]
	if len(hash) != 64 {
		return false
	}
	for _, c := range hash {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}

	return true
}

func validateTONUserFriendly(address string) bool {
	// Must be 48 characters (36 bytes base64-encoded)
	if len(address) != 48 {
		return false
	}

	// Replace URL-safe base64 characters
	b64 := strings.ReplaceAll(address, "-", "+")
	b64 = strings.ReplaceAll(b64, "_", "/")

	decoded, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return false
	}

	// Must decode to exactly 36 bytes: 1 tag + 1 workchain + 32 hash + 2 CRC
	if len(decoded) != 36 {
		return false
	}

	// Tag byte: 0x11 (bounceable) or 0x51 (non-bounceable), with optional +0x80 for testnet
	tag := decoded[0]
	validTags := tag == 0x11 || tag == 0x51 || tag == 0x91 || tag == 0xD1
	if !validTags {
		return false
	}

	return true
}
