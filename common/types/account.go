package types

import (
	"github.com/aximchain/axc-cosmos-sdk/types"
	"github.com/aximchain/axc-cosmos-sdk/x/auth"
	nodeTypes "github.com/aximchain/flash-node/common/types"
)

type (
	AppAccount  = nodeTypes.AppAccount
	BaseAccount = auth.BaseAccount

	Account      = types.Account
	NamedAccount = nodeTypes.NamedAccount
)

// Balance Account definition
type BalanceAccount struct {
	Number    int64          `json:"account_number"`
	Address   string         `json:"address"`
	Balances  []TokenBalance `json:"balances"`
	PublicKey []uint8        `json:"public_key"`
	Sequence  int64          `json:"sequence"`
	Flags     uint64         `json:"flags"`
}
