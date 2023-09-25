package keeper

import (
	"core/x/blog/types"
)

var _ types.QueryServer = Keeper{}
