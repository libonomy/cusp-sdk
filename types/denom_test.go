package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	lby  = "lby"  // 1 (base denom unit)
	mlby = "mlby" // 10^-3 (milli)
	ulby = "ulby" // 10^-6 (micro)
	nlby = "nlby" // 10^-9 (nano)
)

func TestRegisterDenom(t *testing.T) {
	lbyUnit := OneDec() // 1 (base denom unit)

	require.NoError(t, RegisterDenom(lby, lbyUnit))
	require.Error(t, RegisterDenom(lby, lbyUnit))

	res, ok := GetDenomUnit(lby)
	require.True(t, ok)
	require.Equal(t, lbyUnit, res)

	res, ok = GetDenomUnit(mlby)
	require.False(t, ok)
	require.Equal(t, ZeroDec(), res)

	// reset registration
	denomUnits = map[string]Dec{}
}

func TestConvertCoins(t *testing.T) {
	lbyUnit := OneDec() // 1 (base denom unit)
	require.NoError(t, RegisterDenom(lby, lbyUnit))

	mlbyUnit := NewDecWithPrec(1, 3) // 10^-3 (milli)
	require.NoError(t, RegisterDenom(mlby, mlbyUnit))

	ulbyUnit := NewDecWithPrec(1, 6) // 10^-6 (micro)
	require.NoError(t, RegisterDenom(ulby, ulbyUnit))

	nlbyUnit := NewDecWithPrec(1, 9) // 10^-9 (nano)
	require.NoError(t, RegisterDenom(nlby, nlbyUnit))

	testCases := []struct {
		input  Coin
		denom  string
		result Coin
		expErr bool
	}{
		{NewCoin("foo", ZeroInt()), lby, Coin{}, true},
		{NewCoin(lby, ZeroInt()), "foo", Coin{}, true},
		{NewCoin(lby, ZeroInt()), "FOO", Coin{}, true},

		{NewCoin(lby, NewInt(5)), mlby, NewCoin(mlby, NewInt(5000)), false},       // lby => mlby
		{NewCoin(lby, NewInt(5)), ulby, NewCoin(ulby, NewInt(5000000)), false},    // lby => ulby
		{NewCoin(lby, NewInt(5)), nlby, NewCoin(nlby, NewInt(5000000000)), false}, // lby => nlby

		{NewCoin(ulby, NewInt(5000000)), mlby, NewCoin(mlby, NewInt(5000)), false},       // ulby => mlby
		{NewCoin(ulby, NewInt(5000000)), nlby, NewCoin(nlby, NewInt(5000000000)), false}, // ulby => nlby
		{NewCoin(ulby, NewInt(5000000)), lby, NewCoin(lby, NewInt(5)), false},            // ulby => lby

		{NewCoin(mlby, NewInt(5000)), nlby, NewCoin(nlby, NewInt(5000000000)), false}, // mlby => nlby
		{NewCoin(mlby, NewInt(5000)), ulby, NewCoin(ulby, NewInt(5000000)), false},    // mlby => ulby
	}

	for i, tc := range testCases {
		res, err := ConvertCoin(tc.input, tc.denom)
		require.Equal(
			t, tc.expErr, err != nil,
			"unexpected error; tc: #%d, input: %s, denom: %s", i+1, tc.input, tc.denom,
		)
		require.Equal(
			t, tc.result, res,
			"invalid result; tc: #%d, input: %s, denom: %s", i+1, tc.input, tc.denom,
		)
	}

	// reset registration
	denomUnits = map[string]Dec{}
}
