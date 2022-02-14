package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/mattverse/cardhaus/x/namecard/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCardInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "card-info [address]",
		Short: "Query card-info",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.CardInfo(cmd.Context(), &types.QueryCardInfoRequest{Address: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
