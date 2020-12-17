package genutil

import (
	abci "github.com/libonomy/aphelion-staking/abci/types"

	"github.com/libonomy/cusp-sdk/codec"
	sdk "github.com/libonomy/cusp-sdk/types"
	"github.com/libonomy/cusp-sdk/x/genutil/types"
)

// InitGenesis - initialize accounts and deliver genesis transactions
func InitGenesis(
	ctx sdk.Context, cdc *codec.Codec, stakingKeeper types.StakingKeeper,
	deliverTx deliverTxfn, genesisState GenesisState,
) []abci.ValidatorUpdate {

	var validators []abci.ValidatorUpdate
	if len(genesisState.GenTxs) > 0 {
		validators = DeliverGenTxs(ctx, cdc, genesisState.GenTxs, stakingKeeper, deliverTx)
	}

	return validators
}
