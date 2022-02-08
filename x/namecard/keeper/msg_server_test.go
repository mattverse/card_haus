package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/mattverse/cardhaus/testutil/keeper"
	"github.com/mattverse/cardhaus/x/namecard/keeper"
	"github.com/mattverse/cardhaus/x/namecard/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.NamecardKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
