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

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Update local repository",
	Long: "Update local repository",
	Run: func(cmd *cobra.Command, args []string) {

		repo, err := s3git.OpenRepository(".")
		if err != nil {
			er(err)
		}

		var barPulling *pb.ProgressBar

		progressPull := func(total int64) {
			if barPulling == nil {
				barPulling = pb.New64(total).Start()
				barPulling.Prefix("Pulling ")
			}
			if barPulling.Increment() == int(total) {
				barPulling.Finish()
			}
		}

		err = repo.Pull(progressPull)
		if err != nil {
			er(err)
		}

		if barPulling == nil {
			fmt.Println("Already up-to-date.")
		}
	},
}

func init() {
	RootCmd.AddCommand(pullCmd)
}
