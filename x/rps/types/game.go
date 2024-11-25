package types

import (
	"bytes"

	"challenge/x/rps/rules"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const MaxRound = 3

func DefaultGames() (games []Game) {
	return
}

func (g Game) GetPlayerAAddress() (sdk.AccAddress, error) {
	return getPlayerAddress(g.PlayerA)
}

func (g Game) GetPlayerBAddress() (sdk.AccAddress, error) {
	return getPlayerAddress(g.PlayerB)
}

func (g Game) GetPlayerAScore() uint64 {
	return g.Score[0]
}

func (g Game) GetPlayerBScore() uint64 {
	return g.Score[1]
}

func (g *Game) AddWinToPlayerA() {
	g.Score[0]++
}

func (g *Game) AddWinToPlayerB() {
	g.Score[1]++
}

func (g Game) ValidateRounds() error {
	if g.Rounds <= MaxRound && g.Rounds > 0 {
		return nil
	}

	return ErrorRoundsOutOfBounds
}

func (g Game) ValidateMovesCount() error {
	if len(g.PlayerAMoves) <= int(g.Rounds) && len(g.PlayerBMoves) <= int(g.Rounds) {
		return nil
	}
	return ErrorInvalidMovesNumber
}

func (g Game) ValidateGameNumber() error {
	if g.GameNumber > 0 {
		return nil
	}
	return ErrorInvalidGameNumber
}

func (g Game) ValidateGameStatus() error {
	if rules.IsValidStatus(g.Status) {
		return nil
	}
	return ErrorInvalidGameStatus
}

func (g Game) ValidateScore() error {
	if len(g.Score) != 2 {
		return ErrorInvalidScore
	}

	if g.Score[0]+g.Score[1] > g.Rounds {
		return ErrorInvalidScore
	}
	return nil
}

func (g Game) Validate() error {
	// Get player A's wallet address
	accA, err := g.GetPlayerAAddress()
	if err != nil {
		return err
	}

	// Get player B's wallet address
	accB, err := g.GetPlayerBAddress()
	if err != nil {
		return err
	}

	if bytes.Equal(accA, accB) {
		return ErrorInvalidOponent
	}

	if err = g.ValidateRounds(); err != nil {
		return err
	}

	if err = g.ValidateMovesCount(); err != nil {
		return err
	}

	if err = g.ValidateGameNumber(); err != nil {
		return err
	}

	if err = g.ValidateGameStatus(); err != nil {
		return err
	}

	if err = g.ValidateScore(); err != nil {
		return err
	}

	return nil
}

func (g Game) Ended() bool {
	return g.Status == rules.StatusDraw ||
		g.Status == rules.StatusPlayerAWins ||
		g.Status == rules.StatusPlayerBWins ||
		g.Status == rules.StatusCancelled
}

func getPlayerAddress(address string) (sdk.AccAddress, error) {
	// Validate the address has our prefix (it means the wallet is from out blockchain)
	addr, err := sdk.AccAddressFromBech32(address)
	return addr, errors.Wrap(err, ErrorInvalidAddress.Error())
}
