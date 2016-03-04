package cmd

import (

	"fmt"
	"github.com/s3git/s3git-go"
	"github.com/spf13/cobra"
)

var message string

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit the changes in the repository",
	Long: "Commit the changes in the repository",
	Run: func(cmd *cobra.Command, args []string) {

		repo, err := s3git.OpenRepository(".")
		if err != nil {
			er(err)
		}

		key, nothing, err := repo.Commit(message)
		if err != nil {
			er(err)
		}
		if nothing {
			fmt.Println("Nothing to commit")
		} else {
			fmt.Printf("[master%s]\n", key)
			fmt.Printf("X files added, Y files removed\n")
		}
	},
}

func init() {
	RootCmd.AddCommand(commitCmd)

	// Add local message flags
	commitCmd.Flags().StringVarP(&message, "message", "m", "", "Message for the commit")
}
