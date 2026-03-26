package address

// Litecoin address configuration
var litecoinMainnet = bitcoinConfig{
	addressTypes: []string{"30", "32", "05"}, // L/M prefix (P2PKH), 3 prefix (P2SH legacy)
	bech32Hrps:   []string{"ltc"},
	segwitVers:   []int{0, 1},
}

var litecoinTestnet = bitcoinConfig{
	addressTypes: []string{"6f", "c4", "3a"},
	bech32Hrps:   []string{"tltc"},
	segwitVers:   []int{0, 1},
}

// ValidateLitecoin validates a Litecoin address.
func ValidateLitecoin(address string, network Network) bool {
	config := litecoinMainnet
	if network == Testnet {
		config = litecoinTestnet
	}

	// Try base58 first, then SegWit
	return validateBase58Address(address, config.addressTypes) ||
		ValidateSegwitAddress(address, config.bech32Hrps, config.segwitVers)
}
