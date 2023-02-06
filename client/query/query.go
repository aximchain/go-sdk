package query

import (
	"github.com/aximchain/go-sdk/client/basic"
	"github.com/aximchain/go-sdk/common/types"
)

type QueryClient interface {
	GetAccount(string) (*types.BalanceAccount, error)
	GetTime() (*types.Time, error)
	GetTokens(query *types.TokensQuery) ([]types.Token, error)
	GetNodeInfo() (*types.ResultStatus, error)
	GetMiniTokens(query *types.TokensQuery) ([]types.MiniToken, error)
}

type client struct {
	baseClient basic.BasicClient
}

func NewClient(c basic.BasicClient) QueryClient {
	return &client{baseClient: c}
}
