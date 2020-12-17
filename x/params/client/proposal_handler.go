package client

import (
	govclient "github.com/libonomy/cusp-sdk/x/gov/client"
	"github.com/libonomy/cusp-sdk/x/params/client/cli"
	"github.com/libonomy/cusp-sdk/x/params/client/rest"
)

// param change proposal handler
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
