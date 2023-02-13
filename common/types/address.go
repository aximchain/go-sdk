package types

import (
	"github.com/aximchain/axc-cosmos-sdk/types"
)

type AccAddress = types.AccAddress

type ChainNetwork uint8

const (
	TestNetwork ChainNetwork = iota
	ProdNetwork
	TmpTestNetwork
	GangesNetwork
)

const (
	AddrLen = types.AddrLen
)

var Network = ProdNetwork

func SetNetwork(network ChainNetwork) {
	Network = network
	if network != ProdNetwork {
		sdkConfig := types.GetConfig()
		sdkConfig.SetBech32PrefixForAccount("taxc", "axcp")
	}
}

func (this ChainNetwork) Bech32Prefixes() string {
	switch this {
	case TestNetwork:
		return "taxc"
	case TmpTestNetwork:
		return "taxc"
	case GangesNetwork:
		return "taxc"
	case ProdNetwork:
		return "axc"
	default:
		panic("Unknown network type")
	}
}

func init() {
	sdkConfig := types.GetConfig()
	sdkConfig.SetBech32PrefixForAccount("axc", "axcp")
	sdkConfig.SetBech32PrefixForValidator("ava", "avap")
	sdkConfig.SetBech32PrefixForConsensusNode("aca", "acap")
}

var (
	AccAddressFromHex    = types.AccAddressFromHex
	AccAddressFromBech32 = types.AccAddressFromBech32
	GetFromBech32        = types.GetFromBech32
	MustBech32ifyConsPub = types.MustBech32ifyConsPub
	Bech32ifyConsPub     = types.Bech32ifyConsPub
	GetConsPubKeyBech32  = types.GetConsPubKeyBech32
)
