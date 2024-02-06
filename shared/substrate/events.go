// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	events "github.com/Cerebellum-Network/chainbridge-substrate-events"
	"github.com/Cerebellum-Network/go-substrate-rpc-client/v7/types"
)

type EventErc721Minted struct {
	Phase   types.Phase
	Owner   types.AccountID
	TokenId types.U256
	Topics  []types.Hash
}

type EventErc721Transferred struct {
	Phase   types.Phase
	From    types.AccountID
	To      types.AccountID
	TokenId types.U256
	Topics  []types.Hash
}

type EventErc721Burned struct {
	Phase   types.Phase
	TokenId types.AccountID
	Topics  []types.Hash
}

type EventErc20Remark struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

// EventNFTDeposited is emitted when NFT is ready to be deposited to other chain.
type EventNFTDeposited struct {
	Phase  types.Phase
	Asset  types.Hash
	Topics []types.Hash
}

// EventFeeChanged is emitted when a fee for a given key is changed.
type EventFeeChanged struct {
	Phase    types.Phase
	Key      types.Hash
	NewPrice types.U128
	Topics   []types.Hash
}

// EventNewMultiAccount is emitted when a multi account has been created.
// First param is the account that created it, second is the multisig account.
type EventNewMultiAccount struct {
	Phase   types.Phase
	Who, ID types.AccountID
	Topics  []types.Hash
}

// EventMultiAccountUpdated is emitted when a multi account has been updated. First param is the multisig account.
type EventMultiAccountUpdated struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventMultiAccountRemoved is emitted when a multi account has been removed. First param is the multisig account.
type EventMultiAccountRemoved struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventNewMultisig is emitted when a new multisig operation has begun.
// First param is the account that is approving, second is the multisig account.
type EventNewMultisig struct {
	Phase   types.Phase
	Who, ID types.AccountID
	Topics  []types.Hash
}

// TimePoint contains height and index
type TimePoint struct {
	Height types.U32
	Index  types.U32
}

// EventMultisigApproval is emitted when a multisig operation has been approved by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigApproval struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Topics    []types.Hash
}

// EventMultisigExecuted is emitted when a multisig operation has been executed by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigExecuted struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Result    types.DispatchResult
	Topics    []types.Hash
}

// EventMultisigCancelled is emitted when a multisig operation has been cancelled by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigCancelled struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Topics    []types.Hash
}

type EventTreasuryMinting struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventRadClaimsClaimed is emitted when RAD Tokens have been claimed
type EventRadClaimsClaimed struct {
	Phase  types.Phase
	Who    types.AccountID
	Value  types.U128
	Topics []types.Hash
}

// EventRadClaimsRootHashStored is emitted when RootHash has been stored for the correspondent RAD Claims batch
type EventRadClaimsRootHashStored struct {
	Phase    types.Phase
	RootHash types.Hash
	Topics   []types.Hash
}

// EventNftTransferred is emitted when the ownership of the asset has been transferred to the account
type EventNftTransferred struct {
	Phase      types.Phase
	RegistryId RegistryId
	AssetId    AssetId
	Who        types.AccountID
	Topics     []types.Hash
}

// EventRegistryMint is emitted when successfully minting an NFT
type EventRegistryMint struct {
	Phase      types.Phase
	RegistryId RegistryId
	TokenId    TokenId
	Topics     []types.Hash
}

// EventRegistryRegistryCreated is emitted when successfully creating a NFT registry
type EventRegistryRegistryCreated struct {
	Phase      types.Phase
	RegistryId RegistryId
	Topics     []types.Hash
}

// EventRegistryTmp is emitted only for testing
type EventRegistryTmp struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

type EventDataStringSet struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}

type EventDataStringChanged struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}

type EventBillingReportInitialized struct {
	Phase     types.Phase
	ClusterId types.H160
	Era       types.U32
	Topics    []types.Hash
}

type EventChargingStarted struct {
	Phase     types.Phase
	ClusterId types.H160
	Era       types.U32
	Topics    []types.Hash
}

type EventCharged struct {
	Phase      types.Phase
	ClusterId  types.H160
	Era        types.U32
	BatchIndex types.U32
	CustomerId types.AccountID
	Amount     types.U128
	Topics     []types.Hash
}

type EventChargeFailed struct {
	Phase            types.Phase
	ClusterId        types.H160
	Era              types.U32
	BatchIndex       types.U32
	CustomerId       types.AccountID
	Charged          types.U128
	ExpectedToCharge types.U128
	Topics           []types.Hash
}

type EventIndebted struct {
	Phase      types.Phase
	ClusterId  types.H160
	Era        types.U32
	BatchIndex types.U32
	CustomerId types.AccountID
	Amount     types.U128
	Topics     []types.Hash
}

type EventChargingFinished struct {
	Phase     types.Phase
	ClusterId types.H160
	Era       types.U32
	Topics    []types.Hash
}

type EventTreasuryFeesCollected struct {
	Phase     types.Phase
	ClusterId types.H160
	Era       types.U32
	Amount    types.U128
	Topics    []types.Hash
}

type EventClusterReserveFeesCollected struct {
	Phase     types.Phase
	ClusterId types.H160
	Era       types.U32
	Amount    types.U128
	Topics    []types.Hash
}

type EventValidatorFeesCollected struct {
	Phase     types.Phase
	ClusterId types.H160
	Era       types.U32
	Amount    types.U128
	Topics    []types.Hash
}

type EventRewardingStarted struct {
	Phase     types.Phase
	ClusterId types.H160
	Era       types.U32
	Topics    []types.Hash
}

type EventRewarded struct {
	Phase            types.Phase
	ClusterId        types.H160
	Era              types.U32
	BatchIndex       types.U32
	NodeProviderId   types.AccountID
	Rewarded         types.U128
	ExpectedToReward types.U128
	Topics           []types.Hash
}

type EventNotDistributedReward struct {
	Phase             types.Phase
	ClusterId         types.H160
	Era               types.U32
	BatchIndex        types.U32
	NodeProviderId    types.AccountID
	ExpectedReward    types.U128
	DistributedReward types.U128
	Topics            []types.Hash
}

type EventNotDistributedOverallReward struct {
	Phase                  types.Phase
	ClusterId              types.H160
	Era                    types.U32
	ExpectedReward         types.U128
	TotalDistributedReward types.U128
	Topics                 []types.Hash
}

type EventRewardingFinished struct {
	Phase     types.Phase
	ClusterId types.H160
	Era       types.U32
	Topics    []types.Hash
}

type EventBillingReportFinalized struct {
	Phase     types.Phase
	ClusterId types.H160
	Era       types.U32
	Topics    []types.Hash
}

type EventAuthorisedCaller struct {
	Phase            types.Phase
	AuthorisedCaller types.AccountID
	Topics           []types.Hash
}

type Events struct {
	types.EventRecords
	events.Events
	Erc721_Minted                          []EventErc721Minted                //nolint:stylecheck,golint
	Erc721_Transferred                     []EventErc721Transferred           //nolint:stylecheck,golint
	Erc721_Burned                          []EventErc721Burned                //nolint:stylecheck,golint
	Erc20_Remark                           []EventErc20Remark                 //nolint:stylecheck,golint
	Nfts_DepositAsset                      []EventNFTDeposited                //nolint:stylecheck,golint
	Council_Proposed                       []types.EventCouncilProposed       //nolint:stylecheck,golint
	Council_Voted                          []types.EventCouncilVoted          //nolint:stylecheck,golint
	Council_Approved                       []types.EventCouncilApproved       //nolint:stylecheck,golint
	Council_Disapproved                    []types.EventCouncilDisapproved    //nolint:stylecheck,golint
	Council_Executed                       []types.EventCouncilExecuted       //nolint:stylecheck,golint
	Council_MemberExecuted                 []types.EventCouncilMemberExecuted //nolint:stylecheck,golint
	Council_Closed                         []types.EventCouncilClosed         //nolint:stylecheck,golint
	Fees_FeeChanged                        []EventFeeChanged                  //nolint:stylecheck,golint
	MultiAccount_NewMultiAccount           []EventNewMultiAccount             //nolint:stylecheck,golint
	MultiAccount_MultiAccountUpdated       []EventMultiAccountUpdated         //nolint:stylecheck,golint
	MultiAccount_MultiAccountRemoved       []EventMultiAccountRemoved         //nolint:stylecheck,golint
	MultiAccount_NewMultisig               []EventNewMultisig                 //nolint:stylecheck,golint
	MultiAccount_MultisigApproval          []EventMultisigApproval            //nolint:stylecheck,golint
	MultiAccount_MultisigExecuted          []EventMultisigExecuted            //nolint:stylecheck,golint
	MultiAccount_MultisigCancelled         []EventMultisigCancelled           //nolint:stylecheck,golint
	TreasuryReward_TreasuryMinting         []EventTreasuryMinting             //nolint:stylecheck,golint
	Nft_Transferred                        []EventNftTransferred              //nolint:stylecheck,golint
	RadClaims_Claimed                      []EventRadClaimsClaimed            //nolint:stylecheck,golint
	RadClaims_RootHashStored               []EventRadClaimsRootHashStored     //nolint:stylecheck,golint
	Registry_Mint                          []EventRegistryMint                //nolint:stylecheck,golint
	Registry_RegistryCreated               []EventRegistryRegistryCreated     //nolint:stylecheck,golint
	Registry_RegistryTmp                   []EventRegistryTmp                 //nolint:stylecheck,golint
	CereDDCModule_DataStringSet            []EventDataStringSet               //nolint:stylecheck,golint
	CereDDCModule_DataStringChanged        []EventDataStringChanged           //nolint:stylecheck,golint
	DdcPayouts_BillingReportInitialized    []EventBillingReportInitialized    //nolint:stylecheck,golint
	DdcPayouts_ChargingStarted             []EventChargingStarted             //nolint:stylecheck,golint
	DdcPayouts_Charged                     []EventCharged                     //nolint:stylecheck,golint
	DdcPayouts_ChargedFailed               []EventChargeFailed                //nolint:stylecheck,golint
	DdcPayouts_Indebted                    []EventIndebted                    //nolint:stylecheck,golint
	DdcPayouts_ChargingFinished            []EventChargingFinished            //nolint:stylecheck,golint
	DdcPayouts_TreasuryFeesCollected       []EventTreasuryFeesCollected       //nolint:stylecheck,golint
	DdcPayouts_ClusterReserveFeesCollected []EventClusterReserveFeesCollected //nolint:stylecheck,golint
	DdcPayouts_ValidatorFeesCollected      []EventValidatorFeesCollected      //nolint:stylecheck,golint
	DdcPayouts_RewardingStarted            []EventRewardingStarted            //nolint:stylecheck,golint
	DdcPayouts_Rewarded                    []EventRewarded                    //nolint:stylecheck,golint
	DdcPayouts_NotDistributedReward        []EventNotDistributedReward        //nolint:stylecheck,golint
	DdcPayouts_NotDistributedOverallReward []EventNotDistributedOverallReward //nolint:stylecheck,golint
	DdcPayouts_RewardingFinished           []EventRewardingFinished           //nolint:stylecheck,golint
	DdcPayouts_BillingReportFinalized      []EventBillingReportFinalized      //nolint:stylecheck,golint
	DdcPayouts_AuthorisedCaller            []EventAuthorisedCaller            //nolint:stylecheck,golint
}
