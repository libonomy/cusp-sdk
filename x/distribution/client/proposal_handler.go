package client

import (
	"github.com/libonomy/cusp-sdk/x/distribution/client/cli"
	"github.com/libonomy/cusp-sdk/x/distribution/client/rest"
	govclient "github.com/libonomy/cusp-sdk/x/gov/client"
)

// param change proposal handler
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
