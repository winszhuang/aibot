package services

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

func DallEReply(prompt string) (string, error) {
	c := InitOpenAiClient()
	ctx := context.Background()

	// Sample image by link
	reqUrl := openai.ImageRequest{
		Prompt:         prompt,
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatURL,
		N:              1,
	}

	respUrl, err := c.CreateImage(ctx, reqUrl)
	if err != nil {
		return "", err
	}
	// only for test
	// fmt.Println(respUrl.Data[0].URL)

	return respUrl.Data[0].URL, nil
}
