// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

// An available method on the substrate chain
type Method string

var AddRelayerMethod Method = BridgePalletName + ".add_relayer"
var SetResourceMethod Method = BridgePalletName + ".set_resource"
var SetThresholdMethod Method = BridgePalletName + ".set_threshold"
var WhitelistChainMethod Method = BridgePalletName + ".whitelist_chain"
var Erc20TransferNativeMethod Method = "Erc20.transfer_native"
var Erc20TransferErc721Method Method = "Erc20.transfer_erc721"
var Erc20TransferHashMethod Method = "Erc20.transfer_hash"
var Erc20MintErc721Method Method = "Erc20.mint_erc721"
var Erc20TransferMethod Method = "Erc20.transfer"
var Erc20RemarkMethod Method = "Erc20.remark"
var Erc721MintMethod Method = "Erc721.mint"
var SudoMethod Method = "Sudo.sudo"
var BalancesTransferMethod Method = "Balances.transfer"
