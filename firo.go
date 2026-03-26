package address

// FIRO (formerly Zcoin) uses Bitcoin-style addresses with different version bytes.
// Mainnet: a (82) for P2PKH, Z (7) for P2SH
// Addresses are base58check encoded.

var firoMainnet = bitcoinConfig{
	addressTypes: []string{"52", "07"}, // 82 (0x52) for P2PKH, 7 (0x07) for P2SH
	bech32Hrps:   []string{},           // FIRO doesn't use SegWit
	segwitVers:   []int{},
}

var firoTestnet = bitcoinConfig{
	addressTypes: []string{"41", "b2"}, // Testnet versions
	bech32Hrps:   []string{},
	segwitVers:   []int{},
}

// ValidateFiro validates a FIRO address.
func ValidateFiro(address string, network Network) bool {
	config := firoMainnet
	if network == Testnet {
		config = firoTestnet
	}

	return validateBase58Address(address, config.addressTypes)
}
