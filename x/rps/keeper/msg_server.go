package rpsKeeper

import (
	"context"

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
func (ms msgServer) CreateStudent(ctx context.Context, msg *types.MsgCreateStudent) (*types.MsgCreateStudentResponse, error) {
	// TODO: Create function to create game in the blockchain
	return &types.MsgCreateStudentResponse{}, nil // Return the response as the interface needs (and proto file specify)
}

func (ms msgServer) MakeMove(ctx context.Context, msg *types.MsgMakeMove) (*types.MsgMakeMoveResponse, error) {

	return &types.MsgMakeMoveResponse{}, nil
}
