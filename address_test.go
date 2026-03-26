package address

import "testing"

func TestValidateBitcoin(t *testing.T) {
	valid := []string{
		"12KYrjTdVGjFMtaxERSk3gphreJ5US8aUP",
		"12QeMLzSrB8XH8FvEzPMVoRxVAzTr5XM2y",
		"15uwigGExiNQxTNr1QSZYPXJMp9Px2YnVU",
		"3FyVFsEyyBPzHjD3qUEgX7Jsn4tcHNZFkn",
		"38mKdURe1zcQyrFqRLzR8PRao3iLGEPVsU",
		"BC1QW508D6QEJXTDG4Y5R3ZARVARY0C5XW7KV8F3T4",
		"1oNLrsHnBcR6dpaBpwz3LSwutbUNkNSjs",
		"1HVDCg2KrPBH1Mg5SK9fGjAR9KVqyMMdBC",
		"1SQHtwR5oJRKLfiWQ2APsAd9miUc4k2ez",
		"116CGDLddrZhMrTwhCVJXtXQpxygTT1kHd",
		"3NJZLcZEEYBpxYEUGewU4knsQRn1WM5Fkt",
		"bc1q2t63ewm3mvh0ztmnmezxm7s0tefknenxlrlwrk",
	}

	invalid := []string{
		"",
		"invalid",
	}

	for _, addr := range valid {
		if !ValidateBitcoin(addr, Mainnet) {
			t.Errorf("ValidateBitcoin(%q) = false, want true", addr)
		}
	}

	for _, addr := range invalid {
		if ValidateBitcoin(addr, Mainnet) {
			t.Errorf("ValidateBitcoin(%q) = true, want false", addr)
		}
	}
}

func TestValidateEthereum(t *testing.T) {
	valid := []string{
		"0xE37c0D48d68da5c5b14E5c1a9f1CFE802776D9FF",
		"0xa00354276d2fC74ee91e37D085d35748613f4748",
		"0xAff4d6793F584a473348EbA058deb8caad77a288",
		"0xc6d9d2cd449a754c494264e1809c50e34d64562b",
		"0x52908400098527886E0F7030069857D2E4169EE7",
		"0x8617E340B3D01FA5F11F306F4090FD50E238070D",
		"0x27b1fdb04752bbc536007a920d24acb045561c26",
		"0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed",
		"0xfB6916095ca1df60bB79Ce92cE3Ea74c37c5d359",
		"0xdbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB",
		"0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb",
		"0xde709f2102306220921060314715629080e2fb77",
	}

	invalid := []string{
		"",
		"0x",
		"0xGGGG",
		"invalid",
	}

	for _, addr := range valid {
		if !ValidateEthereum(addr) {
			t.Errorf("ValidateEthereum(%q) = false, want true", addr)
		}
	}

	for _, addr := range invalid {
		if ValidateEthereum(addr) {
			t.Errorf("ValidateEthereum(%q) = true, want false", addr)
		}
	}
}

func TestValidateSolana(t *testing.T) {
	valid := []string{
		"833XorXTTx5iya5B3Tr6iqEs9GbRuvVfwyLCP2vpdzhq",
		"6ZRCB7AAqGre6c72PRz3MHLC73VMYvJ8bi9KHf1HFpNk",
		"HgyXhqapicB8zoyyFQ23oUwwFrBACDyDc7bqUuvnEELM",
		"69UwBV4LPg7hHUS5JXiXyfgVnESmDKe8KJppsLj8pRU",
		"G4qGCGF4vWGPzYi2pxc2Djvgv3j8NiWaHQMgTVebCX6W",
	}

	invalid := []string{
		"",
		"invalid",
		"3Yu3ULPjVc4QqS349VCs22br9zH9T6MWNnSM9RBimkw",
	}

	for _, addr := range valid {
		if !ValidateSolana(addr) {
			t.Errorf("ValidateSolana(%q) = false, want true", addr)
		}
	}

	for _, addr := range invalid {
		if ValidateSolana(addr) {
			t.Errorf("ValidateSolana(%q) = true, want false", addr)
		}
	}
}

func TestValidateRipple(t *testing.T) {
	valid := []string{
		"rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
		"r3kmLJN5D28dHuH8vZNUZpMC43pEHpaocV",
		"rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
		"rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
	}

	invalid := []string{
		"",
		"invalid",
	}

	for _, addr := range valid {
		if !ValidateRipple(addr) {
			t.Errorf("ValidateRipple(%q) = false, want true", addr)
		}
	}

	for _, addr := range invalid {
		if ValidateRipple(addr) {
			t.Errorf("ValidateRipple(%q) = true, want false", addr)
		}
	}
}

func TestValidateRippleMemo(t *testing.T) {
	valid := []string{
		"",
		"4294967295",
		"50000",
	}

	invalid := []string{
		"non-numeric-memo",
		"4294967296",
		"-1",
	}

	for _, memo := range valid {
		if !ValidateRippleMemo(memo) {
			t.Errorf("ValidateRippleMemo(%q) = false, want true", memo)
		}
	}

	for _, memo := range invalid {
		if ValidateRippleMemo(memo) {
			t.Errorf("ValidateRippleMemo(%q) = true, want false", memo)
		}
	}
}

func TestValidateTron(t *testing.T) {
	valid := []string{
		"TNDzfERDpxLDS2w1q6yaFC7pzqaSQ3Bg3r",
		"TFEkshkSXo8yMe8vcA6P77XmiLfstNWHyT",
		"TUBBzKNM9gr687ucwj8fvVS2Sf2e4WseVa",
	}

	invalid := []string{
		"",
		"invalid",
		"TJCnKsPa7y5okkXvQAidZBzqx3QyQ6sxM",
	}

	for _, addr := range valid {
		if !ValidateTron(addr) {
			t.Errorf("ValidateTron(%q) = false, want true", addr)
		}
	}

	for _, addr := range invalid {
		if ValidateTron(addr) {
			t.Errorf("ValidateTron(%q) = true, want false", addr)
		}
	}
}

func TestValidateLitecoin(t *testing.T) {
	valid := []string{
		"LVg2kJoFNg45Nbpy53h7Fe1wKyeXVRhMH9",
		"LTpYZG19YmfvY2bBDYtCKpunVRw7nVgRHW",
		"Lb6wDP2kHGyWC7vrZuZAgV7V4ECyDdH7a6",
		"ltc1qg42tkwuuxefutzxezdkdel39gfstuap288mfea",
		"ltc1qhxka8d59lylj76rpsflhagp76nvckggd8x0tt7jtnphpkmfv3snsk6suq0",
		"LLUvRjeoNCa1gRuoVzFAzr3pWdu21JJESa",
	}

	invalid := []string{
		"",
		"invalid",
	}

	for _, addr := range valid {
		if !ValidateLitecoin(addr, Mainnet) {
			t.Errorf("ValidateLitecoin(%q) = false, want true", addr)
		}
	}

	for _, addr := range invalid {
		if ValidateLitecoin(addr, Mainnet) {
			t.Errorf("ValidateLitecoin(%q) = true, want false", addr)
		}
	}
}

func TestValidateBitcoinCash(t *testing.T) {
	valid := []string{
		"12KYrjTdVGjFMtaxERSk3gphreJ5US8aUP",
		"12QeMLzSrB8XH8FvEzPMVoRxVAzTr5XM2y",
		"1oNLrsHnBcR6dpaBpwz3LSwutbUNkNSjs",
		"3NJZLcZEEYBpxYEUGewU4knsQRn1WM5Fkt",
		"bitcoincash:qq4v32mtagxac29my6gwj6fd4tmqg8rysu23dax807",
	}

	invalid := []string{
		"",
		"invalid",
	}

	for _, addr := range valid {
		if !ValidateBitcoinCash(addr, Mainnet) {
			t.Errorf("ValidateBitcoinCash(%q) = false, want true", addr)
		}
	}

	for _, addr := range invalid {
		if ValidateBitcoinCash(addr, Mainnet) {
			t.Errorf("ValidateBitcoinCash(%q) = true, want false", addr)
		}
	}
}

func TestValidateMonero(t *testing.T) {
	valid := []string{
		"47zQ5LAivg6hNCgijXSEFVLX7mke1bgM6YGLFaANDoJbgXDymcAAZvvMNt2PmMpqEe5qRy2zyfMYXdwpmdyitiFh84xnPG2",
		"48bWuoDG75CXMDHbmPEvUF2hm1vLDic7ZJ7hqRkL65QR9p13AQAX4eEACXNk4YP115Q4KRVZnAvmMBHrcGfv9FvKPZnH6vH",
		"88WB4JKdQVhWfkc8cBT9EEJ6vejSAqKJHbV1dXBAXdpgQovtNDNRxfKCS7wB8rHQ5D5zH2Pd1YkyMNNQDie6ZfeZ311fPgn",
	}

	validIntegrated := []string{
		"4Gd4DLiXzBmbVX2FZZ3Cvu6fUaWACup1qDowprUCje1kSP4FmbftiJMSfV8kWZXNqmVwj4m52xqtgFNUudVmsmGkGvkLcCibWfVUfUFVB7",
		"4J5sF94AzXgFgx8LuWc9dcWkJkGkD3cL3L2AuhX6QA9jFvSxxj6QhHqHXqM2b2Go7G8RyDzEbHxYd9G26XUUbuJChipEyBz9fENMU2Ua9b",
	}

	invalid := []string{
		"",
		"invalid",
		"4AdUndXHHZ6cfufTMvppY6JwXNou",
	}

	for _, addr := range valid {
		if !ValidateMonero(addr, Mainnet) {
			t.Errorf("ValidateMonero(%q) = false, want true", addr)
		}
	}

	for _, addr := range validIntegrated {
		if !ValidateMonero(addr, Mainnet) {
			t.Errorf("ValidateMonero integrated (%q) = false, want true", addr)
		}
	}

	for _, addr := range invalid {
		if ValidateMonero(addr, Mainnet) {
			t.Errorf("ValidateMonero(%q) = true, want false", addr)
		}
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		coin    Coin
		address string
		valid   bool
	}{
		{BTC, "12KYrjTdVGjFMtaxERSk3gphreJ5US8aUP", true},
		{BTC, "bc1q2t63ewm3mvh0ztmnmezxm7s0tefknenxlrlwrk", true},
		{ETH, "0xE37c0D48d68da5c5b14E5c1a9f1CFE802776D9FF", true},
		{ETH, "0xde709f2102306220921060314715629080e2fb77", true},
		{SOL, "833XorXTTx5iya5B3Tr6iqEs9GbRuvVfwyLCP2vpdzhq", true},
		{XRP, "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn", true},
		{TRX, "TNDzfERDpxLDS2w1q6yaFC7pzqaSQ3Bg3r", true},
		{LTC, "LVg2kJoFNg45Nbpy53h7Fe1wKyeXVRhMH9", true},
		{BCH, "bitcoincash:qq4v32mtagxac29my6gwj6fd4tmqg8rysu23dax807", true},
		{XMR, "47zQ5LAivg6hNCgijXSEFVLX7mke1bgM6YGLFaANDoJbgXDymcAAZvvMNt2PmMpqEe5qRy2zyfMYXdwpmdyitiFh84xnPG2", true},
		{USDT, "0xE37c0D48d68da5c5b14E5c1a9f1CFE802776D9FF", true},
		{USDT, "TNDzfERDpxLDS2w1q6yaFC7pzqaSQ3Bg3r", true},
		{BTC, "", false},
		{ETH, "invalid", false},
	}

	for _, tt := range tests {
		result := Validate(tt.coin, tt.address)
		if result != tt.valid {
			t.Errorf("Validate(%v, %q) = %v, want %v", tt.coin, tt.address, result, tt.valid)
		}
	}
}

func TestValidateForNetwork(t *testing.T) {
	tests := []struct {
		coin    Coin
		network string
		address string
		valid   bool
	}{
		{USDT, "ERC20", "0xE37c0D48d68da5c5b14E5c1a9f1CFE802776D9FF", true},
		{USDT, "TRC20", "TNDzfERDpxLDS2w1q6yaFC7pzqaSQ3Bg3r", true},
		{USDT, "ETHEREUM", "0xE37c0D48d68da5c5b14E5c1a9f1CFE802776D9FF", true},
		{USDT, "TRON", "TNDzfERDpxLDS2w1q6yaFC7pzqaSQ3Bg3r", true},
		{USDT, "SOL", "833XorXTTx5iya5B3Tr6iqEs9GbRuvVfwyLCP2vpdzhq", true},
		{USDT, "BSC", "0xE37c0D48d68da5c5b14E5c1a9f1CFE802776D9FF", true},
		{USDT, "ERC20", "TNDzfERDpxLDS2w1q6yaFC7pzqaSQ3Bg3r", false},
		{USDT, "TRC20", "0xE37c0D48d68da5c5b14E5c1a9f1CFE802776D9FF", false},
	}

	for _, tt := range tests {
		result := ValidateForNetwork(tt.coin, tt.network, tt.address)
		if result != tt.valid {
			t.Errorf("ValidateForNetwork(%v, %q, %q) = %v, want %v", tt.coin, tt.network, tt.address, result, tt.valid)
		}
	}
}
