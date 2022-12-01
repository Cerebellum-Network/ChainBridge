// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/Cerebellum-Network/chainbridge-utils/msg"
	"github.com/Cerebellum-Network/go-substrate-rpc-client/v4/types"
)

func InitializeChain(client *Client, relayers []types.AccountID, chains []msg.ChainId, resources map[msg.ResourceId]Method, threshold uint32) error {
	calls := []types.Call{}

	// Create AddRelayer calls
	for _, relayer := range relayers {
		call, err := client.NewAddRelayerCall(relayer)
		if err != nil {
			return err
		}
		calls = append(calls, call)
	}
	// Create WhitelistChain calls
	for _, chain := range chains {
		call, err := client.NewWhitelistChainCall(chain)
		if err != nil {
			return err
		}
		calls = append(calls, call)
	}

	// Create SetResource calls
	for id, method := range resources {
		call, err := client.NewRegisterResourceCall(id, string(method))
		if err != nil {
			return err
		}
		calls = append(calls, call)
	}

	// Create a SetThreshold call
	call, err := client.NewSetRelayerThresholdCall(types.U32(threshold))
	if err != nil {
		return err
	}
	calls = append(calls, call)

	// Create a NewBalancesTransfer call
	amount := types.NewUCompactFromUInt(10000000000000000)
	var rId msg.ResourceId
	queryErr := QueryConst(client, "ChainBridge", "BridgeAccountId", &rId)
	if queryErr != nil {
		return queryErr
	}
	multi, err := types.NewMultiAddressFromHexAccountID(rId.Hex())
	if err != nil {
		return err
	}
	nativeTransferCall, err := client.NewBalancesTransferCall(multi, amount)
	if err != nil {
		return err
	}
	calls = append(calls, nativeTransferCall)

	return BatchSubmit(client, calls)
}
