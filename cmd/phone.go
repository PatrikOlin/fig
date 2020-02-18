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
	"math/rand"
	"strconv"
	"time"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// phoneCmd represents the phone command
var phoneCmd = &cobra.Command{
	Use:   "phone",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fstatus, _ := cmd.Flags().GetBool("copy")
		fmt.Println(getPhoneNumber(fstatus))
	},
}

func init() {
	rootCmd.AddCommand(phoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// phoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// phoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	phoneCmd.Flags().BoolP("copy", "c", false, "Copy phone number to clipboard")
}

func getPhoneNumber(copyFlag bool) string {
	rand.Seed(time.Now().UnixNano())
	numbers := rangeIn(10000000, 99999999)
	initial := rangeIn(0, 9)
	phoneNum := "07" + strconv.Itoa(initial) + "-" + strconv.Itoa(numbers)

	if copyFlag == true {
		clipboard.WriteAll(phoneNum)
	}
	return phoneNum
}

func rangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}
