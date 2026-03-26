# address

Cryptocurrency address validation library for Go.

## Installation

```bash
go get github.com/sshwap/address
```

## Usage

```go
package main

import "github.com/sshwap/address"

func main() {
    // Validate an address
    valid := address.Validate(address.BTC, "bc1q...")

    // Validate with network option
    valid = address.Validate(address.BTC, "tb1q...", address.Options{Network: address.Testnet})

    // Validate USDT on a specific network
    valid = address.ValidateForNetwork(address.USDT, "TRC20", "T...")

    // Validate XRP memo
    valid = address.ValidateMemo(address.XRP, "123456")
}
```

## Supported Cryptocurrencies

| Coin | Networks |
|------|----------|
| BTC | Mainnet, Testnet (P2PKH, P2SH, SegWit, Taproot) |
| ETH | EIP-55 checksum |
| USDT | ERC20, TRC20, SPL, BEP20 |
| XRP | Mainnet (with memo validation) |
| SOL | Mainnet |
| BCH | CashAddr, Legacy |
| XMR | Mainnet, Testnet (standard + integrated) |
| LTC | Mainnet, Testnet (P2PKH, P2SH, SegWit) |
| TRX | Mainnet |
| FIRO | Mainnet, Testnet |
| ZANO | Mainnet |
