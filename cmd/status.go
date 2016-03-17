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

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show changes in repository",
	Long: "Show changes in repository",
	Run: func(cmd *cobra.Command, args []string) {

		repo, err := s3git.OpenRepository(".")
		if err != nil {
			fmt.Println(err)
		}

		list, err := repo.Status()
		if err != nil {
			er(err)
		}

		elements := 0
		for elem := range list {
			if elements == 0 {
				fmt.Println("Changes to be committed:")
			}
			fmt.Println("\t", elem)
			elements++
		}

		if elements == 0 {
			fmt.Println("Nothing to commit, staging directory clean")
		}
	},
}

func init() {
	RootCmd.AddCommand(statusCmd)
}
