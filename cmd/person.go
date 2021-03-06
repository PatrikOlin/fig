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
	"github.com/spf13/cobra"
	"encoding/json"
	"fig/api"
	"fig/models"
)


// personCmd represents the person command
var personCmd = &cobra.Command{
	Use:   "person",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		copyFlag, _ := cmd.Flags().GetBool("copy")
		multiFlag, _ := cmd.Flags().GetInt("multi")
		getAndPrintPerson(copyFlag, multiFlag)
	},
}

func init() {
	rootCmd.AddCommand(personCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// personCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// personCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	personCmd.Flags().BoolP("copy", "c", false, "Copy person to clipboard")
	personCmd.Flags().IntP("multi", "m", 1, "Generate as many friends as you damn well please")
}


func getAndPrintPerson(copyFlag bool, amount int) {
	client := api.NewBasicClient("http://localhost:8080/api/v1")

	people, err := client.GetPeople(amount)
	check(err)

	if copyFlag == true {
		json, err := json.Marshal(people)
		check(err)
		clipboard.WriteAll(string(json))
	}

		printPeople(people)
}

func printPeople(people []models.Person) {
	for _, person := range people {
		printPerson(person)
	}
}

func printPerson(person models.Person) {
	fmt.Printf("Namn: %s \n", person.Name)
	fmt.Printf("Personnummer: %s \n", person.Pin)
	fmt.Printf("Adress: %s \n", person.Address)
	fmt.Printf("Telefonnummer: %s \n", person.Phone)
	fmt.Printf("Epost: %s \n", person.Email)
	fmt.Printf("Lösenord: %s \n", person.Password)
	fmt.Println("-----------------------------------------")
}

func getPerson(copyFlag bool) models.Person {
	fullname := getFullName(false)
	email := getEmailForName(fullname)
	person := models.Person{
		Name: fullname,
		Pin: getPIN(false),
		Address: getFullAddress(false),
		Phone: getPhoneNumber(false),
		Email: email,
		Password: getPassword(false),
	}

	if copyFlag == true {
		json, err := json.Marshal(person)
		check(err)
		clipboard.WriteAll(string(json))
	}

	return person
}
