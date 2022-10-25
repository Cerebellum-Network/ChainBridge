// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

// An available method on the substrate chain
type Method string

var AddRelayerMethod Method = BridgePalletName + ".add_relayer"
var SetResourceMethod Method = BridgePalletName + ".set_resource"
var SetThresholdMethod Method = BridgePalletName + ".set_threshold"
var WhitelistChainMethod Method = BridgePalletName + ".whitelist_chain"
var ExampleTransferNativeMethod Method = "Erc20.transfer_native"
var ExampleTransferErc721Method Method = "Erc20.transfer_erc721"
var ExampleTransferHashMethod Method = "Erc20.transfer_hash"
var ExampleMintErc721Method Method = "Erc20.mint_erc721"
var ExampleTransferMethod Method = "Erc20.transfer"
var ExampleRemarkMethod Method = "Erc20.remark"
var Erc721MintMethod Method = "Erc721.mint"
var SudoMethod Method = "Sudo.sudo"
