package cli

import (
    "strconv"
	
	"github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/charleenfei/cosmoverse-workshop/x/eightball/types"
)

var _ = strconv.Itoa(0)

func CmdConnectToDex() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "connect-to-dex [connection-id]",
		Short: "Broadcast message connect-to-dex",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
      		 argConnectionId := args[0]
            
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgConnectToDex(
				clientCtx.GetFromAddress().String(),
				argConnectionId,
				
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