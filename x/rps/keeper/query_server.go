package rpsKeeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	"github.com/0xlb/rps-chain/x/rps/types"
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

func (qs queryServer) GetParams(ctx context.Context, req *types.QueryGetParams) (*types.QueryGetParamsResponse, error) {
	params, err := qs.k.Params.Get(ctx)
	// If there is an error but is 'ErrNotFound' return empty
	if errors.Is(err, collections.ErrNotFound) {
		return &types.QueryGetParamsResponse{Param: &types.Params{}}, nil
	}

	// Return other errors
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Return response
	return &types.QueryGetParamsResponse{Param: &params}, nil
}
