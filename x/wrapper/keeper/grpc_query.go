package keeper

import (
	"github.com/charleenfei/cosmoverse-workshop/x/wrapper/types"
)

var _ types.QueryServer = Keeper{}
