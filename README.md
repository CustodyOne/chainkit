# Chainkit

A modular blockchain toolkit designed for Wallet-as-a-Service (WaaS) infrastructure, providing unified interfaces for multi-chain transaction management.

## Architecture

Chainkit follows a clean separation of concerns through three core abstractions:

- **Client** - Handles all blockchain RPC interactions including balance queries, transaction input fetching, and transaction broadcasting
- **TxBuilder** - Constructs transaction payloads for transfers and staking operations
- **Signer** - Manages cryptographic signing with support for multiple signature algorithms (K256-Keccak, K256-SHA256, Ed25519, Schnorr)

This architecture enables secure key management by isolating signing operations from network communication.

## Supported Protocols

| Protocol | Networks | Signature |
|----------|----------|-----------|
| `btc` | Bitcoin | K256-SHA256 |
| `btc-cash` | Bitcoin Cash | K256-SHA256 |
| `btc-legacy` | Dogecoin, Litecoin | K256-SHA256 |
| `evm` | Ethereum, Avalanche, Polygon, Optimism, Arbitrum, Berachain | K256-Keccak |
| `evm-legacy` | BNB Chain, Fantom, Klaytn, and more | K256-Keccak |
| `cosmos` | Cosmos Hub, Terra, Injective, Sei, Celestia | K256-Keccak |
| `solana` | Solana | Ed25519 |
| `ton` | TON | Ed25519 |
| `tron` | Tron | K256-Keccak |

## Installation

```bash
go get github.com/CustodyOne/chainkit
```

### CLI Tool

Install the `xc` command-line utility for quick blockchain interactions:

```bash
go install -v ./cmd/xc/...
```

## CLI Usage

### Available Commands

```
xc [command]

Commands:
  address     Derive address from PRIVATE_KEY environment variable
  balance     Query asset balance (returned as raw integer value)
  chains      Display supported chain information
  staking     Staking operations (stake, unstake, withdraw)
  transfer    Execute asset transfers with decimal amount input
  tx-info     Retrieve on-chain transaction details
  tx-input    Fetch transaction input parameters

Flags:
      --chain string      Target blockchain (required)
      --config string     Configuration file path
      --not-mainnet       Use testnet/devnet instead of mainnet
      --provider string   Client provider (BTC chains only)
      --rpc string        Custom RPC endpoint
  -v, --verbose           Enable verbose output
```

### Address Derivation

```bash
export PRIVATE_KEY=<your-private-key>
xc address --chain SOL
```

### Transfer Operations

Native asset transfer:

```bash
xc transfer <recipient> 0.1 --chain ETH -v
```

Token transfer (specify contract and decimals):

```bash
xc transfer <recipient> 100 --chain SOL \
  --contract EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v \
  --decimals 6
```

### Staking Operations

Delegate assets to a validator:

```bash
xc staking stake --amount 0.1 --chain SOL \
  --validator he1iusunGwqrNtafDtLdhsUQDFvo13z9sUa36PauBtk \
  --rpc https://api.mainnet-beta.solana.com
```

### Balance Queries

Native balance:

```bash
xc balance 0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5 --chain ETH
```

Token balance:

```bash
xc balance 0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5 --chain ETH \
  --contract 0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48
```

### Transaction Inspection

```bash
xc tx-info --chain BTC b5734126a7b9f5a3a94491c7297959b74099c5c88d2f5f34ea3cb432abdf9c5e
```

### Transaction Input

Fetch required parameters for constructing a new transaction:

```bash
xc tx-input 0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5 --chain ETH
```

## Library Usage

### Initialize Factory

```go
import "github.com/CustodyOne/chainkit/factory"

f := factory.NewDefaultFactory()
```

### Create Client

```go
client, err := f.NewClient(chainConfig)
if err != nil {
    return err
}

balance, err := client.FetchBalance(ctx, address)
```

### Build Transactions

```go
builder, err := f.NewTxBuilder(chainConfig)
if err != nil {
    return err
}

tx, err := builder.NewTransfer(args, txInput)
```

### Sign and Broadcast

```go
signer, err := f.NewSigner(chainConfig, privateKey)
if err != nil {
    return err
}

signature, err := signer.Sign(tx.GetDataToSign())
tx.AddSignature(signature)

err = client.BroadcastTx(ctx, tx)
```

## Staking Providers

Chainkit integrates with institutional staking providers:

- **Native** - Direct on-chain staking
- **Kiln** - Enterprise staking infrastructure
- **Figment** - Institutional staking services
- **Twinstake** - Regulated staking solutions

## Development

### Running Local Devnets

Spin up local blockchain nodes for testing:

```bash
# EVM devnet
cd blockchain/evm/node
docker build -t devnet-evm .
docker run --name devnet-evm -p 10000:10000 -p 10001:10001 devnet-evm
```

Similar configurations are available for Solana, Cosmos, and other supported chains under their respective `node/` directories.

## License

Proprietary - CustodyOne
