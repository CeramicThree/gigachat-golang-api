package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/CeramicThree/gigachat-golang-api/config"
	"github.com/CeramicThree/gigachat-golang-api/domain"
	"github.com/CeramicThree/gigachat-golang-api/http"
)

func main() {
	info_log := log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)

	messages := []domain.Message{
		{
			Role:    "system",
			Content: "Ты ассистен Golang-разработчика",
		},
		{
			Role:    "user",
			Content: "Напиши пример кода Hello World на Golang",
		},
	}

	config := config.Init(&messages)
	res, err := http.GigaSend(config)

	if err != nil {
		info_log.Fatal("Failed send request")
	}

	output, _ := json.Marshal(res.Choices[0].Message.Content)

	info_log.Println(string(output))
}
