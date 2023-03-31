package services

import (
	"os"

	"github.com/sashabaranov/go-openai"
)

var (
	client *openai.Client
)

func InitOpenAiClient() *openai.Client {
	if client == nil {
		aiKey := os.Getenv("OPEN_AI_KEY")
		client = openai.NewClient(aiKey)
	}
	return client
}
