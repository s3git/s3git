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

var pushAfterCreate bool
var dedupe bool
var hash bool
var presignedUrls bool
var jsonOutput bool

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

		repo, err := s3git.OpenRepository(".")
		if err != nil {
			er(err)
		}

		key, nothing, err := repo.SnapshotCreate(args[0], message)
		if err != nil {
			er(err)
		}
		if nothing {
			fmt.Println("No changes to snapshot")
			return
		}

		fmt.Printf("[commit %s]\n", key)
		if pushAfterCreate {
			// TODO: Push new snapshot upstream
		}
	},
}

var snapshotCheckoutCmd = &cobra.Command{
	Use:   "checkout [directory] ([commit])",
	Short: "Checkout a snapshot",
	Long: "Checkout a snapshot",
	Run: func(cmd *cobra.Command, args []string) {

		// TODO: Partial checkout would be nice (eg specify path as filter)

		if len(args) == 0 {
			er("Directory for snapshot must be specified")
		}

		repo, err := s3git.OpenRepository(".")
		if err != nil {
			er(err)
		}

		var commit string
		if len(args) == 2 {
			commit = args[1]
		}

		err = repo.SnapshotCheckout(args[0], commit, dedupe)
		if err != nil {
			er(err)
		}
	},
}

var snapshotListCmd = &cobra.Command{
	Use:   "ls ([commit])",
	Short: "List a snapshot",
	Long: "List a snapshot",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			er("Commit for snapshot must be specified")
		}

		repo, err := s3git.OpenRepository(".")
		if err != nil {
			er(err)
		}

		var commit string
		if len(args) == 1 {
			commit = args[0]
		}

		options := []s3git.SnapshotListOptions{}
		options = append(options, s3git.SnapshotListOptionSetShowHash(hash))
		options = append(options, s3git.SnapshotListOptionSetPresignedUrls(presignedUrls))
		options = append(options, s3git.SnapshotListOptionSetJsonOutput(jsonOutput))

		err = repo.SnapshotList(commit, options...)
		if err != nil {
			er(err)
		}

	},
}

var snapshotStatusCmd = &cobra.Command{
	Use:   "status [directory] ([commit])",
	Short: "Show changes for snapshot",
	Long: "Show changes for snapshot",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			er("Directory for snapshot must be specified")
		}

		repo, err := s3git.OpenRepository(".")
		if err != nil {
			er(err)
		}

		var commit string
		if len(args) == 2 {
			commit = args[1]
		}

		err = repo.SnapshotStatus(args[0], commit)
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
	snapshotCmd.AddCommand(snapshotStatusCmd)

	// Local flags for create
	snapshotCreateCmd.Flags().StringVarP(&message, "message", "m", "", "Message for the commit of create snapshot")
	snapshotCreateCmd.Flags().BoolVarP(&pushAfterCreate, "push", "p", false, "Perform immediate push after create")

	// Local flags for checkout
	snapshotCheckoutCmd.Flags().BoolVar(&dedupe, "dedupe", false, "Checkout in deduped (pointers) format")

	// Local flags for list
	snapshotListCmd.Flags().BoolVar(&hash, "hash", false, "Show hash of object")
	snapshotListCmd.Flags().BoolVar(&presignedUrls, "presigned", false, "Generate presigned urls for direct access from S3")
	snapshotListCmd.Flags().BoolVar(&jsonOutput, "json", false, "Output result in JSON")
}
