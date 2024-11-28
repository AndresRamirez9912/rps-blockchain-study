package rpsKeeper

import (
	"context"
	"time"

	"cosmossdk.io/collections"
	"github.com/0xlb/rps-chain/x/rps/rules"
	"github.com/0xlb/rps-chain/x/rps/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) EndBlocker(ctx context.Context) error {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)
	// Fetch active games whose esxpireaction preiod have ended
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	ranger := collections.NewPrefixUntilPairRange[uint64, uint64](uint64(sdkCtx.BlockHeight()))
	err := k.ActiveGamesQueue.Walk(ctx, ranger, func(key collections.Pair[uint64, uint64]) (stop bool, err error) {
		// Get the games
		game, err := k.Games.Get(ctx, key.K2())
		if err != nil {
			return false, err
		}

		// Check that status is InProgress or Waiting
		// then update the game status to cancelled
		if game.Status == rules.StatusWaiting || game.Status == rules.StatusInProgress {
			game.Status = rules.StatusCancelled
			err = k.Games.Set(ctx, game.GameNumber, game)
			if err != nil {
				return false, err
			}
		}

		// Remove the grame from the queue
		err = k.ActiveGamesQueue.Remove(ctx, collections.Join(game.ExpirationHeight, game.GameNumber))
		if err != nil {
			return false, err
		}
		return false, nil
	})

	if err != nil {
		return err
	}

	return nil
}
