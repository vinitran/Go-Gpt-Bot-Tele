package gpt

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"os"
)

const envGptToken = "GPT_TOKEN"

func Predict(quest string) (string, error) {
	c := openai.NewClient(os.Getenv(envGptToken))

	resp, err := c.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "user",
					Content: quest,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
