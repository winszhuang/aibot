package services

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

var (
	client         *openai.Client
	userMessageMap = make(map[string][]openai.ChatCompletionMessage)
)

func Setup() {
	if client == nil {
		aiKey := os.Getenv("OPEN_AI_KEY")
		client = openai.NewClient(aiKey)
	}
}

func Reply(lineId string, message string) string {
	Setup()

	userMessages := loadUserMessage(lineId)
	userMessages = append(userMessages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	})
	setUserMessages(lineId, userMessages)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: userMessages,
		},
	)

	if err != nil {
		return fmt.Sprintf("ChatCompletion error: %v\n", err)
	}

	content := resp.Choices[0].Message.Content
	userMessages = append(userMessages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: content,
	})
	setUserMessages(lineId, userMessages)

	// printMessages(userMessages)

	return content
}

func setUserMessages(lineId string, messages []openai.ChatCompletionMessage) {
	userMessageMap[lineId] = messages
}

func loadUserMessage(lineId string) []openai.ChatCompletionMessage {
	_, ok := userMessageMap[lineId]
	if !ok {
		userMessageMap[lineId] = make([]openai.ChatCompletionMessage, 0)
	}

	return userMessageMap[lineId]
}

// only for test
// func printMessages(msgs []openai.ChatCompletionMessage) {
// 	fmt.Println("------------------")
// 	for i, m := range msgs {
// 		fmt.Println(i)
// 		fmt.Println("名稱: ", m.Name)
// 		fmt.Println("內容: ", m.Content)
// 	}
// }
