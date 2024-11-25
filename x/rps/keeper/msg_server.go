package rpsKeeper

import (
	"context"

	"challenge/x/rps/rules"
	"challenge/x/rps/types"
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
	newGame := types.Game{
		GameNumber: ms.k.NextGameNumber(ctx),
		PlayerA:    msg.Creator,
		PlayerB:    msg.Oponent,
		Rounds:     msg.Rounds,
		Status:     rules.StatusWaiting, // Default status for creating a game
		Score:      []uint64{0, 0},
	}
	// Validate the game
	err := newGame.Validate()
	if err != nil {
		return nil, err
	}

	// Save newGame in the state
	err = ms.k.Games.Set(ctx, newGame.GameNumber, newGame)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateGameResponse{}, nil // Return the response as the interface needs (and proto file specify)
}

func (ms msgServer) MakeMove(ctx context.Context, msg *types.MsgMakeMove) (*types.MsgMakeMoveResponse, error) {
	// *************** VALIDATION ************
	// Validate whether is a valid move
	ok := rules.IsValidMove(msg.Move)
	if !ok {
		return nil, types.ErrorInvalidMove
	}

	// Validate whether game exists
	game, err := ms.k.Games.Get(ctx, msg.GameIndex)
	if err != nil {
		return nil, err
	}

	// Validate game status is InProgress or Waiting
	if game.Status != rules.StatusInProgress && game.Status != rules.StatusWaiting {
		return nil, types.ErrorGameEnded
	}
	// ***************************************

	// *************** GAME IN PROGRESS ******
	var player rules.Player
	// Add move to the player
	switch msg.Player {
	case game.PlayerA:
		player = rules.PlayerA
		game.PlayerAMoves = append(game.PlayerAMoves, msg.Move)
	case game.PlayerB:
		player = rules.PlayerB
		game.PlayerBMoves = append(game.PlayerBMoves, msg.Move)
	}

	if player == rules.InvalidPlayer {
		return nil, types.ErrorInvalidPlayer
	}

	// Determine whether the player can move
	playerAmovesCount := len(game.PlayerAMoves)
	playerBmovesCount := len(game.PlayerBMoves)
	ok = rules.CanMakeMove(player, playerAmovesCount, playerBmovesCount)
	if !ok {
		return nil, types.ErrorPlayerCantMove
	}

	// Update status of the game
	// if playerA moves count == playerB moves count => then a round is completed
	if playerAmovesCount == playerBmovesCount {
		playerAResult := rules.DetermineRoundWinner(rules.Choice(game.PlayerAMoves[playerAmovesCount-1]), rules.Choice(game.PlayerBMoves[playerBmovesCount-1]))
		if playerAResult == rules.Win {
			game.AddWinToPlayerA()
		}

		if playerAResult == rules.Loss {
			game.AddWinToPlayerB()
		}
	}

	// if playerA moves count != playerB moves count => then game continue
	game.Status = rules.GameResultByMajority(game.GetPlayerAScore(), game.GetPlayerBScore(), game.Rounds)

	// Validate current state of the game
	err = game.Validate()
	if err != nil {
		return nil, err
	}

	// Update the status of the game by gameNumber
	err = ms.k.Games.Set(ctx, game.GameNumber, game)
	if err != nil {
		return nil, err
	}
	// ***************************************

	return &types.MsgMakeMoveResponse{}, nil
}
