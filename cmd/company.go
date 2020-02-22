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
	"github.com/atotto/clipboard"
	"time"			   
	"encoding/json"

	"github.com/spf13/cobra"
)

type Company struct {
	CompanyName     string
	OrgNum          string
	VatCode         string
	BeneficialOwner string
}

// companyCmd represents the company command
var companyCmd = &cobra.Command{
	Use:   "company",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fstatus, _ := cmd.Flags().GetBool("copy")
		printCompany(fstatus)
	},
}

func init() {
	rootCmd.AddCommand(companyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// companyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// companyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	companyCmd.Flags().BoolP("copy", "c", false, "Copy company to clipboard")
}

func printCompany(copyFlag bool) {
	company := getCompany(copyFlag)
	fmt.Printf("Företagsnamn: %s \n", company.CompanyName)
	fmt.Printf("Org.Nr: %s \n", company.OrgNum)
	fmt.Printf("VAT-nr: %s \n", company.VatCode)
	fmt.Printf("Verklig Huvudman: %s \n", company.BeneficialOwner)
} 					   

func getCompany(copyFlag bool) Company {
	// var company strings.Builder
	seed := time.Now().UnixNano()
	// orgNum := getOrgNum()
	// var formattedOrgNum string
	// for i := 6; i < len(orgNum); i += 7 {
	// 	formattedOrgNum = orgNum[:i] + "-" + orgNum[i:]
	// }
	// company.WriteString(getCompanyname(false))
	// company.WriteString("\n")
	// company.WriteString(formattedOrgNum)
	// company.WriteString("\n")
	// company.WriteString(getVatCodeForOrgNum(orgNum))
	// company.WriteString("\n")
	// company.WriteString(getFullName(false))
	company := Company{
		CompanyName: getCompanyname(false),
		OrgNum: getFormattedOrgNum(false, seed), 
		VatCode: getVatCodeForOrgNum(getOrgNum(seed)),
		BeneficialOwner: getFullName(false),
	}				   

	if copyFlag == true {
		json, err := json.Marshal(company)
		check(err)
		clipboard.WriteAll(string(json))
	}

	return company
}					   
