package config

import (
	"github.com/CeramicThree/gigachat-golang-api/domain"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type ApiConfig struct {
	Model              string            `json:"model"`
	Temperature        float64           `json:"temperature"`
	Top_P              float64           `json:"top_p"`
	Answer_num         int               `json:"n"`
	Max_tokens         int               `json:"max_tokens"`
	Repetition_penalty float64           `json:"repetition_penalty"`
	Stream             bool              `json:"stream"`
	Update_interval    int               `json:"update_interval"`
	Messages           *[]domain.Message `json:"messages"`
}

func Init(messages *[]domain.Message) *ApiConfig {
	err := godotenv.Load("../config.env")

	if err != nil {
		log.Fatal("Can't load enviroment variables: ", err)
	}

	apiConfig := new(ApiConfig)

	apiConfig.Model = os.Getenv("model")
	apiConfig.Temperature, _ = strconv.ParseFloat(os.Getenv("temperature"), 64)
	apiConfig.Top_P, _ = strconv.ParseFloat(os.Getenv("top_p"), 64)
	apiConfig.Answer_num, _ = strconv.Atoi(os.Getenv("n"))
	apiConfig.Max_tokens, _ = strconv.Atoi(os.Getenv("max_tokens"))
	apiConfig.Repetition_penalty, _ = strconv.ParseFloat(os.Getenv("repetition_penalty"), 64)
	apiConfig.Stream, _ = strconv.ParseBool(os.Getenv("stream"))
	apiConfig.Update_interval, _ = strconv.Atoi(os.Getenv("update_interval"))
	apiConfig.Messages = messages

	return apiConfig
}
