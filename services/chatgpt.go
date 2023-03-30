package services

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

var (
	client *openai.Client
)

func Setup() {
	if client == nil {
		fmt.Println("第一次初始化哦")
		aiKey := os.Getenv("OPEN_AI_KEY")
		client = openai.NewClient(aiKey)
	}
}

func Reply(message string) string {
	Setup()
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)

	if err != nil {
		return fmt.Sprintf("ChatCompletion error: %v\n", err)
	}

	return resp.Choices[0].Message.Content
}
