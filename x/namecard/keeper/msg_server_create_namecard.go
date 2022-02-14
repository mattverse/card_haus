package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mattverse/cardhaus/x/namecard/types"
)

func (k msgServer) CreateNamecard(goCtx context.Context, msg *types.MsgCreateNamecard) (*types.MsgCreateNamecardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	var nameCard = types.NameCard{
		Address:   msg.Address,
		Name:      msg.Name,
		ImageUrl:  msg.ImageUrl,
		SelfIntro: msg.SelfIntro,
	}

	err := k.SetNameCard(ctx, &nameCard)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	return &types.MsgCreateNamecardResponse{}, nil
}
