package rpsKeeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/core/store"
	"github.com/0xlb/rps-chain/x/rps/types"
	"github.com/cosmos/cosmos-sdk/codec"
)

type Keeper struct {
	cdc          codec.BinaryCodec // Serialize and deserialize
	addressCodec address.Codec

	// authority is the address capable of executing a MsgUpdateParams and other authority-gated message.
	// typically, this should be the x/gov module account.
	authority string

	// state management
	Schema           collections.Schema
	Params           collections.Item[types.Params]
	GameNumber       collections.Sequence                                 // Counter the games
	Games            collections.Map[uint64, types.Game]                  // Map [GameId] => Game
	ActiveGamesQueue collections.KeySet[collections.Pair[uint64, uint64]] //ActiveGamesQueue
}

// NewKeeper creates a new Keeper instance
func NewKeeper(cdc codec.BinaryCodec, addressCodec address.Codec, storeService storetypes.KVStoreService, authority string) Keeper {
	// Decode authority and check
	_, err := addressCodec.StringToBytes(authority)
	if err != nil {
		panic(fmt.Errorf("invalid authority address: %w", err))
	}

	sb := collections.NewSchemaBuilder(storeService) // instance used to define the keeper
	k := Keeper{
		cdc:              cdc,
		addressCodec:     addressCodec,
		authority:        authority,
		Params:           collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		GameNumber:       collections.NewSequence(sb, types.GameNumberKey, "game_number"),
		Games:            collections.NewMap(sb, types.GamesKey, "games", collections.Uint64Key, codec.CollValue[types.Game](cdc)),
		ActiveGamesQueue: collections.NewKeySet(sb, types.ActiveGamesQueueKey, "active_games", collections.PairKeyCodec(collections.Uint64Key, collections.Uint64Key)),
	}

	schema, err := sb.Build() // Build the whole squema
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// NextGameNumber returns and increments the global game number counter
func (k Keeper) NextGameNumber(ctx context.Context) uint64 {
	gameNumber, err := k.GameNumber.Next(ctx)
	if err != nil {
		panic(err)
	}
	return gameNumber + 1 // Sequence starts from zero
}
