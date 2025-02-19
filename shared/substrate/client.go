// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"fmt"
	"math/big"

	"github.com/Cerebellum-Network/chainbridge-utils/msg"
	gsrpc "github.com/Cerebellum-Network/go-substrate-rpc-client/v9"
	"github.com/Cerebellum-Network/go-substrate-rpc-client/v9/signature"
	"github.com/Cerebellum-Network/go-substrate-rpc-client/v9/types"
	"github.com/Cerebellum-Network/go-substrate-rpc-client/v9/types/codec"
	"github.com/ChainSafe/log15"
)

// Client is a container for all the components required to submit extrinsics
// TODO: Perhaps this would benefit an interface so we can interchange Connection and a client like this
type Client struct {
	Api     *gsrpc.SubstrateAPI
	Meta    *types.Metadata
	Genesis types.Hash
	Key     *signature.KeyringPair
}

func CreateClient(key *signature.KeyringPair, endpoint string) (*Client, error) {
	c := &Client{Key: key}
	api, err := gsrpc.NewSubstrateAPI(endpoint)
	if err != nil {
		return nil, err
	}
	c.Api = api

	// Fetch metadata
	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		return nil, err
	}
	c.Meta = meta

	// Fetch genesis hash
	genesisHash, err := c.Api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return nil, err
	}
	c.Genesis = genesisHash

	return c, nil
}

// Admin calls

func (c *Client) SetRelayerThreshold(threshold types.U32) error {
	log15.Info("Setting threshold", "threshold", threshold)
	return SubmitSudoTx(c, SetThresholdMethod, threshold)
}

func (c *Client) AddRelayer(relayer types.AccountID) error {
	log15.Info("Adding relayer", "accountId", relayer)
	return SubmitSudoTx(c, AddRelayerMethod, relayer)
}

func (c *Client) WhitelistChain(id msg.ChainId) error {
	log15.Info("Whitelisting chain", "chainId", id)
	return SubmitSudoTx(c, WhitelistChainMethod, types.U8(id))
}

func (c *Client) RegisterResource(id msg.ResourceId, method string) error {
	log15.Info("Registering resource", "rId", id, "method", []byte(method))
	return SubmitSudoTx(c, SetResourceMethod, types.NewBytes32(id), []byte(method))
}

// Standard transfer calls

func (c *Client) InitiateNativeTransfer(amount types.U128, recipient []byte, destId msg.ChainId) error {
	log15.Info("Initiating Substrate native transfer", "amount", amount, "recipient", fmt.Sprintf("%x", recipient), "destId", destId)
	return SubmitTx(c, Erc20TransferNativeMethod, amount, recipient, types.U8(destId))
}

func (c *Client) InitiateNonFungibleTransfer(tokenId types.U256, recipient []byte, destId msg.ChainId) error {
	log15.Info("Initiating Substrate nft transfer", "tokenId", tokenId, "recipient", recipient, "destId", destId)
	return SubmitTx(c, Erc20TransferErc721Method, recipient, tokenId, types.U8(destId))
}

func (c *Client) InitiateHashTransfer(hash types.Hash, destId msg.ChainId) error {
	log15.Info("Initiating hash transfer", "hash", hash.Hex())
	return SubmitTx(c, Erc20TransferHashMethod, hash, types.U8(destId))
}

// Call creation methods for batching

func (c *Client) NewSudoCall(call types.Call) (types.Call, error) {
	return types.NewCall(c.Meta, string(SudoMethod), call)
}

func (c *Client) NewSetRelayerThresholdCall(threshold types.U32) (types.Call, error) {
	call, err := types.NewCall(c.Meta, string(SetThresholdMethod), threshold)
	if err != nil {
		return types.Call{}, err
	}
	return c.NewSudoCall(call)
}

func (c *Client) NewAddRelayerCall(relayer types.AccountID) (types.Call, error) {
	call, err := types.NewCall(c.Meta, string(AddRelayerMethod), relayer)
	if err != nil {
		return types.Call{}, err
	}
	return c.NewSudoCall(call)
}

func (c *Client) NewWhitelistChainCall(id msg.ChainId) (types.Call, error) {
	call, err := types.NewCall(c.Meta, string(WhitelistChainMethod), id)
	if err != nil {
		return types.Call{}, err
	}
	return c.NewSudoCall(call)
}

func (c *Client) NewRegisterResourceCall(id msg.ResourceId, method string) (types.Call, error) {
	call, err := types.NewCall(c.Meta, string(SetResourceMethod), id, method)
	if err != nil {
		return types.Call{}, err
	}
	return c.NewSudoCall(call)
}

func (c *Client) NewNativeTransferCall(amount types.U128, recipient []byte, destId msg.ChainId) (types.Call, error) {
	return types.NewCall(c.Meta, string(Erc20TransferNativeMethod), amount, recipient, types.U8(destId))
}

func (c *Client) NewBalancesTransferCall(recipient types.MultiAddress, amount types.UCompact) (types.Call, error) {
	return types.NewCall(c.Meta, string(BalancesTransferMethod), recipient, amount)
}

// Utility methods

func (c *Client) LatestBlock() (uint64, error) {
	head, err := c.Api.RPC.Chain.GetHeaderLatest()
	if err != nil {
		return 0, err
	}
	return uint64(head.Number), nil
}

func (c *Client) MintErc721(tokenId *big.Int, metadata []byte, recipient *signature.KeyringPair) error {
	fmt.Printf("Mint info: account %x amount: %x meta: %x\n", recipient.PublicKey, types.NewU256(*tokenId), types.Bytes(metadata))
	accountId, err := types.NewAccountID(recipient.PublicKey)
	if err != nil {
		return err
	}
	return SubmitSudoTx(c, Erc721MintMethod, accountId, types.NewU256(*tokenId), types.Bytes(metadata))
}

func (c *Client) OwnerOf(tokenId *big.Int) (types.AccountID, error) {
	var owner types.AccountID
	tokenIdBz, err := codec.Encode(types.NewU256(*tokenId))
	if err != nil {
		return types.AccountID{}, err
	}

	exists, err := QueryStorage(c, "TokenStorage", "TokenOwner", tokenIdBz, nil, &owner)
	if err != nil {
		return types.AccountID{}, err
	}
	if !exists {
		return types.AccountID{}, fmt.Errorf("token %s doesn't have an owner", tokenId.String())
	}
	return owner, nil
}

func (c *Client) GetDepositNonce(chain msg.ChainId) (uint64, error) {
	var count types.U64
	chainId, err := codec.Encode(types.U8(chain))
	if err != nil {
		return 0, err
	}
	exists, err := QueryStorage(c, BridgeStoragePrefix, "ChainNonces", chainId, nil, &count)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, nil
	}
	return uint64(count), nil
}
