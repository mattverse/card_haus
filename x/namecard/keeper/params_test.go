package keeper_test

import (
	"testing"

	testkeeper "github.com/mattverse/cardhaus/testutil/keeper"
	"github.com/mattverse/cardhaus/x/namecard/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NamecardKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
