package types

import (
	"github.com/libonomy/cusp-sdk/codec"
)

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateValidator{}, "cusp-sdk/MsgCreateValidator", nil)
	cdc.RegisterConcrete(MsgEditValidator{}, "cusp-sdk/MsgEditValidator", nil)
	cdc.RegisterConcrete(MsgDelegate{}, "cusp-sdk/MsgDelegate", nil)
	cdc.RegisterConcrete(MsgUndelegate{}, "cusp-sdk/MsgUndelegate", nil)
	cdc.RegisterConcrete(MsgBeginRedelegate{}, "cusp-sdk/MsgBeginRedelegate", nil)
}

// generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
