package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tsl",
	Short: "Translate words",
	Long:  "Translate words between English and Chinese",
	Run: func(cmd *cobra.Command, args []string) {
		english, _ := cmd.Flags().GetString("english")
		chinese, _ := cmd.Flags().GetString("chinese")

		if english != "" {
			translateEnglish(english)
		} else if chinese != "" {
			translateChinese(chinese)
		} else {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&english, "english", "e", "", "Translate English word")
	rootCmd.Flags().StringVarP(&chinese, "chinese", "c", "", "Translate Chinese word")
}

var english string
var chinese string

type TranslationResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func translateEnglish(english string) {
	url := "http://39.104.17.28:9000/word/?english=" + english
	translationRequest(url)
}

func translateChinese(chinese string) {
	url := "http://39.104.17.28:9000/chinese/?chinese=" + chinese
	translationRequest(url)
}

func translationRequest(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	var translationResponse TranslationResponse
	err = json.NewDecoder(response.Body).Decode(&translationResponse)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if translationResponse.Code == 1 {
		fmt.Println("Translation:", translationResponse.Data)
	} else {
		fmt.Println("Translation failed:", translationResponse.Message)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
