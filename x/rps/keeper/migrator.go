package rpsKeeper

import (
	"cosmossdk.io/collections"
	v2 "github.com/0xlb/rps-chain/x/rps/migrations/v2"
	"github.com/0xlb/rps-chain/x/rps/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Migrator struct {
	keeper Keeper
}

func NewMigrator(k Keeper) Migrator {
	return Migrator{
		keeper: k,
	}
}

func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	// migrate games - add the expiration height for existing games
	err := m.keeper.Games.Walk(ctx, nil, func(key uint64, game types.Game) (stop bool, err error) {
		mg := v2.MigrateGame(ctx, game)
		err = m.keeper.Games.Set(ctx, key, mg)
		if err != nil {
			return true, err
		}
		// Add these to the active game list
		err = m.keeper.ActiveGamesQueue.Set(ctx, collections.Join(mg.ExpirationHeight, key))
		if err != nil {
			return true, err
		}
		return false, nil
	})
	if err != nil {
		return err
	}
	return m.keeper.Params.Set(ctx, types.DefaultParams())
}
