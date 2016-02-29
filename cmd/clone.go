package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a repository into a new directory",
	Long: "Clone a repository into a new directory",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("clone called")
	},
}

func init() {
	RootCmd.AddCommand(cloneCmd)
}
