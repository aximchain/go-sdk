package types

import (
	"github.com/aximchain/axc-cosmos-sdk/types"
	paramHubTypes "github.com/aximchain/axc-cosmos-sdk/x/paramHub/types"
)

const (
	FeeForProposer = types.FeeForProposer
	FeeForAll      = types.FeeForAll
	FeeFree        = types.FeeFree
)

type (
	FeeDistributeType = types.FeeDistributeType

	FeeParam         = paramHubTypes.FeeParam
	DexFeeParam      = paramHubTypes.DexFeeParam
	DexFeeField      = paramHubTypes.DexFeeField
	FixedFeeParams   = paramHubTypes.FixedFeeParams
	TransferFeeParam = paramHubTypes.TransferFeeParam
)
