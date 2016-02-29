package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show changes to repository",
	Long: "Show changes to repository",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("status called")
	},
}

func init() {
	RootCmd.AddCommand(statusCmd)
}
