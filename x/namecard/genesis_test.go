package namecard_test

import (
	"testing"

	keepertest "github.com/mattverse/cardhaus/testutil/keeper"
	"github.com/mattverse/cardhaus/testutil/nullify"
	"github.com/mattverse/cardhaus/x/namecard"
	"github.com/mattverse/cardhaus/x/namecard/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NamecardKeeper(t)
	namecard.InitGenesis(ctx, *k, genesisState)
	got := namecard.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
