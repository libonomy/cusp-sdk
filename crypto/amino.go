package crypto

import (
	amino "github.com/evdatsion/go-amino"
	cryptoAmino "github.com/libonomy/aphelion-staking/crypto/encoding/amino"
)

var cdc = amino.NewCodec()

func init() {
	RegisterAmino(cdc)
	cryptoAmino.RegisterAmino(cdc)
}

// RegisterAmino registers all go-crypto related types in the given (amino) codec.
func RegisterAmino(cdc *amino.Codec) {
	cdc.RegisterConcrete(PrivKeyLedgerSecp256k1{},
		"aphelion/PrivKeyLedgerSecp256k1", nil)
}
