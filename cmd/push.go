package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Update remote repositories",
	Long: "Update remote repositories",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("push called")
	},
}

func init() {
	RootCmd.AddCommand(pushCmd)
}
