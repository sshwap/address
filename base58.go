package address

import (
	"errors"
	"math/big"
)

// Standard Bitcoin base58 alphabet
const base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

// Ripple uses a different base58 alphabet
const rippleAlphabet = "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz"

var (
	base58AlphabetMap = make(map[byte]int)
	rippleAlphabetMap = make(map[byte]int)
	bigRadix          = big.NewInt(58)
	bigZero           = big.NewInt(0)
)

func init() {
	for i := 0; i < len(base58Alphabet); i++ {
		base58AlphabetMap[base58Alphabet[i]] = i
	}
	for i := 0; i < len(rippleAlphabet); i++ {
		rippleAlphabetMap[rippleAlphabet[i]] = i
	}
}

// Base58Decode decodes a base58 encoded string using the standard Bitcoin alphabet.
func Base58Decode(input string) ([]byte, error) {
	return base58DecodeWithAlphabet(input, base58AlphabetMap)
}

// Base58DecodeRipple decodes a base58 encoded string using the Ripple alphabet.
func Base58DecodeRipple(input string) ([]byte, error) {
	return base58DecodeWithAlphabet(input, rippleAlphabetMap)
}

func base58DecodeWithAlphabet(input string, alphabetMap map[byte]int) ([]byte, error) {
	if len(input) == 0 {
		return []byte{}, nil
	}

	result := big.NewInt(0)
	for i := 0; i < len(input); i++ {
		val, ok := alphabetMap[input[i]]
		if !ok {
			return nil, errors.New("invalid base58 character")
		}
		result.Mul(result, bigRadix)
		result.Add(result, big.NewInt(int64(val)))
	}

	// Convert to bytes
	decoded := result.Bytes()

	// Count leading zeros (represented as '1' in standard base58 or 'r' in Ripple)
	leadingZeros := 0
	leadingChar := byte('1')
	if _, ok := alphabetMap['r']; ok && alphabetMap['r'] == 0 {
		leadingChar = 'r'
	}
	for i := 0; i < len(input) && input[i] == leadingChar; i++ {
		leadingZeros++
	}

	// Prepend leading zeros
	if leadingZeros > 0 {
		zeros := make([]byte, leadingZeros)
		decoded = append(zeros, decoded...)
	}

	return decoded, nil
}

// Base58Encode encodes bytes to a base58 string using the standard Bitcoin alphabet.
func Base58Encode(input []byte) string {
	return base58EncodeWithAlphabet(input, base58Alphabet)
}

func base58EncodeWithAlphabet(input []byte, alphabet string) string {
	if len(input) == 0 {
		return ""
	}

	// Count leading zeros
	leadingZeros := 0
	for i := 0; i < len(input) && input[i] == 0; i++ {
		leadingZeros++
	}

	// Convert to big integer
	num := new(big.Int).SetBytes(input)

	// Encode
	var encoded []byte
	mod := new(big.Int)
	for num.Cmp(bigZero) > 0 {
		num.DivMod(num, bigRadix, mod)
		encoded = append([]byte{alphabet[mod.Int64()]}, encoded...)
	}

	// Add leading zeros
	for i := 0; i < leadingZeros; i++ {
		encoded = append([]byte{alphabet[0]}, encoded...)
	}

	return string(encoded)
}
