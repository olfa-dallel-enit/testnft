package keeper

import (
	"testnft/x/testnft/types"
)

var _ types.QueryServer = Keeper{}
