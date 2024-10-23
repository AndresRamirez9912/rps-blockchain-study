package rpsKeeper

import (
	"context"

	"github.com/0xlb/rps-chain/x/rps/types"
)

// InitGenesis initializes the module state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *types.GenesisState) error {
	// Set the genesis params into the state
	if err := k.Params.Set(ctx, data.Params); err != nil {
		return err
	}

	// Set the genesis games into the state
	for _, game := range data.Games {
		err := k.Games.Set(ctx, game.GameNumber, game)
		if err != nil {
			return nil
		}
	}

	return nil
}

// ExportGenesis exports the module state to a genesis state.
func (k *Keeper) ExportGenesis(ctx context.Context) (*types.GenesisState, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	var games []types.Game
	// The fuction walk just iterate the Games map
	err = k.Games.Walk(ctx, nil, func(gameNumber uint64, game types.Game) (stop bool, err error) {
		games = append(games, game)
		return false, nil // False because I don't want to stop the iteration
	})

	if err != nil {
		return nil, err
	}

	return &types.GenesisState{
		Params: params,
		Games:  games,
	}, nil
}
