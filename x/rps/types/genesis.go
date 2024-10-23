package types

// NewGenesisState creates a new genesis state with default values.
func NewGenesisState() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
		Games:  DefaultGames(), // Return the genesis value for a game
	}
}

// Validate performs basic genesis state validation returning an error upon any
func (gs *GenesisState) Validate() error {
	// Validate params
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	// Validate the genesis games
	unique := make(map[uint64]bool)
	for _, game := range gs.Games {
		// Validate if the game number already exists
		_, ok := unique[game.GameNumber]
		if ok {
			return ErrorDuplicatedGameIndex
		}

		// Validate each game
		err := game.Validate()
		if err != nil {
			return err
		}
		unique[game.GameNumber] = true
	}

	return nil
}
