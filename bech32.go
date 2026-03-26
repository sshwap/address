package address

import (
	"strings"
)

// Bech32 encoding constants
const bech32Charset = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

var bech32Generator = []uint32{0x3b6a57b2, 0x26508e6d, 0x1ea119fa, 0x3d4233dd, 0x2a1462b3}

// Bech32 encoding types
const (
	bech32Encoding  = 1
	bech32mEncoding = 0x2bc830a3
)

// bech32Polymod computes the Bech32 checksum polynomial.
func bech32Polymod(values []int) uint32 {
	chk := uint32(1)
	for _, v := range values {
		top := chk >> 25
		chk = (chk&0x1ffffff)<<5 ^ uint32(v)
		for i := 0; i < 5; i++ {
			if (top>>i)&1 == 1 {
				chk ^= bech32Generator[i]
			}
		}
	}
	return chk
}

// bech32HrpExpand expands the HRP for checksum computation.
func bech32HrpExpand(hrp string) []int {
	ret := make([]int, 0, len(hrp)*2+1)
	for i := 0; i < len(hrp); i++ {
		ret = append(ret, int(hrp[i])>>5)
	}
	ret = append(ret, 0)
	for i := 0; i < len(hrp); i++ {
		ret = append(ret, int(hrp[i])&31)
	}
	return ret
}

// bech32VerifyChecksum verifies a Bech32 checksum.
func bech32VerifyChecksum(hrp string, data []int, encoding uint32) bool {
	values := append(bech32HrpExpand(hrp), data...)
	return bech32Polymod(values) == encoding
}

// bech32CreateChecksum creates a Bech32 checksum.
func bech32CreateChecksum(hrp string, data []int, encoding uint32) []int {
	values := append(bech32HrpExpand(hrp), data...)
	values = append(values, 0, 0, 0, 0, 0, 0)
	mod := bech32Polymod(values) ^ encoding
	ret := make([]int, 6)
	for i := 0; i < 6; i++ {
		ret[i] = int((mod >> (5 * (5 - i))) & 31)
	}
	return ret
}

// Bech32Decode decodes a Bech32 string.
func Bech32Decode(bechString string) (hrp string, data []int, encoding uint32, err error) {
	// Check for mixed case
	hasLower := false
	hasUpper := false
	for i := 0; i < len(bechString); i++ {
		c := bechString[i]
		if c >= 'a' && c <= 'z' {
			hasLower = true
		}
		if c >= 'A' && c <= 'Z' {
			hasUpper = true
		}
	}
	if hasLower && hasUpper {
		return "", nil, 0, nil
	}

	bechString = strings.ToLower(bechString)

	// Find separator
	pos := strings.LastIndex(bechString, "1")
	if pos < 1 || pos+7 > len(bechString) || len(bechString) > 110 {
		return "", nil, 0, nil
	}

	hrp = bechString[:pos]
	data = make([]int, 0, len(bechString)-pos-1)

	for i := pos + 1; i < len(bechString); i++ {
		d := strings.IndexByte(bech32Charset, bechString[i])
		if d == -1 {
			return "", nil, 0, nil
		}
		data = append(data, d)
	}

	// Try BECH32 first, then BECH32M
	if bech32VerifyChecksum(hrp, data, bech32Encoding) {
		return hrp, data[:len(data)-6], bech32Encoding, nil
	}
	if bech32VerifyChecksum(hrp, data, bech32mEncoding) {
		return hrp, data[:len(data)-6], bech32mEncoding, nil
	}

	return "", nil, 0, nil
}

// Bech32Encode encodes data to a Bech32 string.
func Bech32Encode(hrp string, data []int, encoding uint32) string {
	combined := append(data, bech32CreateChecksum(hrp, data, encoding)...)
	ret := hrp + "1"
	for _, d := range combined {
		ret += string(bech32Charset[d])
	}
	return ret
}

// convertBits converts data between bit widths.
func convertBits(data []int, fromBits, toBits int, pad bool) []int {
	acc := 0
	bits := 0
	var ret []int
	maxv := (1 << toBits) - 1

	for _, value := range data {
		if value < 0 || (value>>fromBits) != 0 {
			return nil
		}
		acc = (acc << fromBits) | value
		bits += fromBits
		for bits >= toBits {
			bits -= toBits
			ret = append(ret, (acc>>bits)&maxv)
		}
	}

	if pad {
		if bits > 0 {
			ret = append(ret, (acc<<(toBits-bits))&maxv)
		}
	} else if bits >= fromBits || ((acc<<(toBits-bits))&maxv) != 0 {
		return nil
	}

	return ret
}

// SegwitDecode decodes a SegWit address.
func SegwitDecode(hrp, addr string) (version int, program []int, valid bool) {
	decHrp, data, encoding, _ := Bech32Decode(addr)
	if decHrp == "" || decHrp != hrp || len(data) < 1 || data[0] > 16 {
		return 0, nil, false
	}

	program = convertBits(data[1:], 5, 8, false)
	if program == nil || len(program) < 2 || len(program) > 40 {
		return 0, nil, false
	}

	// Version 0 must be 20 or 32 bytes
	if data[0] == 0 && len(program) != 20 && len(program) != 32 {
		return 0, nil, false
	}

	// Version 0 must use BECH32, other versions must use BECH32M
	if data[0] == 0 && encoding != bech32Encoding {
		return 0, nil, false
	}
	if data[0] != 0 && encoding != bech32mEncoding {
		return 0, nil, false
	}

	return data[0], program, true
}

// SegwitEncode encodes a SegWit address.
func SegwitEncode(hrp string, version int, program []int) string {
	encoding := uint32(bech32Encoding)
	if version > 0 {
		encoding = bech32mEncoding
	}

	data := convertBits(program, 8, 5, true)
	if data == nil {
		return ""
	}

	return Bech32Encode(hrp, append([]int{version}, data...), encoding)
}

// ValidateSegwitAddress validates a SegWit address for the given HRP and allowed versions.
func ValidateSegwitAddress(address string, hrps []string, allowedVersions []int) bool {
	for _, hrp := range hrps {
		version, program, valid := SegwitDecode(hrp, address)
		if valid {
			for _, allowedVersion := range allowedVersions {
				if version == allowedVersion {
					// Verify by re-encoding
					encoded := SegwitEncode(hrp, version, program)
					return encoded == strings.ToLower(address)
				}
			}
		}
	}
	return false
}
