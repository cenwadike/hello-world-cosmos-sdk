package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "ignite-example/testutil/keeper"
	"ignite-example/x/igniteexample/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.IgniteexampleKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
