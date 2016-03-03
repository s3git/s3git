package cmd

import (
	"fmt"

	"github.com/s3git/s3git-go"
	"github.com/cheggaaa/pb"
	"github.com/spf13/cobra"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a repository into a new directory",
	Long: "Clone a repository into a new directory",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Cloning into ...")

		var barDownloading, barProcessing *pb.ProgressBar

		progressDownload := func(total int64) {
			if barDownloading == nil {
				barDownloading = pb.New64(total).Start()
				barDownloading.Prefix("Downloading ")
			}
			if barDownloading.Increment() == int(total) {
				barDownloading.Finish()
			}
		}

		progressProcessing := func(total int64) {
			if barProcessing == nil {
				barProcessing = pb.New64(total).Start()
				barProcessing.Prefix("Processing  ")
			}
			if barProcessing.Increment() == int(total) {
				barProcessing.Finish()
			}
		}

		_, err := s3git.Clone("s3://s3git-100m", ".", progressDownload, progressProcessing)
		if err != nil {
			er(err)
		}

		fmt.Println("Done. Totalling 97974067 objects")
	},
}

func init() {
	RootCmd.AddCommand(cloneCmd)
}