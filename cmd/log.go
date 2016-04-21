/*
 * Copyright 2016 Frank Wessels <fwessels@xs4all.nl>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"fmt"
	"github.com/s3git/s3git-go"
	"github.com/spf13/cobra"
	"github.com/fatih/color"
)

var oneline bool
var snapshots bool

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show commit log",
	Long: "Show commit log for the repository",
	Run: func(cmd *cobra.Command, args []string) {

		repo, err := s3git.OpenRepository(".")
		if err != nil {
			er(err)
		}

		options := []s3git.ListCommitOptions{}
		options = append(options, s3git.ListCommitOptionSetOnlySnapshots(true))

		list, err := repo.ListCommits("", options...)
		if err != nil {
			er(err)
		}

		for commit := range list {
			if oneline {
				color.Set(color.FgYellow)
				fmt.Print(commit.Hash)
				color.Unset()
				fmt.Print(" ")
				fmt.Println(commit.Message)

			} else {
				color.Set(color.FgYellow)
				fmt.Println("commit", commit.Hash)
				color.Unset()


				fmt.Println("Date:", commit.TimeStamp)
				fmt.Println()
				fmt.Println("   ", commit.Message)
				fmt.Println()
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(logCmd)

	// Add local message flags
	logCmd.Flags().BoolVarP(&oneline, "pretty", "p", false, "Pretty format")
	logCmd.Flags().BoolVar(&snapshots, "snapshots", false, "Just show snapshot commits")
}
