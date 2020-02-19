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

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// emailCmd represents the email command
var emailCmd = &cobra.Command{
	Use:   "email",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fstatus, _ := cmd.Flags().GetBool("copy")
		fmt.Println(getEmail(fstatus))
	},
}

func init() {
	rootCmd.AddCommand(emailCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// emailCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// emailCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	emailCmd.Flags().BoolP("copy", "c", false, "Copy person to clipboard")
}

func getEmail(copyFlag bool) string {
	var emailAddress strings.Builder
	emailAddress.WriteString(getFirstName())
	emailAddress.WriteString(".")
	emailAddress.WriteString(getSurname())
	emailAddress.WriteString(getRandomLine("emaildomains"))

	if copyFlag == true {
		clipboard.WriteAll(strings.ToLower(emailAddress.String()))
	}

	return strings.ToLower(emailAddress.String())
}

func getEmailForName(fullname string) string {
	names := strings.Fields(fullname)
	var emailAddress strings.Builder
	emailAddress.WriteString(names[0])
	emailAddress.WriteString(".")
	emailAddress.WriteString(names[1])
	emailAddress.WriteString(getRandomLine("emaildomains"))

	return strings.ToLower(emailAddress.String())
}
