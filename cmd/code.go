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
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/slayer321/uzzo/util"

	"github.com/spf13/cobra"
)

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "code <zip_file_path>",
	Short: "It will open the directory in Visual Studio Code",
	Long: `This command will help to open the unzipped folder
	to Visual Studio Code.
	In order for this command to work, Visual Studio code should be installed in your system`,

	DisableFlagsInUseLine: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if File != "" && len(args) < 1 {
			return errors.New("accept(s) 1 argument")
		} else if len(args) == 0 {
			return errors.New("Enter the help flag -h or --help")
		}
		return nil
	},
	Example: `uzzo code demo.zip
	          uzzo code /Downloads/application.zip`,
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
			//fmt.Printf("Inside %v\n", filename)
			//fmt.Printf("Inside %v\n", argument)
		} else {
			fmt.Printf("File %v does not exist \n", argument)
			return
		}

		// get the working dir
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}
		//fmt.Printf("Outside %v", filename)

		// using Unzip function to Unzip the file
		util.Unzip(filename, wd)

		//changing the dir and removing .zip extension
		os.Chdir(util.FilenameWithoutExtension(filename))

		// getting the working dir after changing the dir
		wd, err = os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		// executing the code command in the current working dir to open file in vscode
		commandCode := exec.Command("code", wd)
		err = commandCode.Run()

		// File not found
		if err != nil {
			log.Fatal("VS Code executable file not found in %PATH%")
		}

	},
}

func init() {
	rootCmd.AddCommand(codeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	codeCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "enter the file dir")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// codeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
