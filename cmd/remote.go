package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// remoteCmd represents the remote command
var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Manage remote reposities",
	Long: "Manage remote reposities",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("remote called")

		subCommand := "add"
		if subCommand == "add" {

		} else if subCommand == "show" {

		}
	},
}

func init() {
	RootCmd.AddCommand(remoteCmd)
}
