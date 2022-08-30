package keeper

import (
	"github.com/charleenfei/cosmoverse-workshop/x/eightball/types"
)

var _ types.QueryServer = Keeper{}
