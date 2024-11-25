package rpsKeeper

import (
	"context"
	"errors"

	"challenge/x/rps/types"

	"cosmossdk.io/collections"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = queryServer{}

type queryServer struct {
	k Keeper
}

func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return queryServer{
		k: keeper,
	}
}

func (qs queryServer) GetGame(ctx context.Context, req *types.QueryGetGameRequest) (*types.QueryGetGameResponse, error) {
	game, err := qs.k.Games.Get(ctx, req.Index)

	// If there is an error but is 'ErrNotFound' return empty
	if errors.Is(err, collections.ErrNotFound) {
		return &types.QueryGetGameResponse{Game: &types.Game{}}, nil
	}

	// Return other errors
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Return response
	return &types.QueryGetGameResponse{Game: &game}, nil

}
