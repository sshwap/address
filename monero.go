package address

import (
	"encoding/hex"
	"math/big"
	"regexp"
)

// Monero address patterns
var moneroStandardRegex = regexp.MustCompile(`^[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{95}$`)
var moneroIntegratedRegex = regexp.MustCompile(`^[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{106}$`)

// Monero network address types
var moneroAddressTypes = map[Network][]int{
	Mainnet: {18, 42}, // 18 = standard, 42 = subaddress
	Testnet: {53, 63},
}

var moneroIntegratedTypes = map[Network][]int{
	Mainnet: {19},
	Testnet: {54},
}

// CryptoNote base58 constants
const cnBase58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

var cnBase58AlphabetMap = make(map[byte]int)
var cnEncodedBlockSizes = []int{0, 2, 3, 5, 6, 7, 9, 10, 11}

func init() {
	for i := 0; i < len(cnBase58Alphabet); i++ {
		cnBase58AlphabetMap[cnBase58Alphabet[i]] = i
	}
}

// ValidateMonero validates a Monero address.
func ValidateMonero(address string, network Network) bool {
	isStandard := moneroStandardRegex.MatchString(address)
	isIntegrated := moneroIntegratedRegex.MatchString(address)

	if !isStandard && !isIntegrated {
		return false
	}

	// Decode using CryptoNote base58
	decoded := cnBase58Decode(address)
	if decoded == "" {
		return false
	}

	// Validate network type
	addrType, err := hex.DecodeString(decoded[:2])
	if err != nil || len(addrType) == 0 {
		return false
	}

	var validTypes []int
	if isIntegrated {
		validTypes = moneroIntegratedTypes[network]
	} else {
		validTypes = moneroAddressTypes[network]
	}

	typeValid := false
	for _, t := range validTypes {
		if int(addrType[0]) == t {
			typeValid = true
			break
		}
	}
	if !typeValid {
		return false
	}

	// Verify checksum (last 8 hex chars = 4 bytes)
	if len(decoded) < 8 {
		return false
	}

	addrChecksum := decoded[len(decoded)-8:]
	payload := decoded[:len(decoded)-8]

	// Convert payload hex to bytes
	payloadBytes, err := hex.DecodeString(payload)
	if err != nil {
		return false
	}

	// Compute Keccak-256 checksum
	hashChecksum := keccak256Checksum(payloadBytes)

	return addrChecksum == hashChecksum
}

// cnBase58Decode decodes a CryptoNote base58 encoded string to hex.
func cnBase58Decode(encoded string) string {
	if len(encoded) == 0 {
		return ""
	}

	fullBlockCount := len(encoded) / 11
	lastBlockSize := len(encoded) % 11
	lastBlockDecodedSize := -1

	for i, size := range cnEncodedBlockSizes {
		if size == lastBlockSize {
			lastBlockDecodedSize = i
			break
		}
	}

	if lastBlockDecodedSize < 0 {
		return ""
	}

	dataSize := fullBlockCount*8 + lastBlockDecodedSize
	data := make([]byte, dataSize)

	for i := 0; i < fullBlockCount; i++ {
		block := encoded[i*11 : i*11+11]
		decoded, ok := cnDecodeBlock(block)
		if !ok {
			return ""
		}
		copy(data[i*8:], decoded)
	}

	if lastBlockSize > 0 {
		block := encoded[fullBlockCount*11:]
		decoded, ok := cnDecodeBlock(block)
		if !ok {
			return ""
		}
		copy(data[fullBlockCount*8:], decoded[:lastBlockDecodedSize])
	}

	return hex.EncodeToString(data)
}

// cnDecodeBlock decodes a single CryptoNote base58 block.
func cnDecodeBlock(block string) ([]byte, bool) {
	if len(block) < 1 || len(block) > 11 {
		return nil, false
	}

	resSize := -1
	for i, size := range cnEncodedBlockSizes {
		if size == len(block) {
			resSize = i
			break
		}
	}
	if resSize <= 0 {
		return nil, false
	}

	// Convert to big integer
	num := big.NewInt(0)
	base := big.NewInt(58)

	for i := 0; i < len(block); i++ {
		digit, ok := cnBase58AlphabetMap[block[i]]
		if !ok {
			return nil, false
		}
		num.Mul(num, base)
		num.Add(num, big.NewInt(int64(digit)))
	}

	// Check for overflow
	maxVal := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(8*resSize)), nil)
	if num.Cmp(maxVal) >= 0 {
		return nil, false
	}

	// Convert to bytes (big-endian)
	result := make([]byte, resSize)
	numBytes := num.Bytes()

	// Pad with zeros if necessary
	offset := resSize - len(numBytes)
	copy(result[offset:], numBytes)

	return result, true
}
