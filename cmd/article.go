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
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"fig/api"								
	"fig/models"						

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)


// articleCmd represents the article command
var articleCmd = &cobra.Command{
	Use:   "article",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		copyFlag, _ := cmd.Flags().GetBool("copy")
		multiFlag, _ := cmd.Flags().GetInt("multi")
		getAndPrintArticle(copyFlag, multiFlag)
	},
}

func init() {
	rootCmd.AddCommand(articleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// articleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// articleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	articleCmd.Flags().BoolP("copy", "c", false, "Copy article to clipboard")
	articleCmd.Flags().IntP("multi", "m", 1, "Generate multiple articles")
}

func getAndPrintArticle(copyFlag bool, amount int) {
	client := api.NewBasicClient("http://localhost:8080/api/v1")

	articles, _ := client.GetArticles(amount)

		if copyFlag == true {
			json, err := json.Marshal(articles)
			check(err)
			clipboard.WriteAll(string(json))
		}

		printArticles(articles)
	}
										
func printArticles(articles []models.Article) {
	for _, article := range articles {
		printArticle(article)
	}
}

func printArticle(article models.Article) {
	fmt.Printf("Artikelid: %s \n", article.ID)
	fmt.Printf("Artikel: %s \n", article.Description)
	fmt.Printf("Momskod: %d \n", article.VatCode)
	fmt.Printf("Pris: %s:- %s \n", article.Price, article.Unit)
	fmt.Println("-----------------------------------------")
}

func getArticle(copyFlag bool) models.Article {
	rand.Seed(time.Now().UnixNano())

	article := models.Article{
		ID:          strconv.Itoa(rangeIn(1, 9999)),
		Description: getArticleName(),
		VatCode:     rangeIn(0, 3),
		Price:       strconv.Itoa(rangeIn(1, 99999)),
		Unit:        "st",
	}

	if copyFlag == true {
		json, err := json.Marshal(article)
		check(err)
		clipboard.WriteAll(string(json))
	}
	return article
}

func getArticleName() string {
	article := getRandomLine("articles")
	return article						
}
