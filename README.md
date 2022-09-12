# ChainBridge

<a href="https://discord.gg/ykXsJKfhgq">
  <img alt="discord" src="https://img.shields.io/discord/593655374469660673?label=Discord&logo=discord&style=flat" />
</a>
<a href="https://github.com/ChainSafe/ChainBridge/actions">
  <img alt="build status" src="https://github.com/ChainSafe/ChainBridge/workflows/Tests/badge.svg?branch=master" />
</a>

# Release Notes

## vNext
- 

## v2.6.0
- Fixed unit tests and linter
- Updated `go-substrate-rpc-client` lib to v2.0.2-cere
- Added startBlock logs
- Implemented liveness/readiness probe

## v2.5.2
- Increased ExecuteBlockWatchLimit to 500 for Ethereum/Polygon
- Added logs for gasPrice, supportedGasPrice

## v2.5.1
- Increased BlockRetryLimit to 10 for Ethereum/Polygon

## v2.5.0
- Migrated to GitHub actions

## v2.4.0
- Updated chainbridge-utils to v1.1.0

## v2.3.2
- Added min gas price check
- Updated DefaultGasPrice
- All undecoded blocks will be skipped

## v2.3.1
- Block with `System_CodeUpdated` event will be skipped 

## v2.3.0
- Added condition to prevent excess execution of executeProposal method

## v2.2.0
- Added ability to disable network scanning

## 2.1.2
- Fixed health-check endpoint initialization for networks

## 2.1.1
- Updated `chainbridge-utils`
- Added DDC custom events

## 2.1.0
- Removed bytecode check for `genericHandler`

## 2.0.0
- Changed Lib name to Cerebellum-Network
- Integrated with chainbridge-substrate-events v2
- Ingegrated with chainbridge-utils v2

# Contents

- [Getting Started](#getting-started)
- [Installation](#installation)
- [Configuration](#configuration)
- [Chain Implementations](#chain-implementations)
- [Testing](#testing)
- [Simulations](#simulations)

# Getting Started
- Check out our [documentation](https://chainbridge.chainsafe.io).
- Try [running ChainBridge locally](https://chainbridge.chainsafe.io/local/).
- Chat with us on [discord](https://discord.gg/ykXsJKfhgq).

# Installation

## Dependencies

- [Subkey](https://substrate.dev/docs/en/knowledgebase/integrate/subkey): 
Used for substrate key management. Only required if connecting to a substrate chain.


## Building

`make build`: Builds `chainbridge` in `./build`.

**or**

`make install`: Uses `go install` to add `chainbridge` to your GOBIN.

# Configuration

> Note: TOML configs have been deprecated in favour of JSON

A chain configurations take this form:

```
{
    "name": "eth",                      // Human-readable name
    "type": "ethereum",                 // Chain type (eg. "ethereum" or "substrate")
    "id": "0",                          // Chain ID
    "endpoint": "ws://<host>:<port>",   // Node endpoint
    "from": "0xff93...",                // On-chain address of relayer
    "opts": {},                         // Chain-specific configuration options (see below)
}
```

See `config.json.example` for an example configuration. 

### Ethereum Options

Ethereum chains support the following additional options:

```
{
    "bridge": "0x12345...",          // Address of the bridge contract (required)
    "erc20Handler": "0x1234...",     // Address of erc20 handler (required)
    "erc721Handler": "0x1234...",    // Address of erc721 handler (required)
    "genericHandler": "0x1234...",   // Address of generic handler (required)
    "maxGasPrice": "0x1234",         // Gas price for transactions (default: 20000000000)
    "gasLimit": "0x1234",            // Gas limit for transactions (default: 6721975)
    "gasMultiplier": "1.25",         // Multiplies the gas price by the supplied value (default: 1)
    "http": "true",                  // Whether the chain connection is ws or http (default: false)
    "startBlock": "1234",            // The block to start processing events from (default: 0)
    "blockConfirmations": "10"       // Number of blocks to wait before processing a block
    "useExtendedCall": "true"        // Extend extrinsic calls to substrate with ResourceID. Used for backward compatibility with example pallet. *Default: false*
}
```

### Substrate Options

Substrate supports the following additonal options:

```
{
    "startBlock": "1234" // The block to start processing events from (default: 0)
}
```

## Blockstore

The blockstore is used to record the last block the relayer processed, so it can pick up where it left off. 

If a `startBlock` option is provided (see [Configuration](#configuration)), then the greater of `startBlock` and the latest block in the blockstore is used at startup.

To disable loading from the blockstore specify the `--fresh` flag. A custom path for the blockstore can be provided with `--blockstore <path>`. For development, the `--latest` flag can be used to start from the current block and override any other configuration.

## Keystore

ChainBridge requires keys to sign and submit transactions, and to identify each bridge node on chain.

To use secure keys, see `chainbridge accounts --help`. The keystore password can be supplied with the `KEYSTORE_PASSWORD` environment variable.

To import external ethereum keys, such as those generated with geth, use `chainbridge accounts import --ethereum /path/to/key`.

To import private keys as keystores, use `chainbridge account import --privateKey key`.

For testing purposes, chainbridge provides 5 test keys. The can be used with `--testkey <name>`, where `name` is one of `Alice`, `Bob`, `Charlie`, `Dave`, or `Eve`. 

## Metrics

See [metrics.md](/docs/metrics.md).

# Chain Implementations

- Ethereum (Solidity): [chainbridge-solidity](https://github.com/ChainSafe/chainbridge-solidity) 

    The Solidity contracts required for chainbridge. Includes deployment and interaction CLI.
    
    The bindings for the contracts live in `bindings/`. To update the bindings modify `scripts/setup-contracts.sh` and then run `make clean && make setup-contracts`

- Substrate: [chainbridge-substrate](https://github.com/ChainSafe/chainbridge-substrate)

    A substrate pallet that can be integrated into a chain, as well as an example pallet to demonstrate chain integration.

# Docs

MKdocs will generate static HTML files for Chainsafe markdown files located in `Chainbridge/docs/`

`make install-mkdocs`: Pull the docker image MkDocs

`make mkdocs`: Run MkDoc's docker image, building and hosting the html files on `localhost:8000`  

# Local run
## Prepare

1. Build executable:
    ```
    make build
    ```
2. Import keys from Devnet:
    ```
    build/chainbridge accounts import --privateKey "improve nominee response kangaroo keen gain antenna pepper spike credit pony parrot" --sr25519
    ```
    ```
    build/chainbridge accounts import --privateKey 5502ab637600e552b45d0608f8f2888a5feafa2a5df1bc67eb9ecc69a9c6b990
    ```
    ```
    build/chainbridge accounts import --privateKey a5485b158e045febc939d9cedc79910f26efe48c31f65643fba3e6dd0f734282
    ```
    ```
    Enter a password after a prompt appears: <any_password>
    ```
3. Copy `config/config.json.example` to `config/config.json` folder:
    ```
    cp config/config.json.example config/config.json
    ```
    
## Run app in docker

```
docker-compose up
```

## Run app from source code

1. Export password:
    ```
    export KEYSTORE_PASSWORD=<any_password>
    ```
2. Run executable:
    ```
    build/chainbridge --config config/config.json --keystore ./keys --verbosity trce  --metrics --blockstore ./blockstore
    ```

# Testing

Unit tests require an ethereum node running on `localhost:8545` and a substrate node running on `localhost:9944`. E2E tests require an additional ethereum node on `localhost:8546`. 

A docker-compose file is provided to run two Geth nodes and a chainbridge-substrate-chain node in isolated environments:
```
$ docker-compose -f ./docker-compose-e2e.yml up
```

See [chainbridge-solidity](https://github.com/chainsafe/chainbridge-solidity) and [chainbridge-substrate-chain](https://github.com/ChainSafe/chainbridge-substrate-chain) for more information on testing facilities.

All Go tests can be run with:
```
$ make test
```
Go tests specifically for ethereum, substrate and E2E can be run with
```
$ make test-eth
$ make test-sub
$ make test-e2e
```

# ChainSafe Security Policy

## Reporting a Security Bug

We take all security issues seriously, if you believe you have found a security issue within a ChainSafe
project please notify us immediately. If an issue is confirmed, we will take all necessary precautions 
to ensure a statement and patch release is made in a timely manner.

Please email us a description of the flaw and any related information (e.g. reproduction steps, version) to
[security at chainsafe dot io](mailto:security@chainsafe.io).
