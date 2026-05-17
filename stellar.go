package address

import (
	"encoding/base32"
	"encoding/binary"
)

// crc16XmodemTable for Stellar's CRC-16/XMODEM checksum.
var crc16XmodemTable [256]uint16

func init() {
	for i := 0; i < 256; i++ {
		crc := uint16(i << 8)
		for j := 0; j < 8; j++ {
			if crc&0x8000 != 0 {
				crc = (crc << 1) ^ 0x1021
			} else {
				crc <<= 1
			}
		}
		crc16XmodemTable[i] = crc
	}
}

func crc16Xmodem(data []byte) uint16 {
	crc := uint16(0)
	for _, b := range data {
		crc = (crc << 8) ^ crc16XmodemTable[byte(crc>>8)^b]
	}
	return crc
}

// ValidateStellar validates a Stellar (XLM) address.
// Public keys: G + 55 base32 chars = 56 total, decoding to 35 bytes.
// Version byte for public key: 6 << 3 = 0x30.
func ValidateStellar(address string) bool {
	if address == "" || len(address) != 56 || address[0] != 'G' {
		return false
	}

	decoded, err := base32.StdEncoding.DecodeString(address)
	if err != nil || len(decoded) != 35 {
		return false
	}

	if decoded[0] != 6<<3 {
		return false
	}

	// CRC-16 XMODEM of first 33 bytes, stored little-endian in last 2 bytes
	expected := crc16Xmodem(decoded[:33])
	actual := binary.LittleEndian.Uint16(decoded[33:35])
	return expected == actual
}

// ValidateStellarMemo validates a Stellar memo.
func ValidateStellarMemo(memo string) bool {
	if memo == "" {
		return true
	}
	return len(memo) <= 28
}
