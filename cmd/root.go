/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/slayer321/uzzo/util"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:               "uzzo",
	Short:             "uzzo (Unzip and Open) can be used to unzip and open the directory with any IDE",
	Long:              `📂 uzzo will help you to directly unzip and open the directory in the specified IDE or text editior you want.`,
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	Version:           "1.0.0",

	DisableFlagsInUseLine: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if File != "" && len(args) < 1 {
			return errors.New("accept(s) 1 argument")
		} else if len(args) == 0 {
			return errors.New("Enter the help flag -h or --help")
		}
		return nil
	},
	Example: `uzzo demo.zip`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var filename string
		var err error
		var argument string

		argument = args[0]

		fileExists, err := util.FileExists(argument)
		if err != nil {
			fmt.Println(err.Error())
		}
		// check if file exists or not
		if fileExists {
			filename, err = filepath.Abs(argument)
			if err != nil {
				fmt.Println(err.Error())
			}

		} else {
			fmt.Printf("File %v does not exist \n", argument)
			return
		}

		// get the working dir
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}

		// using Unzip function to Unzip the file
		util.Unzip(filename, wd)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".uzzo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".uzzo")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
