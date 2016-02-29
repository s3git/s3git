package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Read a file from the repository",
	Long: "Read a file from the repository and dump to stdout",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("cat called")
	},
}

func init() {
	RootCmd.AddCommand(catCmd)
}
