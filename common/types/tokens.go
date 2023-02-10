package types

import (
	nodeTypes "github.com/aximchain/flash-node/common/types"
	"github.com/aximchain/flash-node/plugins/tokens/client/rest"
)

type (
	Token        = nodeTypes.Token
	MiniToken    = nodeTypes.MiniToken
	TokenBalance = rest.TokenBalance
)
