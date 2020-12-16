package client

import (
	"github.com/evdatsion/cusp-sdk/x/distribution/client/cli"
	"github.com/evdatsion/cusp-sdk/x/distribution/client/rest"
	govclient "github.com/evdatsion/cusp-sdk/x/gov/client"
)

// param change proposal handler
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
