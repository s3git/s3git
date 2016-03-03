package cmd

import (
	"io"
	"os"

	"github.com/fwessels/s3git-go"
	"github.com/spf13/cobra"
)

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Read a file from the repository",
	Long: "Read a file from the repository and dump to stdout",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			er("Argument missing")
		}

		repo, err := s3git.OpenRepository("/Users/frankw/golang/src/github.com/fwessels/ldcli-test/alice/lifedrive-100m-euc1-objs")
		if err != nil {
			er(err)
		}

		hash, err := repo.MakeUnique(args[0])
		if err != nil {
			er(err)
		}

		r, err := repo.Get(hash)
		if err != nil {
			er(err)
		}

		io.Copy(os.Stdout, r)
	},
}

func init() {
	RootCmd.AddCommand(catCmd)
}
