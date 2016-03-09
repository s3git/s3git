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
	"io"
	"os"

	"github.com/s3git/s3git-go"
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

		repo, err := s3git.OpenRepository(".")
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
