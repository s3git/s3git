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

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List files in the repository",
	Long: "List files in the repository",
	Run: func(cmd *cobra.Command, args []string) {

		repo, err := s3git.OpenRepository(".")
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
