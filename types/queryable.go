package types

import abci "github.com/libonomy/aphelion-staking/abci/types"

// Type for querier functions on keepers to implement to handle custom queries
type Querier = func(ctx Context, path []string, req abci.RequestQuery) (res []byte, err Error)
