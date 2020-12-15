package gov

import (
	"strings"
	"testing"

	sdk "github.com/evdatsion/cosmos-sdk/types"
	abci "github.com/evdatsion/tendermint/abci/types"

	"github.com/stretchr/testify/require"
)

func TestInvalidMsg(t *testing.T) {
	k := Keeper{}
	h := NewHandler(k)

	res := h(sdk.NewContext(nil, abci.Header{}, false, nil), sdk.NewTestMsg())
	require.False(t, res.IsOK())
	require.True(t, strings.Contains(res.Log, "unrecognized gov message type"))
}
