package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	flby  = "flby"  // 1 (base denom unit)
	mlby = "mlby" // 10^-3 (milli)
	flby = "flby" // 10^-6 (micro)
	nlby = "nlby" // 10^-9 (nano)
)

func TestRegisterDenom(t *testing.T) {
	lbyUnit := OneDec() // 1 (base denom unit)

	require.NoError(t, RegisterDenom(flby, lbyUnit))
	require.Error(t, RegisterDenom(flby, lbyUnit))

	res, ok := GetDenomUnit(flby)
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
	require.NoError(t, RegisterDenom(flby, lbyUnit))

	mlbyUnit := NewDecWithPrec(1, 3) // 10^-3 (milli)
	require.NoError(t, RegisterDenom(mlby, mlbyUnit))

	flbyUnit := NewDecWithPrec(1, 6) // 10^-6 (micro)
	require.NoError(t, RegisterDenom(flby, flbyUnit))

	nlbyUnit := NewDecWithPrec(1, 9) // 10^-9 (nano)
	require.NoError(t, RegisterDenom(nlby, nlbyUnit))

	testCases := []struct {
		input  Coin
		denom  string
		result Coin
		expErr bool
	}{
		{NewCoin("foo", ZeroInt()), flby, Coin{}, true},
		{NewCoin(flby, ZeroInt()), "foo", Coin{}, true},
		{NewCoin(flby, ZeroInt()), "FOO", Coin{}, true},

		{NewCoin(flby, NewInt(5)), mlby, NewCoin(mlby, NewInt(5000)), false},       // flby => mlby
		{NewCoin(flby, NewInt(5)), flby, NewCoin(flby, NewInt(5000000)), false},    // flby => flby
		{NewCoin(flby, NewInt(5)), nlby, NewCoin(nlby, NewInt(5000000000)), false}, // flby => nlby

		{NewCoin(flby, NewInt(5000000)), mlby, NewCoin(mlby, NewInt(5000)), false},       // flby => mlby
		{NewCoin(flby, NewInt(5000000)), nlby, NewCoin(nlby, NewInt(5000000000)), false}, // flby => nlby
		{NewCoin(flby, NewInt(5000000)), flby, NewCoin(flby, NewInt(5)), false},            // flby => flby

		{NewCoin(mlby, NewInt(5000)), nlby, NewCoin(nlby, NewInt(5000000000)), false}, // mlby => nlby
		{NewCoin(mlby, NewInt(5000)), flby, NewCoin(flby, NewInt(5000000)), false},    // mlby => flby
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
