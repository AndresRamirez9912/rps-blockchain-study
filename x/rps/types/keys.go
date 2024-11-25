package types

import "cosmossdk.io/collections"

const ModuleName = "rps"

// Here store the stateless collection keys
var (
	ParamsKey        = collections.NewPrefix(0)
	StudentsKey      = collections.NewPrefix(1)
	StudentNumberKey = collections.NewPrefix(2)
)
