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
	"path/filepath"
	"github.com/spf13/cobra"
	"github.com/s3git/s3git-go"
)

var resource string
var accessKey string
var secretKey string
var endpoint string

var leafSize uint32
var maxRepoSize uint64

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create an empty repository",
	Long: "Create an empty repository",
	Run: func(cmd *cobra.Command, args []string) {

		// TODO: Check whether not already initialized -- abort if so

		var dir string
		var err error
		if len(args) > 0 {
			dir, err = filepath.Abs(filepath.Dir(args[0]))
		} else {
			dir, err = filepath.Abs(filepath.Dir("."))
		}
		if err != nil {
			er(err)
		}

		options := []s3git.InitOptions{}
		options = append(options, s3git.InitOptionSetLeafSize(leafSize))
		options = append(options, s3git.InitOptionSetMaxRepoSize(maxRepoSize))

		repo, err := s3git.InitRepository(dir, options...)
		if err != nil {
			er(err)
		}

		fmt.Printf("Initialized empty s3git repository in %s\n", dir)

		// Add remote when resource specifier is not empty (access & secret may be omitted for public access)
		if resource != "" {
			err := repo.RemoteAdd("primary", resource, accessKey, secretKey)
			if err != nil {
				er(err)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(initCmd)

	// Add local message flags
	initCmd.Flags().Uint32VarP(&leafSize, "leafsize", "l", 5*1024*1024, "Leaf size for nodes")
	initCmd.Flags().Uint64VarP(&maxRepoSize, "maxreposize", "m", 25*1024*1024*1024, "Maximum local disk space that repository will consume")
}
