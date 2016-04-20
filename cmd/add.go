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
	"os"
	"fmt"
	"path/filepath"

	"github.com/s3git/s3git-go"
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
			key, newBlob, err := repo.Add(os.Stdin)
			if err != nil {
				er(err)
			}
			printKey(key, newBlob)

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
					defer file.Close()

					key, newBlob, err := repo.Add(file)
					if err != nil {
						er(err)
					}
					printKey(key, newBlob)
				}
			}
		}
	},
}

func printKey(key string, newBlob bool) {
	if newBlob {
		fmt.Println("Added:", key)
	} else {
		fmt.Println("Already in repo:", key)
	}
}

func init() {
	RootCmd.AddCommand(addCmd)
}
