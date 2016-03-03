package cmd

import (
	"fmt"
	"github.com/fwessels/s3git-go"
	"github.com/spf13/cobra"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show commit logs",
	Long: "Show commit logs for the repository",
	Run: func(cmd *cobra.Command, args []string) {

		repo, err := s3git.OpenRepository(".")
		if err != nil {
			er(err)
		}

		list, err := repo.ListCommits()
		if err != nil {
			er(err)
		}

		for commit := range list {
			fmt.Println(commit.Hash, commit.Message)
		}
	},
}

func init() {
	RootCmd.AddCommand(logCmd)
}
