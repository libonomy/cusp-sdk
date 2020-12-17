package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	libocoin  = "libocoin"  // 1 (base denom unit)
	mlibocoin = "mlibocoin" // 10^-3 (milli)
	ulibocoin = "ulibocoin" // 10^-6 (micro)
	nlibocoin = "nlibocoin" // 10^-9 (nano)
)

func TestRegisterDenom(t *testing.T) {
	libocoinUnit := OneDec() // 1 (base denom unit)

	require.NoError(t, RegisterDenom(libocoin, libocoinUnit))
	require.Error(t, RegisterDenom(libocoin, libocoinUnit))

	res, ok := GetDenomUnit(libocoin)
	require.True(t, ok)
	require.Equal(t, libocoinUnit, res)

	res, ok = GetDenomUnit(mlibocoin)
	require.False(t, ok)
	require.Equal(t, ZeroDec(), res)

	// reset registration
	denomUnits = map[string]Dec{}
}

func TestConvertCoins(t *testing.T) {
	libocoinUnit := OneDec() // 1 (base denom unit)
	require.NoError(t, RegisterDenom(libocoin, libocoinUnit))

	mlibocoinUnit := NewDecWithPrec(1, 3) // 10^-3 (milli)
	require.NoError(t, RegisterDenom(mlibocoin, mlibocoinUnit))

	ulibocoinUnit := NewDecWithPrec(1, 6) // 10^-6 (micro)
	require.NoError(t, RegisterDenom(ulibocoin, ulibocoinUnit))

	nlibocoinUnit := NewDecWithPrec(1, 9) // 10^-9 (nano)
	require.NoError(t, RegisterDenom(nlibocoin, nlibocoinUnit))

	testCases := []struct {
		input  Coin
		denom  string
		result Coin
		expErr bool
	}{
		{NewCoin("foo", ZeroInt()), libocoin, Coin{}, true},
		{NewCoin(libocoin, ZeroInt()), "foo", Coin{}, true},
		{NewCoin(libocoin, ZeroInt()), "FOO", Coin{}, true},

		{NewCoin(libocoin, NewInt(5)), mlibocoin, NewCoin(mlibocoin, NewInt(5000)), false},       // libocoin => mlibocoin
		{NewCoin(libocoin, NewInt(5)), ulibocoin, NewCoin(ulibocoin, NewInt(5000000)), false},    // libocoin => ulibocoin
		{NewCoin(libocoin, NewInt(5)), nlibocoin, NewCoin(nlibocoin, NewInt(5000000000)), false}, // libocoin => nlibocoin

		{NewCoin(ulibocoin, NewInt(5000000)), mlibocoin, NewCoin(mlibocoin, NewInt(5000)), false},       // ulibocoin => mlibocoin
		{NewCoin(ulibocoin, NewInt(5000000)), nlibocoin, NewCoin(nlibocoin, NewInt(5000000000)), false}, // ulibocoin => nlibocoin
		{NewCoin(ulibocoin, NewInt(5000000)), libocoin, NewCoin(libocoin, NewInt(5)), false},            // ulibocoin => libocoin

		{NewCoin(mlibocoin, NewInt(5000)), nlibocoin, NewCoin(nlibocoin, NewInt(5000000000)), false}, // mlibocoin => nlibocoin
		{NewCoin(mlibocoin, NewInt(5000)), ulibocoin, NewCoin(ulibocoin, NewInt(5000000)), false},    // mlibocoin => ulibocoin
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
