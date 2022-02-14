package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/mattverse/cardhaus/x/namecard/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateNamecard() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-namecard [address] [name] [image-url] [self-intro]",
		Short: "Broadcast message create-namecard",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress := args[0]
			argName := args[1]
			argImageUrl := args[2]
			argSelfIntro := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateNamecard(
				clientCtx.GetFromAddress().String(),
				argAddress,
				argName,
				argImageUrl,
				argSelfIntro,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
