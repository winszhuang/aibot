package services

import (
	"context"
	"fmt"
	"sync"

	"github.com/sashabaranov/go-openai"
)

var (
	userMessageMap = sync.Map{}
)

func GPTReply(lineId string, message string) string {
	client := InitOpenAiClient()

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
	userMessageMap.Store(lineId, messages)
}

func loadUserMessage(lineId string) []openai.ChatCompletionMessage {
	v, ok := userMessageMap.Load(lineId)
	if !ok {
		v = make([]openai.ChatCompletionMessage, 0)
		userMessageMap.Store(lineId, v)
	}

	return v.([]openai.ChatCompletionMessage)
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
