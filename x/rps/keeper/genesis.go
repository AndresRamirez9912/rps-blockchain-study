package rpsKeeper

import (
	"context"

	"challenge/x/rps/types"
)

// InitGenesis initializes the module state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *types.GenesisState) error {

	// Set the genesis games into the state
	for _, student := range data.Students {
		err := k.Students.Set(ctx, student.Id, student)
		if err != nil {
			return nil
		}
	}

	return nil
}
