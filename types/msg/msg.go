package msg

import (
	cTypes "github.com/aximchain/axc-cosmos-sdk/types"
	"github.com/aximchain/axc-cosmos-sdk/x/bank"
	"github.com/aximchain/axc-cosmos-sdk/x/gov"
	oracleTypes "github.com/aximchain/axc-cosmos-sdk/x/oracle/types"
	"github.com/aximchain/axc-cosmos-sdk/x/slashing"
	stakeTypes "github.com/aximchain/axc-cosmos-sdk/x/stake/types"
	"github.com/aximchain/flash-node/plugins/account"
	bTypes "github.com/aximchain/flash-node/plugins/bridge/types"
	"github.com/aximchain/flash-node/plugins/dex/order"
	dexTypes "github.com/aximchain/flash-node/plugins/dex/types"
	"github.com/aximchain/flash-node/plugins/tokens/burn"
	"github.com/aximchain/flash-node/plugins/tokens/freeze"
	"github.com/aximchain/flash-node/plugins/tokens/issue"
	"github.com/aximchain/flash-node/plugins/tokens/ownership"
	"github.com/aximchain/flash-node/plugins/tokens/seturi"
	"github.com/aximchain/flash-node/plugins/tokens/swap"
	"github.com/aximchain/flash-node/plugins/tokens/timelock"
)

// Msg definition
type (
	SmartChainAddress = cTypes.SmartChainAddress

	// bridge module
	BindMsg        = bTypes.BindMsg
	TransferOutMsg = bTypes.TransferOutMsg
	UnbindMsg      = bTypes.UnbindMsg

	// token module
	TokenBurnMsg         = burn.BurnMsg
	DexListMsg           = dexTypes.ListMsg
	ListMiniMsg          = dexTypes.ListMiniMsg
	TokenFreezeMsg       = freeze.FreezeMsg
	TokenUnfreezeMsg     = freeze.UnfreezeMsg
	TokenIssueMsg        = issue.IssueMsg
	MiniTokenIssueMsg    = issue.IssueMiniMsg
	TinyTokenIssueMsg    = issue.IssueTinyMsg
	MintMsg              = issue.MintMsg
	SendMsg              = bank.MsgSend
	SetURIMsg            = seturi.SetURIMsg
	TimeLockMsg          = timelock.TimeLockMsg
	TimeRelockMsg        = timelock.TimeRelockMsg
	TimeUnlockMsg        = timelock.TimeUnlockMsg
	TransferOwnershipMsg = ownership.TransferOwnershipMsg

	// gov module
	SubmitProposalMsg          = gov.MsgSubmitProposal
	DepositMsg                 = gov.MsgDeposit
	VoteMsg                    = gov.MsgVote
	SideChainSubmitProposalMsg = gov.MsgSideChainSubmitProposal
	SideChainDepositMsg        = gov.MsgSideChainDeposit
	SideChainVoteMsg           = gov.MsgSideChainVote

	// atomic swap module
	HTLTMsg        = swap.HTLTMsg
	DepositHTLTMsg = swap.DepositHTLTMsg
	ClaimHTLTMsg   = swap.ClaimHTLTMsg
	RefundHTLTMsg  = swap.RefundHTLTMsg

	// oracle claim module
	Claim    = oracleTypes.Claim
	ClaimMsg = oracleTypes.ClaimMsg

	// trade module
	CreateOrderMsg = order.NewOrderMsg
	CancelOrderMsg = order.CancelOrderMsg

	// account module
	SetAccountFlagsMsg = account.SetAccountFlagsMsg

	// slash module
	MsgSideChainUnjail = slashing.MsgSideChainUnjail
	MsgUnjail          = slashing.MsgUnjail

	// stake module
	CreateSideChainValidatorMsg = stakeTypes.MsgCreateSideChainValidator
	EditSideChainValidatorMsg   = stakeTypes.MsgEditSideChainValidator
	SideChainDelegateMsg        = stakeTypes.MsgSideChainDelegate
	SideChainRedelegateMsg      = stakeTypes.MsgSideChainRedelegate
	SideChainUndelegateMsg      = stakeTypes.MsgSideChainUndelegate
	MsgCreateValidatorOpen      = stakeTypes.MsgCreateValidatorOpen
	MsgRemoveValidator          = stakeTypes.MsgRemoveValidator
	MsgEditValidator            = stakeTypes.MsgEditValidator
	MsgDelegate                 = stakeTypes.MsgDelegate
	MsgRedelegate               = stakeTypes.MsgRedelegate
	MsgUndelegate               = stakeTypes.MsgUndelegate
)

var (
	NewSmartChainAddress = cTypes.NewSmartChainAddress

	// bridge module
	NewBindMsg        = bTypes.NewBindMsg
	NewTransferOutMsg = bTypes.NewTransferOutMsg
	NewUnbindMsg      = bTypes.NewUnbindMsg

	// token module
	NewTokenBurnMsg         = burn.NewMsg
	NewDexListMsg           = dexTypes.NewListMsg
	NewListMiniMsg          = dexTypes.NewListMiniMsg
	NewFreezeMsg            = freeze.NewFreezeMsg
	NewUnfreezeMsg          = freeze.NewUnfreezeMsg
	NewTokenIssueMsg        = issue.NewIssueMsg
	NewMiniTokenIssueMsg    = issue.NewIssueMiniMsg
	NewTinyTokenIssueMsg    = issue.NewIssueTinyMsg
	NewMintMsg              = issue.NewMintMsg
	NewMsgSend              = bank.NewMsgSend
	NewSetUriMsg            = seturi.NewSetUriMsg
	NewTimeLockMsg          = timelock.NewTimeLockMsg
	NewTimeRelockMsg        = timelock.NewTimeRelockMsg
	NewTimeUnlockMsg        = timelock.NewTimeUnlockMsg
	NewTransferOwnershipMsg = ownership.NewTransferOwnershipMsg

	// gov module
	NewDepositMsg                 = gov.NewMsgDeposit
	NewMsgVote                    = gov.NewMsgVote
	NewMsgSubmitProposal          = gov.NewMsgSubmitProposal
	NewSideChainSubmitProposalMsg = gov.NewMsgSideChainSubmitProposal
	NewSideChainDepositMsg        = gov.NewMsgSideChainDeposit
	NewSideChainVoteMsg           = gov.NewMsgSideChainVote

	// atomic swap module
	NewHTLTMsg        = swap.NewHTLTMsg
	NewDepositHTLTMsg = swap.NewDepositHTLTMsg
	NewClaimHTLTMsg   = swap.NewClaimHTLTMsg
	NewRefundHTLTMsg  = swap.NewRefundHTLTMsg

	// oracle claim module
	NewClaim    = oracleTypes.NewClaim
	NewClaimMsg = oracleTypes.NewClaimMsg

	// trade module
	NewCreateOrderMsg = order.NewNewOrderMsg
	NewCancelOrderMsg = order.NewCancelOrderMsg

	// account module
	NewSetAccountFlagsMsg = account.NewSetAccountFlagsMsg

	// slash module
	NewMsgSideChainUnjail = slashing.NewMsgSideChainUnjail
	NewMsgUnjail          = slashing.NewMsgUnjail

	// stake module
	NewCreateSideChainValidatorMsg           = stakeTypes.NewMsgCreateSideChainValidator
	NewMsgCreateSideChainValidatorOnBehalfOf = stakeTypes.NewMsgCreateSideChainValidatorOnBehalfOf
	NewEditSideChainValidatorMsg             = stakeTypes.NewMsgEditSideChainValidator
	NewSideChainDelegateMsg                  = stakeTypes.NewMsgSideChainDelegate
	NewSideChainRedelegateMsg                = stakeTypes.NewMsgSideChainRedelegate
	NewSideChainUndelegateMsg                = stakeTypes.NewMsgSideChainUndelegate
	NewMsgCreateValidatorOpen                = stakeTypes.NewMsgRemoveValidator
	NewMsgRemoveValidator                    = stakeTypes.NewMsgRemoveValidator
	NewMsgEditValidator                      = stakeTypes.NewMsgEditValidator
	NewMsgDelegate                           = stakeTypes.NewMsgDelegate
	NewMsgRedelegate                         = stakeTypes.NewMsgRedelegate
	NewMsgUndelegate                         = stakeTypes.NewMsgUndelegate
)
