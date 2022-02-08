package keeper

import (
	"github.com/mattverse/cardhaus/x/namecard/types"
)

var _ types.QueryServer = Keeper{}
