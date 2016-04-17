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
	"github.com/s3git/s3git-go"
	"github.com/spf13/cobra"
)

// snapshotCmd represents the snapshot command
var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Manage snapshots",
	Long: "Create, checkout and list snapshots",
}

var snapshotCreateCmd = &cobra.Command{
	Use:   "create [directory]",
	Short: "Create a snapshot",
	Long: "Create a snapshot",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			er("Directory for snapshot must be specified")
		} else if message == "" {
			er("Commit message for snapshot must be specified")
		}

		_/*repo*/, err := s3git.OpenRepository(".")
		if err != nil {
			er(err)
		}

		//err = repo.SnapshotCreate()
		//if err != nil {
		//	er(err)
		//}
	},
}

var snapshotCheckoutCmd = &cobra.Command{
	Use:   "checkout [commit] [directory]",
	Short: "Checkout a snapshot",
	Long: "Checkout a snapshot",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			er("Commit for snapshot must be specified")
		}

		_ /*repo*/, err := s3git.OpenRepository(".")
		if err != nil {
			er(err)
		}

		//err = repo.RemoteRemove(args[0])
		//if err != nil {
		//	er(err)
		//}
	},
}

var snapshotListCmd = &cobra.Command{
	Use:   "list [commit]",
	Short: "List a snapshot",
	Long: "List a snapshot",
	Run: func(cmd *cobra.Command, args []string) {

		_/*repo*/, err := s3git.OpenRepository(".")
		if err != nil {
			er(err)
		}
	},
}

var snapshotLogCmd = &cobra.Command{
	Use:   "log",
	Short: "Show commit log for snapshots",
	Long: "Show commit log for snapshots",
	Run: func(cmd *cobra.Command, args []string) {

		_/*repo*/, err := s3git.OpenRepository(".")
		if err != nil {
			er(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(snapshotCmd)
	snapshotCmd.AddCommand(snapshotCreateCmd)
	snapshotCmd.AddCommand(snapshotCheckoutCmd)
	snapshotCmd.AddCommand(snapshotListCmd)
	snapshotCmd.AddCommand(snapshotLogCmd)

	// Add local message flags
	snapshotCreateCmd.Flags().StringVarP(&message, "message", "m", "", "Message for the commit of create snapshot")
}
