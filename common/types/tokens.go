package types

import (
	nodeTypes "github.com/aximchain/beacon-node/common/types"
	"github.com/aximchain/beacon-node/plugins/tokens/client/rest"
)

type (
	Token        = nodeTypes.Token
	MiniToken    = nodeTypes.MiniToken
	TokenBalance = rest.TokenBalance
)
