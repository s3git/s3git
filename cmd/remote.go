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

// remoteCmd represents the remote command
var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Manage remote reposities",
	Long: "Manage remote reposities",
}

var remoteAddCmd = &cobra.Command{
	Use:   "add [name]",
	Short: "Add a remote repository",
	Long: "Add a remote repository",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			er("Name for remote must be specified")
		}

		repo, err := s3git.OpenRepository(".")
		if err != nil {
			er(err)
		}

		if resource == "" {
			er("Name for remote must be specified")
		} else {
			err := repo.RemoteAdd(args[0], resource, accessKey, secretKey)
			if err != nil {
				er(err)
			}
		}
	},
}

var remoteShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show remote repositories",
	Long: "Show remote repositories",
	Run: func(cmd *cobra.Command, args []string) {

		repo, err := s3git.OpenRepository(".")
		if err != nil {
			er(err)
		}

		remotes, err := repo.RemotesShow()
		if err != nil {
			er(err)
		}

		for _, r := range remotes {
			fmt.Println(r.Name, r.Resource)
		}
	},
}

func init() {
	RootCmd.AddCommand(remoteCmd)
	remoteCmd.AddCommand(remoteAddCmd)
	remoteCmd.AddCommand(remoteShowCmd)

	// Add local message flags
	remoteAddCmd.Flags().StringVarP(&resource, "resource", "r", "", "URL for remote S3")
	remoteAddCmd.Flags().StringVarP(&accessKey, "access", "a", "", "Access key for remote S3")
	remoteAddCmd.Flags().StringVarP(&secretKey, "secret", "s", "", "Secret key for remote S3")
}

