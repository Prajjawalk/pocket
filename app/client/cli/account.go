package cli

import (
	"fmt"

	"github.com/pokt-network/pocket/shared/crypto"
	"github.com/pokt-network/pocket/utility/types"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(NewAccountCommand())
}

func NewAccountCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "Account",
		Short:   "Account specific commands",
		Aliases: []string{"account"},
		Args:    cobra.ExactArgs(0),
	}

	cmd.AddCommand(accountCommands()...)

	return cmd
}

func accountCommands() (cmds []*cobra.Command) {
	sendCmd := &cobra.Command{
		Use:     "Send <from> <to> <amount>",
		Short:   "Send <from> <to> <amount>",
		Long:    "Sends <amount> to address <to> from address <from>",
		Aliases: []string{"send"},
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO(pocket/issues/150): update when we have keybase
			pk, err := readEd25519PrivateKeyFromFile(privateKeyFilePath)
			if err != nil {
				return err
			}
			// currently ignored since we are using the address from the PrivateKey
			// fromAddr := crypto.AddressFromString(args[0])
			toAddr := crypto.AddressFromString(args[1])
			amount := args[2]

			_ = &types.MessageSend{
				FromAddress: pk.Address(),
				ToAddress:   toAddr,
				Amount:      amount,
			}

			// TODO(deblasis): implement RPC client, route and handler
			fmt.Printf("sending %s from %s to %s\n", args[2], args[0], args[1])
			return nil
		},
	}

	cmds = append(cmds, sendCmd)

	return cmds
}