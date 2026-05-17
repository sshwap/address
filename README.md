# address

Cryptocurrency address validation library for Go. Supports 35 coins with format, checksum, and network-aware validation.

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

    // Validate on a specific network (for multi-chain tokens)
    valid = address.ValidateForNetwork(address.USDT, "TRC20", "T...")
    valid = address.ValidateForNetwork(address.USDC, "SOL", "6ZRCB...")
    valid = address.ValidateForNetwork(address.DOGE, "BSC", "0x...")

    // Validate memo/tag
    valid = address.ValidateMemo(address.XRP, "123456")
    valid = address.ValidateMemo(address.XLM, "memo text")
}
```

## Supported Cryptocurrencies

| Coin | Validation | Networks |
|------|-----------|----------|
| BTC | Base58 checksum + SegWit bech32 | Mainnet, Testnet (P2PKH, P2SH, SegWit v0, Taproot v1) |
| ETH | Regex + EIP-55 checksum | Mainnet |
| USDT | ETH/TRX/SOL format | ERC20, TRC20, SPL, BEP20 |
| USDC | ETH/TRX/SOL/ALGO format | ERC20, TRC20, SPL, BEP20, Algorand, Arbitrum, Optimism, Base, Avalanche, Aptos, SUI |
| XRP | Ripple base58 checksum + memo validation | Mainnet |
| BNB | Ethereum format | BSC, OPBNB |
| SOL | Base58 (32 bytes) | Mainnet |
| ADA | Bech32 (Shelley addr1) + Base58 (Byron) | Mainnet, Testnet |
| DOGE | Base58 checksum (version 0x1e/0x16) | Mainnet |
| LINK | Ethereum format | ETH, BSC, Arbitrum, Polygon |
| TON | Base64 user-friendly (EQ/UQ) + raw (0:hex) | Mainnet |
| AVAX | Ethereum format (C-Chain) | C-Chain, BSC |
| SUI | 0x + 64 hex | Mainnet |
| DOT | SS58 with blake2b-512 checksum | Mainnet |
| XLM | Base32 with CRC-16/XMODEM checksum + memo validation | Mainnet |
| HBAR | Account ID format (0.0.xxxxx) + memo validation | Mainnet |
| NEAR | Named accounts (*.near) + implicit (64 hex) | Mainnet |
| ATOM | Bech32 "cosmos" HRP (20 bytes) + memo validation | Mainnet, BSC |
| XMR | CryptoNote base58 + keccak256 checksum | Mainnet, Testnet (standard + integrated) |
| LTC | Base58 checksum + SegWit bech32 | Mainnet, Testnet |
| BCH | CashAddr polymod + Legacy base58 | Mainnet, Testnet |
| TRX | Base58 checksum (0x41 prefix) | Mainnet |
| ETC | Ethereum format | Mainnet, BSC |
| APT | 0x + 1-64 hex | Mainnet |
| KAS | Bech32-like with "kaspa:" prefix | Mainnet |
| FIL | f0 (ID) / f1-f2 (secp256k1/actor) / f3 (BLS) / f4 (delegated) | Mainnet |
| VET | Ethereum format | Mainnet, BSC |
| ARB | Ethereum format | Arbitrum, ETH |
| OP | Ethereum format | Optimism |
| ALGO | Base32 with SHA-512/256 checksum | Mainnet |
| STX | C32check format (SP/ST prefix) + memo validation | Mainnet, Testnet |
| INJ | Bech32 "inj" HRP (20 bytes) | Mainnet |
| SEI | Bech32 "sei" HRP + Ethereum format | Mainnet, SEI EVM |
| FIRO | Base58 checksum | Mainnet, Testnet |
| ZANO | CryptoNote base58 + keccak256 checksum | Mainnet |

## Memo/Tag Validation

Coins that support or require memo/destination tag fields:

| Coin | Field | Validation |
|------|-------|-----------|
| XRP | Destination tag | Numeric, 0-4294967295 |
| XLM | Memo | Text, max 28 chars |
| HBAR | Memo | Text, max 100 chars |
| ATOM | Memo | Text, max 256 chars |
| INJ | Memo | Text, max 256 chars |
| TON | Comment | Text, max 512 chars |
| STX | Memo | Text, max 34 chars |

## Multi-Network Support

`ValidateForNetwork` handles coins that exist on multiple chains. It routes to the correct address format validator based on the network parameter:

- **BSC/BEP20 wrapped tokens**: Always validated as Ethereum format
- **L2 networks** (Arbitrum, Optimism, Base, zkSync, Polygon, Avalanche C-Chain): Ethereum format
- **USDT/USDC**: Network-specific routing (ERC20->ETH, TRC20->TRX, SOL->Solana, etc.)
