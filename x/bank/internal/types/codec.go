package types

import (
	"github.com/libonomy/cusp-sdk/codec"
)

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSend{}, "cusp-sdk/MsgSend", nil)
	cdc.RegisterConcrete(MsgMultiSend{}, "cusp-sdk/MsgMultiSend", nil)
}

// module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}
