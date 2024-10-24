package types

import (
	types "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterInterfaces registers the interfaces types with the interface registry.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{}, // Register my MsgUpdateParams struct
		&MsgCreateGame{},   // Register my MsgCreateGame struct
		&MsgMakeMove{},     // Register my MsgMakeMove struct
		&MsgReviewMove{},   // Register my MsgReviewMove struct
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
