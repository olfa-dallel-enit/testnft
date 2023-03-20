package testnft_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "testnft/testutil/keeper"
	"testnft/testutil/nullify"
	"testnft/x/testnft"
	"testnft/x/testnft/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TestnftKeeper(t)
	testnft.InitGenesis(ctx, *k, genesisState)
	got := testnft.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
