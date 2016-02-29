package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Update local repository",
	Long: "Update local repository",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("pull called")
	},
}

func init() {
	RootCmd.AddCommand(pullCmd)
}
