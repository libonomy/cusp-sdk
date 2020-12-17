package client

import (
	"github.com/gorilla/mux"

	"github.com/libonomy/cusp-sdk/client/context"
)

// Register routes
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	RegisterRPCRoutes(cliCtx, r)
}
