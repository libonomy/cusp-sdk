package bank

import (
	"strings"
	"testing"

	sdk "github.com/libonomy/cusp-sdk/types"
	abci "github.com/libonomy/aphelion-staking/abci/types"

	"github.com/stretchr/testify/require"
)

func TestInvalidMsg(t *testing.T) {
	h := NewHandler(nil)

	res := h(sdk.NewContext(nil, abci.Header{}, false, nil), sdk.NewTestMsg())
	require.False(t, res.IsOK())
	require.True(t, strings.Contains(res.Log, "unrecognized bank message type"))
}
