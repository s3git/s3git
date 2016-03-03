package cmd

import (
	"fmt"

	"github.com/s3git/s3git-go"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show changes to repository",
	Long: "Show changes to repository",
	Run: func(cmd *cobra.Command, args []string) {

		repo, err := s3git.OpenRepository(".")
		if err != nil {
			fmt.Println(err)
		}

		list, err := repo.Status()
		if err != nil {
			er(err)
		}

		for elem := range list {
			fmt.Println(elem)
		}

	},
}

func init() {
	RootCmd.AddCommand(statusCmd)
}
