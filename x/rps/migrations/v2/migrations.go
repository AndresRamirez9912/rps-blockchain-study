package v2

import (
	"github.com/0xlb/rps-chain/x/rps/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MigrateGame(ctx sdk.Context, game types.Game) types.Game {
	curentHeight := ctx.BlockHeight()
	game.ExpirationHeight = uint64(curentHeight) + types.DefaultParams().Ttl
	return game
}
