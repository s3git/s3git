package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add file(s) to the repository",
	Long: "Add file(s) to the repository",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("add called")
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
