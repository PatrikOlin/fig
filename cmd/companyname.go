/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"strings"
	"math/rand"
	"time"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// companynameCmd represents the companyname command
var companynameCmd = &cobra.Command{
	Use:   "companyname",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fstatus, _ := cmd.Flags().GetBool("copy")
		fmt.Println(getCompanyname(fstatus))
	},
}

func init() {
	rootCmd.AddCommand(companynameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// companynameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// companynameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	companynameCmd.Flags().BoolP("copy", "c", false, "Copy company name to clipboard")
}

func getCompanyname(copyFlag bool) string {
	rand.Seed(time.Now().UnixNano())
	numOfWords := rangeIn(2, 5)
	var companyname strings.Builder
	
	for i := 0; i < numOfWords; i++ {
		companyname.WriteString(getRandomLine("companynameparts"))
		companyname.WriteString(" ")
	}

	finishedCompanyname := strings.TrimSpace(companyname.String())
	if copyFlag == true {
		clipboard.WriteAll(finishedCompanyname)
	}

	return finishedCompanyname
}
