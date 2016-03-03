package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/fwessels/s3git-go"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create an empty repository",
	Long: "Create an empty repository",
	Run: func(cmd *cobra.Command, args []string) {

		repo, err := s3git.InitRepository(".")
		if err != nil {
			er(err)
		}

		repo.Remotes = append(repo.Remotes, s3git.Remote{}) // .Add("s3://mybucket", "ACIA1234", "SECRETKEY")
		fmt.Println(len(repo.Remotes))
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
