// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"math/big"

	"github.com/Cerebellum-Network/chainbridge-utils/msg"
	"github.com/Cerebellum-Network/go-substrate-rpc-client/v9/scale"
	"github.com/Cerebellum-Network/go-substrate-rpc-client/v9/types"
	"github.com/Cerebellum-Network/go-substrate-rpc-client/v9/types/codec"
)

type voteState struct {
	VotesFor     []types.AccountID
	VotesAgainst []types.AccountID
	Status       voteStatus
}

type voteStatus struct {
	IsActive   bool
	IsApproved bool
	IsRejected bool
}

func (m *voteStatus) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	if b == 0 {
		m.IsActive = true
	} else if b == 1 {
		m.IsApproved = true
	} else if b == 2 {
		m.IsRejected = true
	}

	return nil
}

// proposal represents an on-chain proposal
type proposal struct {
	depositNonce types.U64
	call         types.Call
	sourceId     types.U8
	resourceId   types.Bytes32
	method       string
}

// encode takes only nonce and call and encodes them for storage queries
func (p *proposal) encode() ([]byte, error) {
	return codec.Encode(struct {
		types.U64
		types.Call
	}{p.depositNonce, p.call})
}

func (w *writer) createFungibleProposal(m msg.Message) (*proposal, error) {
	bigAmt := big.NewInt(0).SetBytes(m.Payload[0].([]byte))
	amount := types.NewU128(*bigAmt)
	recipient, err := types.NewAccountID(m.Payload[1].([]byte))
	if err != nil {
		return nil, err
	}
	depositNonce := types.U64(m.DepositNonce)

	meta := w.conn.getMetadata()
	method, err := w.resolveResourceId(m.ResourceId)
	if err != nil {
		return nil, err
	}
	call, err := types.NewCall(
		&meta,
		method,
		recipient,
		amount,
	)
	if err != nil {
		return nil, err
	}
	if w.extendCall {
		eRID, err := codec.Encode(m.ResourceId)
		if err != nil {
			return nil, err
		}
		call.Args = append(call.Args, eRID...)
	}

	return &proposal{
		depositNonce: depositNonce,
		call:         call,
		sourceId:     types.U8(m.Source),
		resourceId:   types.NewBytes32(m.ResourceId),
		method:       method,
	}, nil
}

func (w *writer) createNonFungibleProposal(m msg.Message) (*proposal, error) {
	tokenId := types.NewU256(*big.NewInt(0).SetBytes(m.Payload[0].([]byte)))
	recipient, err := types.NewAccountID(m.Payload[1].([]byte))
	if err != nil {
		return nil, err
	}
	metadata := types.Bytes(m.Payload[2].([]byte))
	depositNonce := types.U64(m.DepositNonce)

	meta := w.conn.getMetadata()
	method, err := w.resolveResourceId(m.ResourceId)
	if err != nil {
		return nil, err
	}

	call, err := types.NewCall(
		&meta,
		method,
		recipient,
		tokenId,
		metadata,
	)
	if err != nil {
		return nil, err
	}
	if w.extendCall {
		eRID, err := codec.Encode(m.ResourceId)
		if err != nil {
			return nil, err
		}
		call.Args = append(call.Args, eRID...)
	}

	return &proposal{
		depositNonce: depositNonce,
		call:         call,
		sourceId:     types.U8(m.Source),
		resourceId:   types.NewBytes32(m.ResourceId),
		method:       method,
	}, nil
}

func (w *writer) createGenericProposal(m msg.Message) (*proposal, error) {
	meta := w.conn.getMetadata()
	method, err := w.resolveResourceId(m.ResourceId)
	if err != nil {
		return nil, err
	}

	call, err := types.NewCall(
		&meta,
		method,
		types.NewHash(m.Payload[0].([]byte)),
	)
	if err != nil {
		return nil, err
	}
	if w.extendCall {
		eRID, err := codec.Encode(m.ResourceId)
		if err != nil {
			return nil, err
		}

		call.Args = append(call.Args, eRID...)
	}
	return &proposal{
		depositNonce: types.U64(m.DepositNonce),
		call:         call,
		sourceId:     types.U8(m.Source),
		resourceId:   types.NewBytes32(m.ResourceId),
		method:       method,
	}, nil
}
