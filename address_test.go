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

func TestValidateDogecoin(t *testing.T) {
	valid := []string{
		"DRSqEwcnJX3GZWH9Twtwk8D5ewqdJzi13k",
		"DBXu2kgc3xtvCUWFcxFE3r9hEYgmuaaCyD",
	}

	invalid := []string{
		"",
		"invalid",
		"0xE37c0D48d68da5c5b14E5c1a9f1CFE802776D9FF",
		"12KYrjTdVGjFMtaxERSk3gphreJ5US8aUP",
	}

	for _, addr := range valid {
		if !ValidateDogecoin(addr) {
			t.Errorf("ValidateDogecoin(%q) = false, want true", addr)
		}
	}

	for _, addr := range invalid {
		if ValidateDogecoin(addr) {
			t.Errorf("ValidateDogecoin(%q) = true, want false", addr)
		}
	}
}

func TestValidateCardano(t *testing.T) {
	valid := []string{
		"addr1qx2fxv2umyhttkxyxp8x0dlpdt3k6cwng5pxj3jhsydzer3jcu5d8ps7zex2k2xt3uqxgjqnnj83ws8lhrn648jjxtwq2ytjqp",
		"addr1q9u5sjt4qmjvlkry3qnlfjx746skkf2w7g2c4nnvpv68kfm4gp2haq5fkcq0fhxg0hpcj9r76h4l0fp7gw4tz6ct7m7qsnfjcf",
	}

	invalid := []string{
		"",
		"invalid",
		"addr1notavalidaddress",
	}

	for _, addr := range valid {
		if !ValidateCardano(addr, Mainnet) {
			t.Errorf("ValidateCardano(%q) = false, want true", addr)
		}
	}

	for _, addr := range invalid {
		if ValidateCardano(addr, Mainnet) {
			t.Errorf("ValidateCardano(%q) = true, want false", addr)
		}
	}
}

func TestValidateTON(t *testing.T) {
	valid := []string{
		"EQDtFpEwcFAEcRe5mLVh2N6C0x-_hJEM7W61_JLnSF74p4q2",
		"UQBvW8Z5huBkMJYdnfAEM5JqTNkuWX3diqYENkWsIL0XggGG",
		"0:83dfd552e63729b472fcbcc8c45ebcc6691702558b68ec7527e1ba403a0f31a8",
	}

	invalid := []string{
		"",
		"invalid",
		"EQshort",
		"0:notahexhash",
	}

	for _, addr := range valid {
		if !ValidateTON(addr) {
			t.Errorf("ValidateTON(%q) = false, want true", addr)
		}
	}

	for _, addr := range invalid {
		if ValidateTON(addr) {
			t.Errorf("ValidateTON(%q) = true, want false", addr)
		}
	}
}

func TestValidatePolkadot(t *testing.T) {
	valid := []string{
		"1FRMM8PEiWXYax7rpS6X4XZX1aAAxSWx1CrKTyrVYhV24fg",
	}

	invalid := []string{
		"",
		"invalid",
		"5short",
	}

	for _, addr := range valid {
		if !ValidatePolkadot(addr) {
			t.Errorf("ValidatePolkadot(%q) = false, want true", addr)
		}
	}

	for _, addr := range invalid {
		if ValidatePolkadot(addr) {
			t.Errorf("ValidatePolkadot(%q) = true, want false", addr)
		}
	}
}

func TestValidateStellar(t *testing.T) {
	valid := []string{
		"GAAZI4TCR3TY5OJHCTJC2A4QSY6CJWJH5IAJTGKIN2ER7LBNVKOCCWN7",
	}

	invalid := []string{
		"",
		"invalid",
		"GABCDEFG",
		"GAAZI4TCR3TY5OJHCTJC2A4QSY6CJWJH5IAJTGKIN2ER7LBNVKOCCXXX",
	}

	for _, addr := range valid {
		if !ValidateStellar(addr) {
			t.Errorf("ValidateStellar(%q) = false, want true", addr)
		}
	}

	for _, addr := range invalid {
		if ValidateStellar(addr) {
			t.Errorf("ValidateStellar(%q) = true, want false", addr)
		}
	}
}

func TestValidateHedera(t *testing.T) {
	valid := []string{
		"0.0.12345",
		"0.0.1",
		"0.0.98",
		"0.0.1234567890",
	}

	invalid := []string{
		"",
		"invalid",
		"0.0.0",
		"1.2",
		"0.0",
		"abc.def.ghi",
	}

	for _, addr := range valid {
		if !ValidateHedera(addr) {
			t.Errorf("ValidateHedera(%q) = false, want true", addr)
		}
	}

	for _, addr := range invalid {
		if ValidateHedera(addr) {
			t.Errorf("ValidateHedera(%q) = true, want false", addr)
		}
	}
}

func TestValidateNEAR(t *testing.T) {
	valid := []string{
		"alice.near",
		"bob.testnet",
		"example-account.near",
		"a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2",
		"system",
	}

	invalid := []string{
		"",
		"A",
		"UPPERCASE.near",
		"a",
		"has spaces.near",
		"double..dots.near",
	}

	for _, addr := range valid {
		if !ValidateNEAR(addr) {
			t.Errorf("ValidateNEAR(%q) = false, want true", addr)
		}
	}

	for _, addr := range invalid {
		if ValidateNEAR(addr) {
			t.Errorf("ValidateNEAR(%q) = true, want false", addr)
		}
	}
}

func TestValidateCosmos(t *testing.T) {
	valid := []string{
		"cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02",
	}

	invalid := []string{
		"",
		"invalid",
		"cosmos1invalid",
	}

	for _, addr := range valid {
		if !ValidateCosmos(addr) {
			t.Errorf("ValidateCosmos(%q) = false, want true", addr)
		}
	}

	for _, addr := range invalid {
		if ValidateCosmos(addr) {
			t.Errorf("ValidateCosmos(%q) = true, want false", addr)
		}
	}
}

func TestValidateAptos(t *testing.T) {
	valid := []string{
		"0x1",
		"0xf22bede237a07e121b56d91a491eb7bcdfd1f5907926a9e58338f964a01b17fa",
		"0x0000000000000000000000000000000000000000000000000000000000000001",
	}

	invalid := []string{
		"",
		"invalid",
		"1234",
		"0xGGGG",
	}

	for _, addr := range valid {
		if !ValidateAptos(addr) {
			t.Errorf("ValidateAptos(%q) = false, want true", addr)
		}
	}

	for _, addr := range invalid {
		if ValidateAptos(addr) {
			t.Errorf("ValidateAptos(%q) = true, want false", addr)
		}
	}
}

func TestValidateSUI(t *testing.T) {
	valid := []string{
		"0xfa7ac3951fdca92c9200f1e12e8c0a87c8e40101a2ea448f887a33ff0127d56b",
		"0x0000000000000000000000000000000000000000000000000000000000000002",
	}

	invalid := []string{
		"",
		"invalid",
		"0x1",
		"0xshort",
	}

	for _, addr := range valid {
		if !ValidateSUI(addr) {
			t.Errorf("ValidateSUI(%q) = false, want true", addr)
		}
	}

	for _, addr := range invalid {
		if ValidateSUI(addr) {
			t.Errorf("ValidateSUI(%q) = true, want false", addr)
		}
	}
}

func TestValidateKaspa(t *testing.T) {
	valid := []string{
		"kaspa:qpamkvhgh0kzx50gwvvp5xs8fakrqf6asvnlrd80esxqq4uvkwfk2ez2w6pf8",
		"kaspa:qr0jhre5a26c37fwew0mzs6yjav54r54j2axyz7xx3r04lfuwqzqcw0lsdvve",
	}

	invalid := []string{
		"",
		"invalid",
		"kaspa:short",
		"bitcoin:qpamkvhgh0kzx50gwvvp5xs8fakrqf6asvnlrd80esxqq4uvkwfk2ez2w6pf8",
	}

	for _, addr := range valid {
		if !ValidateKaspa(addr) {
			t.Errorf("ValidateKaspa(%q) = false, want true", addr)
		}
	}

	for _, addr := range invalid {
		if ValidateKaspa(addr) {
			t.Errorf("ValidateKaspa(%q) = true, want false", addr)
		}
	}
}

func TestValidateFilecoin(t *testing.T) {
	valid := []string{
		"f01234",
		"f17uoq6tp427uzv7fztkbsnn64iwotfrristwpryy",
		"f3vvmn62lofvhjd2ugzca6sof2j2ubwok6cj4xxbfzz4yuj7f2xfmzcd25bxa6ex7unvwddcq7sucifdcwmma",
	}

	invalid := []string{
		"",
		"invalid",
		"f0",
		"f00",
		"g17uoq6tp427uzv7fztkbsnn64iwotfrristwpryy",
	}

	for _, addr := range valid {
		if !ValidateFilecoin(addr) {
			t.Errorf("ValidateFilecoin(%q) = false, want true", addr)
		}
	}

	for _, addr := range invalid {
		if ValidateFilecoin(addr) {
			t.Errorf("ValidateFilecoin(%q) = true, want false", addr)
		}
	}
}

func TestValidateAlgorand(t *testing.T) {
	valid := []string{
		"737777777777777777777777777777777777777777777777777UFEJ2CI",
	}

	invalid := []string{
		"",
		"invalid",
		"short",
		"737777777777777777777777777777777777777777777777777XXXXXX",
	}

	for _, addr := range valid {
		if !ValidateAlgorand(addr) {
			t.Errorf("ValidateAlgorand(%q) = false, want true", addr)
		}
	}

	for _, addr := range invalid {
		if ValidateAlgorand(addr) {
			t.Errorf("ValidateAlgorand(%q) = true, want false", addr)
		}
	}
}

func TestValidateStacks(t *testing.T) {
	valid := []string{
		"SP2J6ZY48GV1EZ5V2V5RB9MP66SW86PYKKNRV9EJ7",
		"SP000000000000000000002Q6VF78",
	}

	invalid := []string{
		"",
		"invalid",
		"12KYrjTdVGjFMtaxERSk3gphreJ5US8aUP",
	}

	for _, addr := range valid {
		if !ValidateStacks(addr, Mainnet) {
			t.Errorf("ValidateStacks(%q) = false, want true", addr)
		}
	}

	for _, addr := range invalid {
		if ValidateStacks(addr, Mainnet) {
			t.Errorf("ValidateStacks(%q) = true, want false", addr)
		}
	}
}

func TestValidateInjective(t *testing.T) {
	valid := []string{
		"inj1ady3s7whq30l4fx8sj3x6muv5mx4dfdlcpv8n7",
	}

	invalid := []string{
		"",
		"invalid",
		"cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02",
	}

	for _, addr := range valid {
		if !ValidateInjective(addr) {
			t.Errorf("ValidateInjective(%q) = false, want true", addr)
		}
	}

	for _, addr := range invalid {
		if ValidateInjective(addr) {
			t.Errorf("ValidateInjective(%q) = true, want false", addr)
		}
	}
}

func TestValidateSei(t *testing.T) {
	valid := []string{
		// Sei EVM addresses
		"0xE37c0D48d68da5c5b14E5c1a9f1CFE802776D9FF",
		"0xde709f2102306220921060314715629080e2fb77",
	}

	invalid := []string{
		"",
		"invalid",
		"cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02",
	}

	for _, addr := range valid {
		if !ValidateSei(addr) {
			t.Errorf("ValidateSei(%q) = false, want true", addr)
		}
	}

	for _, addr := range invalid {
		if ValidateSei(addr) {
			t.Errorf("ValidateSei(%q) = true, want false", addr)
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
		// New coins
		{BNB, "0xde709f2102306220921060314715629080e2fb77", true},
		{DOGE, "DRSqEwcnJX3GZWH9Twtwk8D5ewqdJzi13k", true},
		{LINK, "0xde709f2102306220921060314715629080e2fb77", true},
		{AVAX, "0xde709f2102306220921060314715629080e2fb77", true},
		{ETC, "0xde709f2102306220921060314715629080e2fb77", true},
		{VET, "0xde709f2102306220921060314715629080e2fb77", true},
		{ARB, "0xde709f2102306220921060314715629080e2fb77", true},
		{OP, "0xde709f2102306220921060314715629080e2fb77", true},
		{HBAR, "0.0.12345", true},
		{APT, "0x1", true},
		// Invalid
		{BTC, "", false},
		{ETH, "invalid", false},
		{DOGE, "invalid", false},
		{HBAR, "invalid", false},
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
		// USDC multi-network
		{USDC, "ETH", "0xE37c0D48d68da5c5b14E5c1a9f1CFE802776D9FF", true},
		{USDC, "SOL", "833XorXTTx5iya5B3Tr6iqEs9GbRuvVfwyLCP2vpdzhq", true},
		{USDC, "TRX", "TNDzfERDpxLDS2w1q6yaFC7pzqaSQ3Bg3r", true},
		// BSC wrapped tokens
		{DOGE, "BSC", "0xE37c0D48d68da5c5b14E5c1a9f1CFE802776D9FF", true},
		{ATOM, "BSC", "0xE37c0D48d68da5c5b14E5c1a9f1CFE802776D9FF", true},
		// L2 networks
		{LINK, "ARBITRUM", "0xde709f2102306220921060314715629080e2fb77", true},
	}

	for _, tt := range tests {
		result := ValidateForNetwork(tt.coin, tt.network, tt.address)
		if result != tt.valid {
			t.Errorf("ValidateForNetwork(%v, %q, %q) = %v, want %v", tt.coin, tt.network, tt.address, result, tt.valid)
		}
	}
}

func TestValidateMemo(t *testing.T) {
	tests := []struct {
		coin  Coin
		memo  string
		valid bool
	}{
		{XRP, "", true},
		{XRP, "12345", true},
		{XRP, "4294967296", false},
		{XLM, "", true},
		{XLM, "hello", true},
		{XLM, "this memo is way too long for stellar to accept at all", false},
		{HBAR, "", true},
		{HBAR, "test memo", true},
		{ATOM, "", true},
		{ATOM, "cosmos memo", true},
		{TON, "", true},
		{TON, "ton memo", true},
		{BTC, "anything", true}, // BTC doesn't use memos
	}

	for _, tt := range tests {
		result := ValidateMemo(tt.coin, tt.memo)
		if result != tt.valid {
			t.Errorf("ValidateMemo(%v, %q) = %v, want %v", tt.coin, tt.memo, result, tt.valid)
		}
	}
}
