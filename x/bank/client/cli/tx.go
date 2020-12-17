package cli

import (
	"github.com/spf13/cobra"

	"github.com/evdatsion/cusp-sdk/client"
	"github.com/evdatsion/cusp-sdk/client/context"
	"github.com/evdatsion/cusp-sdk/codec"
	sdk "github.com/evdatsion/cusp-sdk/types"
	"github.com/evdatsion/cusp-sdk/x/auth"
	"github.com/evdatsion/cusp-sdk/x/auth/client/utils"
	"github.com/evdatsion/cusp-sdk/x/bank/internal/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Bank transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(
		SendTxCmd(cdc),
	)
	return txCmd
}

// SendTxCmd will create a send tx and sign it with the given key.
func SendTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send [from_key_or_address] [to_address] [amount]",
		Short: "Create and sign a send tx",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithFrom(args[0]).WithCodec(cdc)

			to, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			// parse coins trying to be sent
			coins, err := sdk.ParseCoins(args[2])
			if err != nil {
				return err
			}

			// build and sign the transaction, then broadcast to Aphelion
			msg := types.NewMsgSend(cliCtx.GetFromAddress(), to, coins)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd = client.PostCommands(cmd)[0]

	return cmd
}
