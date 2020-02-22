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
	"github.com/ShiraazMoollatjie/goluhn"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// orgnumCmd represents the orgnum command
var orgnumCmd = &cobra.Command{
	Use:   "orgnum",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fstatus, _ := cmd.Flags().GetBool("copy")
		timeSeed := time.Now().UnixNano()
		fmt.Println(getFormattedOrgNum(fstatus, timeSeed))
	},
}

func init() {
	rootCmd.AddCommand(orgnumCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// orgnumCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// orgnumCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	orgnumCmd.Flags().BoolP("copy", "c", false, "Copy phone number to clipboard")
}

func getFormattedOrgNum(copyFlag bool, timeSeed int64) string {
	orgNumLuhn := getOrgNum(timeSeed)

	for i := 6; i < len(orgNumLuhn); i += 7 {
		orgNumLuhn = orgNumLuhn[:i] + "-" + orgNumLuhn[i:]
	}

	if copyFlag == true {
		clipboard.WriteAll(orgNumLuhn)
	}

	return orgNumLuhn
}

func getOrgNum(timeSeed int64) string {
	rand.Seed(timeSeed)
	orgNum := orgNumPartial()

	_, orgNumLuhn, err := goluhn.Calculate(orgNum)
	check(err)

	return orgNumLuhn
}

func orgNumPartial() string {
	group := strconv.Itoa(rangeIn(2, 9))
	second := strconv.Itoa(rangeIn(1, 9))
	leading := strconv.Itoa(rangeIn(20, 99))
	numbers := strconv.Itoa(rangeIn(20, 99))
	ending := strconv.Itoa(rangeIn(100, 999))

	stringArray := []string{group, second, leading, numbers, ending}

	return strings.Join(stringArray, "")

}

func toOrgNumString(org []int) string {

	fmt.Println(org)
	string := make([]string, len(org))

	for v := range org {
		string = append(string, strconv.Itoa(v))
	}
	return strings.Join(string, "")
}
