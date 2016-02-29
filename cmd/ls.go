package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List files in the repository",
	Long: "List files in the repository",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("ls called")
	},
}

func init() {
	RootCmd.AddCommand(lsCmd)
}
