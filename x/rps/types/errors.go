package types

import "cosmossdk.io/errors"

var (
	ErrorInvalidAddress      = errors.Register(ModuleName, 1, "Invalid address")
	ErrorRoundsOutOfBounds   = errors.Register(ModuleName, 2, "Game rounds out of bounds")
	ErrorInvalidMovesNumber  = errors.Register(ModuleName, 3, "Invalid player moves count")
	ErrorInvalidGameNumber   = errors.Register(ModuleName, 4, "Invalid game number. Should be greater than zero")
	ErrorInvalidGameStatus   = errors.Register(ModuleName, 5, "Invalid game status")
	ErrorInvalidScore        = errors.Register(ModuleName, 6, "Invalid score size")
	ErrorInvalidOponent      = errors.Register(ModuleName, 7, "Invalid oponent address")
	ErrorDuplicatedGameIndex = errors.Register(ModuleName, 8, "Invalid game, duplicated game index")
	ErrorInvalidMove         = errors.Register(ModuleName, 9, "Invalid move")
	ErrorGameEnded           = errors.Register(ModuleName, 10, "The error has ended")
	ErrorInvalidPlayer       = errors.Register(ModuleName, 11, "Player invalid")
	ErrorPlayerCantMove      = errors.Register(ModuleName, 12, "Player cant move")
)
