package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"translate_all/model"
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
	var url string
	url = "http://39.104.17.28:9000/word/?english=" + english
	translationRequest(url)
	// 再发一个请求
	url = "http://39.104.17.28:9000/english?english=" + english
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
	// 反序列化
	// 根据返回状态码来确定查询类型

	if translationResponse.Code > 0 {
		data := translationResponse.Data
		// 先序列化
		jsonData, err := json.Marshal(data)
		if translationResponse.Code == 1 {

			//word
			var word = model.WordTranslation{}
			err = json.Unmarshal(jsonData, &word)
			if err != nil {
				fmt.Println(err)
			}
			//type WordTranslation struct {
			//	ID          uint   `gorm:"primaryKey"`
			//	Word        string `gorm:"column:word"`
			//	Translation string `gorm:"column:translation"`
			//}
			// 打印单词信息
			fmt.Println("result:")
			fmt.Println("Word: ", word.Word)
			fmt.Println("Trans: ", word.Translation)

		} else if translationResponse.Code == 2 {
			//chinese
			var word = model.WordTranslation{}
			err = json.Unmarshal(jsonData, &word)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("related info: ")
			fmt.Println("Trans: ", word.Translation)
			fmt.Println("Word: ", word.Word)
			fmt.Println("extended info about the eng_word: ")
			translateEnglish(word.Word)

		} else if translationResponse.Code == 3 {
			//english
			var word = model.Words{}
			err = json.Unmarshal(jsonData, &word)
			if err != nil {
				fmt.Println(err)
			}
			//type Words struct {
			//	ID              string  `gorm:"column:vc_id"`               // 单词id	57067c89a172044907c6698e
			//	Vocabulary      string  `gorm:"column:vc_vocabulary"`       // 单词	superspecies
			//	PhoneticUK      string  `gorm:"column:vc_phonetic_uk"`      // uk英音音标	[su:pərsˈpi:ʃi:z]
			//	PhoneticUS      string  `gorm:"column:vc_phonetic_us"`      // us美音音标	[supɚsˈpiʃiz]
			//	Frequency       float32 `gorm:"column:vc_frequency"`        // 词频	0.000000
			//	Difficulty      int     `gorm:"column:vc_difficulty"`       // 难度	1
			//	AcknowledgeRate float32 `gorm:"column:vc_acknowledge_rate"` // 认识率 0.664
			//}
			fmt.Println("details:")
			fmt.Println("word: ", word.Vocabulary)
			fmt.Println("US: ", word.PhoneticUS)
			fmt.Println("UK: ", word.PhoneticUK)
			fmt.Println("Difficulty: ", word.Difficulty)
			fmt.Println("Frequency: ", word.Frequency)

		}
	} else {
		fmt.Println("No result:", translationResponse.Message)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
