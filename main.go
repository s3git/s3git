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

package main

import (
	"fmt"
	"github.com/s3git/s3git/cmd"
	"github.com/spf13/viper"
	"os"
)

var version = ""

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" && version != "" {
		fmt.Println(version)
		os.Exit(0)
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("s3git")
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath("/etc/s3git/")
	viper.AddConfigPath("$HOME/.s3git")

	// AWS_PROFILE or S3GIT_PROFILE can set the profile
	viper.BindEnv("profile", "AWS_PROFILE");

	for _, ext := range []string{"json", "yaml", "yml"} {
		viper.SetConfigType(ext)
		err := viper.ReadInConfig() // Find and read config files
		if err != nil { // Handle errors reading the config file
			if _, ok := err.(viper.ConfigParseError); ok {
				panic(fmt.Errorf("Fatal error reading " + ext + " config file: %s \n", err))
			}
		}
	}

	for _, extrafile := range []string{"$HOME/.aws/config", "$HOME/.aws/credentials"} {

		viper.SetConfigType("yaml") // TODO: this should be INI
		viper.SetConfigFile(extrafile) // name of config file (without extension)
		err := viper.ReadInConfig() // Find and read the config file

		if err != nil { // Handle errors reading the config file
			if _, ok := err.(viper.ConfigParseError); ok {
				panic(fmt.Errorf("Fatal error reading config file: %s \n", err))
			}
		}
	}

	//TODO: Investigate: do we set the max procs somewhere?
	cmd.Execute()
}
