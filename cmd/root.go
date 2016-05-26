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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(-1)
}

var endpoint string
var accessKey string
var secretKey string

// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "s3git",
	Short: "git for Cloud Storage",
	Long: `s3git applies the git philosophy to Cloud Storage. If you know git, you will know how to use s3git.

s3git is a simple CLI tool that allows you to create a distributed, decentralized and versioned repository.
It scales limitlessly to 100s of millions of files and PBs of storage and stores your data safely in S3.
Yet huge repos can be cloned on the SSD of your laptop for making local changes, committing and pushing back.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		er(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	cmd := RootCmd
	cmd.PersistentFlags().StringVarP(&accessKey, "access", "a", "", "Access key for S3 remote")
	viper.BindPFlag("access", cmd.PersistentFlags().Lookup("access"))
	cmd.PersistentFlags().StringVarP(&secretKey, "secret", "s", "", "Secret key for S3 remote")
	viper.BindPFlag("secret", cmd.PersistentFlags().Lookup("secret"))
	cmd.PersistentFlags().StringVarP(&endpoint, "endpoint", "e", "", "Endpoint for S3 remote")
	viper.BindPFlag("endpoint", cloneCmd.PersistentFlags().Lookup("endpoint"))
}

// initConfig reads in config file
func initConfig() {
}
