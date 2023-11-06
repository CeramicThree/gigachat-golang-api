package http

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/CeramicThree/gigachat-golang-api/config"
	"github.com/CeramicThree/gigachat-golang-api/domain"
	"github.com/joho/godotenv"
)

const (
	endpoint = "/chat/completions"
)

func GigaSend(config *config.ApiConfig) (*domain.GigaResponse, error) {
	log.Printf("Sending gigachat-api request...")

	err := godotenv.Load("../config.env")

	if err != nil {
		log.Fatal("Can't load enviroment variables")
		return nil, err
	}

	marshalled, err := json.Marshal(config)

	if err != nil {
		log.Fatal("Can't marshal api config")
		return nil, err
	}

	client := &http.Client{}

	req, err := http.NewRequest(
		"POST", 
		os.Getenv("baseUrl") + endpoint, 
		bytes.NewReader(marshalled),
	)
	req.Header.Set("Authorization", "Bearer " + os.Getenv("access_token"))

	if err != nil {
		log.Fatal("Can't build api request")
		return nil, err
	}

	response, err := client.Do(req)

	if err != nil {
		log.Fatal("Can't send api request: ", err)
		return nil, err
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	
	if err != nil {
		log.Fatal("Can't read response body")
		return nil, err
	}

	var giga_response *domain.GigaResponse

	error := json.Unmarshal(responseBody, &giga_response) 

	if error != nil {
		log.Fatal(error, response)
		return nil, err
	}

	return giga_response, nil
}