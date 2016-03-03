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

		repo, err := s3git.OpenRepository("/Users/frankw/golang/src/github.com/fwessels/ldcli-test/alice/lifedrive-100m-euc1-objs")
		if err != nil {
			er(err)
		}

		arg := ""
		if len(args) > 0 {
			arg = args[0]
		}
		list, err := repo.List(arg)
		if err != nil {
			er(err)
		}

		for elem := range list {
			fmt.Println(elem)
		}
	},
}

func init() {
	RootCmd.AddCommand(lsCmd)
}
