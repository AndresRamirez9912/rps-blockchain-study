package types

import "cosmossdk.io/collections"

const ModuleName = "rps"

// Here store the stateless collection keys
var (
	ParamsKey     = collections.NewPrefix(0)
	GamesKey      = collections.NewPrefix(1)
	GameNumberKey = collections.NewPrefix(2)
)

/*
Storage root
│
├── ParamsKey (Prefix: 0)
│   ├── "MaxGamesAllowed" : 10
│   └── "GameTimeout" : 300
│
├── GamesKey (Prefix: 1)
│   ├── GameID: 1
│   │   ├── "Player1" : "addr1..."
│   │   ├── "Player2" : "addr2..."
│   │   ├── "Winner" : "addr1..."
│   │   └── "Moves" : ["rock", "paper"]
│   ├── GameID: 2
│   │   ├── "Player1" : "addr3..."
│   │   ├── "Player2" : "addr4..."
│   │   └── "Moves" : ["scissors", "rock"]
│   └── GameID: 3
│       ├── "Player1" : "addr5..."
│       └── "Player2" : "addr6..."
│
└── GameNumberKey (Prefix: 2)
    └── "TotalGames" : 3
*/
