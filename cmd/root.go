package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(-1)
}

// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "s3git",
	Short: "git for Cloud Storage",
	Long: `s3git applies the git philosophy to Cloud Storage. If you know git, you will know how to use s3git.

s3git is a simple CLI tool that allows you to create a distributed, decentralized and versioned repository.
It scales limitlessly to 100s of millions of files and PBs of storage and stores your data safely in S3.
Yet huge repos can be cloned on the SSD of your laptop for making local changes, committing and pushing back.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		er(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file
func initConfig() {
}
