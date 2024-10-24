package rpsKeeper

import (
	"context"
	"fmt"
	"strings"

	"github.com/0xlb/rps-chain/x/rps/rules"
	"github.com/0xlb/rps-chain/x/rps/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	k Keeper
}

var _ types.MsgServer = msgServer{}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{
		k: keeper,
	}
}

// Implement MsgServer interface
func (ms msgServer) CreateGame(ctx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	// Get params
	params, err := ms.k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	expirationHeight := uint64(sdkCtx.BlockHeight()) + params.Ttl

	newGame := types.Game{
		GameNumber:       ms.k.NextGameNumber(ctx),
		PlayerA:          msg.Creator,
		PlayerB:          msg.Oponent,
		Rounds:           msg.Rounds,
		Status:           rules.StatusWaiting, // Default status for creating a game
		Score:            []uint64{0, 0},
		ExpirationHeight: expirationHeight,
	}
	// Validate the game
	err = newGame.Validate()
	if err != nil {
		return nil, err
	}

	// Save newGame in the state
	err = ms.k.Games.Set(ctx, newGame.GameNumber, newGame)
	if err != nil {
		return nil, err
	}

	// *************** EMIT EVENT ************
	sdkCtx.EventManager().EmitTypedEvent(&types.EventCreateGame{
		GameNumber: newGame.GameNumber,
		PlayerA:    newGame.PlayerA,
		PlayerB:    newGame.PlayerB,
	})
	// ***************************************

	return &types.MsgCreateGameResponse{}, nil // Return the response as the interface needs (and proto file specify)
}

func (ms msgServer) MakeMove(ctx context.Context, msg *types.MsgMakeMove) (*types.MsgMakeMoveResponse, error) {
	// *************** VALIDATION ************
	// Validate whether game exists
	game, err := ms.k.Games.Get(ctx, msg.GameIndex)
	if err != nil {
		return nil, err
	}
	// ***************************************

	// Validate game status is InProgress or Waiting
	if game.Ended() {
		return nil, types.ErrorGameEnded
	}

	// Add move to the player array
	err = game.AddPlayerMove(msg.Player, msg.Move)
	if err != nil {
		return nil, err
	}

	// *************** GAME IN PROGRESS ******
	game.Status = rules.StatusInProgress
	err = game.Validate()
	if err != nil {
		return nil, err
	}

	err = ms.k.Games.Set(ctx, game.GameNumber, game)
	if err != nil {
		return nil, err
	}
	// ***************************************

	// *************** EMIT EVENT ************
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitTypedEvent(&types.EventMakeMove{
		GameNumber: game.GameNumber,
		Player:     msg.Player,
		Move:       msg.Move,
	})
	// ***************************************

	return &types.MsgMakeMoveResponse{}, nil
}

func (ms msgServer) ReviewMove(ctx context.Context, msg *types.MsgReviewMove) (*types.MsgReviewMoveResponse, error) {
	// *************** VALIDATION ************
	// Game exists
	game, err := ms.k.Games.Get(ctx, msg.GameIndex) // Get game by id
	if err != nil {
		return nil, err
	}

	// Game status is Inprogress or waiting
	if game.Ended() {
		return nil, types.ErrorGameEnded
	}

	err = game.RevealPlayerMove(msg.Player, msg.RevealedMove, msg.Salt)
	if err != nil {
		return nil, err
	}

	// Get the new game's status
	// if playerAmoveCount == playerBMovesCount => the round is completed
	if game.IsRoundRevealed() {
		playerAResult := rules.DetermineRoundWinner(
			rules.Choice(game.GetPlayerALastMove()),
			rules.Choice(game.GetPlayerBLastMove()),
		)

		if playerAResult == rules.Win {
			game.AddWinToPlayerA()
		}

		if playerAResult == rules.Loss {
			game.AddWinToPlayerA()
		}
	}

	game.Status = rules.GameResultByMajority(game.GetPlayerAScore(), game.GetPlayerBScore(), game.Rounds)

	// *************** GAME ENDED ************
	if game.Ended() {

	}
	// ***************************************

	// *************** EMIT EVENT ************
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitTypedEvent(&types.EventEndGame{
		GameNumber: game.GameNumber,
		Status:     game.Status,
	})
	// ***************************************
	return &types.MsgReviewMoveResponse{}, nil
}

func (ms msgServer) UpdateParams(ctx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	// *************** VALIDATION ************
	// Validate exists Authority
	_, err := ms.k.addressCodec.StringToBytes(msg.Authority)
	if err != nil {
		return nil, err
	}

	// Validate the correct Authority
	authority := ms.k.GetAuthority()
	if !strings.EqualFold(msg.Authority, authority) {
		return nil, fmt.Errorf("unauthorized, got: %s, want %s", msg.Authority, ms.k.GetAuthority())
	}

	// Validate mesage's params
	err = msg.Params.Validate()
	if err != nil {
		return nil, err
	}
	// ***************************************

	// ************ UPDATE PARAMS ************
	err = ms.k.Params.Set(ctx, msg.Params)
	if err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
	// ***************************************
}
