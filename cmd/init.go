package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/spf13/cobra"
	"github.com/s3git/s3git-go"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create an empty repository",
	Long: "Create an empty repository",
	Run: func(cmd *cobra.Command, args []string) {

		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			er(err)
		}

		_ /*repo*/, err = s3git.InitRepository(dir)
		if err != nil {
			er(err)
		}

		fmt.Printf("Initialized empty s3git repository in %s\n", dir)

//		repo.Remotes = append(repo.Remotes, s3git.Remote{}) // .Add("s3://mybucket", "ACIA1234", "SECRETKEY")
//		fmt.Println(len(repo.Remotes))
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
