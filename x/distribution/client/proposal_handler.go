package client

import (
	"github.com/evdatsion/cosmos-sdk/x/distribution/client/cli"
	"github.com/evdatsion/cosmos-sdk/x/distribution/client/rest"
	govclient "github.com/evdatsion/cosmos-sdk/x/gov/client"
)

// param change proposal handler
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
