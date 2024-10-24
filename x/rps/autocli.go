package rps

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	rpsv1 "github.com/0xlb/rps-chain/api/rps/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: rpsv1.Msg_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "CreateGame",                // This is the name that I want for the command
					Use:       "create [oponent] [rounds]", // This is the command I want
					Short:     "Creates a new Rock, Paper and Scissors for the message sender and the chose oponent",
					Long:      "Creates a new Rock, Paper and Scissors for the message sender and the chose oponent. Input parameters are the oponent address and the rounds numebr for the game",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "oponent"}, // This must be name as the proto field
						{ProtoField: "rounds"},  // This must be name as the proto field
					},
				},
				{
					RpcMethod: "MakeMove",
					Use:       "make [game_index] [movement]", // This is the command I want
					Short:     "Make a Rock, Paper, Scissors movement",
					Long:      "Make a Rock, Paper, Scissors movement. Valid move options are 'Rock' - 'Paper' - 'Scissors'",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "game_index"}, // This must be name as the proto field
						{ProtoField: "move"},       // This must be name as the proto field
					},
				},
				{
					RpcMethod: "ReviewMove",
					Use:       "review-move [game_index] [revealed_move] [salt]",
					Short:     "Reveals a submitted commitment for a specific Rock, Paper & Scissors Game",
					Long:      "Reveals a submitted commitment for a specific Rock, Paper & Scissors Game",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "game_index"},
						{ProtoField: "revealed_move"},
						{ProtoField: "salt"},
					},
				},
			},
		},
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: rpsv1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "GetGame",
					Use:       "game [game_index]", // This is the command I want
					Short:     "Get the current game with the provided game index",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "index"}, // This must be name as the proto field
					},
				},
				{
					RpcMethod: "GetParams",
					Use:       "params",
					Short:     "Get the current params",
				},
			},
		},
	}
}
