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

		elements := 0
		for elem := range list {
			if elements == 0 {
				fmt.Println("Changes to be committed:")
			}
			fmt.Println("\t", elem)
			elements++
		}

		if elements == 0 {
			fmt.Println("Nothing to commit, staging directory clean")
		}
	},
}

func init() {
	RootCmd.AddCommand(statusCmd)
}
