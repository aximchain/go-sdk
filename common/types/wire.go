package types

import "github.com/tendermint/go-amino"

func RegisterWire(cdc *amino.Codec) {

	cdc.RegisterConcrete(Token{}, "axcchain/Token", nil)
	cdc.RegisterInterface((*Account)(nil), nil)
	cdc.RegisterInterface((*NamedAccount)(nil), nil)
	cdc.RegisterConcrete(&AppAccount{}, "axcchain/Account", nil)

	cdc.RegisterInterface((*FeeParam)(nil), nil)
	cdc.RegisterConcrete(&FixedFeeParams{}, "params/FixedFeeParams", nil)
	cdc.RegisterConcrete(&TransferFeeParam{}, "params/TransferFeeParams", nil)
	cdc.RegisterConcrete(&DexFeeParam{}, "params/DexFeeParam", nil)

	cdc.RegisterInterface((*Proposal)(nil), nil)
	cdc.RegisterConcrete(&TextProposal{}, "gov/TextProposal", nil)

	cdc.RegisterConcrete(MiniToken{}, "axcchain/MiniToken", nil)
}
