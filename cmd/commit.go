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
			fmt.Printf("[commit %s]\n", key)
			//fmt.Printf("X files added, Y files removed\n")
		}
	},
}

func init() {
	RootCmd.AddCommand(commitCmd)

	// Add local message flags
	commitCmd.Flags().StringVarP(&message, "message", "m", "", "Message for the commit")
	//commitCmd.Flags().StringVarP(&message, "all", "a", "", "Add all parent commits")

	// Use --all flags to include all top most commits as parents, or explicitly specify one or more on the command line
}
