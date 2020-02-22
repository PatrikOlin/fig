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
	"time"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// vatcodeCmd represents the vatcode command
var vatcodeCmd = &cobra.Command{
	Use:   "vatcode",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fstatus, _ := cmd.Flags().GetBool("copy")
		fmt.Println(getVatCode(fstatus))

	},
}

func init() {
	rootCmd.AddCommand(vatcodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vatcodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vatcodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	vatcodeCmd.Flags().BoolP("copy", "c", false, "Copy VAT-code to clipboard")
}

func getVatCode(copyFlag bool) string {

	seed := time.Now().UnixNano()
	prefix := "SE"
	orgNum := getOrgNum(seed)
	suffix := "01"

	parts := []string{prefix, orgNum, suffix}

	vatcode := strings.Join(parts, "")

	if copyFlag == true {
		clipboard.WriteAll(vatcode)
	}

	return vatcode
}
				 
func getVatCodeForOrgNum(orgNum string) string {
	prefix := "SE"
	suffix := "01"

	parts := []string{prefix, orgNum, suffix}

	vatcode := strings.Join(parts, "")

	return vatcode
}
