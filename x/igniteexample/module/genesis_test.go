package igniteexample_test

import (
	"testing"

	keepertest "ignite-example/testutil/keeper"
	"ignite-example/testutil/nullify"
	igniteexample "ignite-example/x/igniteexample/module"
	"ignite-example/x/igniteexample/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IgniteexampleKeeper(t)
	igniteexample.InitGenesis(ctx, k, genesisState)
	got := igniteexample.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
