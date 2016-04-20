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
	"github.com/cheggaaa/pb"
)

var hydrate bool

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Update remote repositories",
	Long: "Update remote repositories",
	Run: func(cmd *cobra.Command, args []string) {

		repo, err := s3git.OpenRepository(".")
		if err != nil {
			er(err)
		}

		var barPushing *pb.ProgressBar

		progressPush := func(total int64) {
			if barPushing == nil {
				barPushing = pb.New64(total).Start()
				barPushing.Prefix("Pushing ")
			}
			if barPushing.Increment() == int(total) {
				barPushing.Finish()
			}
		}

		err = repo.Push(hydrate, progressPush)
		if err != nil {
			er(err)
		}

		if barPushing == nil {
			fmt.Println("Everything up-to-date.")
		}
	},
}

func init() {
	RootCmd.AddCommand(pushCmd)

	// Add local message flags
	pushCmd.Flags().BoolVar(&hydrate, "hydrate", false, "Store in hydrated (original) format at remote")
}
