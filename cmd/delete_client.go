package cmd

import (
	"code.cloudfoundry.org/uaa-cli/uaa"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var deleteClientCmd = &cobra.Command{
	Use:   "delete-client CLIENT_ID",
	Short: "Delete a client registration",
	PreRun: func(cmd *cobra.Command, args []string) {
		EnsureTarget()
	},
	Run: func(cmd *cobra.Command, args []string) {
		cm := &uaa.ClientManager{GetHttpClient(), GetSavedConfig()}
		_, err := cm.Delete(args[0])
		if err != nil {
			fmt.Println(err)
			TraceRetryMsg(GetSavedConfig())
			os.Exit(1)
		}

		fmt.Printf("Successfully deleted client %v.\n", args[0])
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return MissingArgument("client_id")
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(deleteClientCmd)
	deleteClientCmd.Flags().StringVarP(&zoneSubdomain, "zone", "z", "", "the identity zone subdomain in which to delete the client")
}