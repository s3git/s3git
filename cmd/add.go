package cmd

import (
	"os"
	"path/filepath"

	"github.com/fwessels/s3git-go"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add stream or file(s) to the repository",
	Long: "Add a stream or one or more file(s) to the repository",
	Run: func(cmd *cobra.Command, args []string) {

		repo, err := s3git.OpenRepository(".")
		if err != nil {
			er(err)
		}

		if len(args) == 0 {
			// Read from stdin
			repo.Add(os.Stdin)
		} else {
			// Iterate over file list
			// TODO: Add support for '...' operator (** wildcard)?
			for _, pattern := range args {
				fileList, err := filepath.Glob(pattern)
				if err != nil {
					er(err)
				}

				for _, filename := range fileList {
					file, err := os.Open(filename)
					if err != nil {
						er(err)
					}

					_, err = repo.Add(file)
					if err != nil {
						er(err)
					}
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
