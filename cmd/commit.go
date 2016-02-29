package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit the changes in the repository",
	Long: "Commit the changes in the repository",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("commit called")
	},
}

func init() {
	RootCmd.AddCommand(commitCmd)

	// Add local message flags
	commitCmd.Flags().StringP("message", "m", "", "Message for the commit")
}
