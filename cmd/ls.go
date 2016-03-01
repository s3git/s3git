package cmd

import (
	"fmt"

	"github.com/fwessels/s3git-go"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List files in the repository",
	Long: "List files in the repository",
	Run: func(cmd *cobra.Command, args []string) {

		repo, err := s3git.OpenRepository(".")
		if err != nil {
			fmt.Println(err)
		}

		for elem := range repo.List() {
			fmt.Println(elem)
		}
	},
}

func init() {
	RootCmd.AddCommand(lsCmd)
}
