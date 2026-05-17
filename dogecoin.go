package address

// Dogecoin uses base58check like Bitcoin with different version bytes.
// P2PKH: 0x1e (D prefix), P2SH: 0x16 (9 or A prefix)

var dogecoinMainnet = bitcoinConfig{
	addressTypes: []string{"1e", "16"},
}

// ValidateDogecoin validates a Dogecoin address.
func ValidateDogecoin(address string) bool {
	return validateBase58Address(address, dogecoinMainnet.addressTypes)
}
