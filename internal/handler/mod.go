package handler

import (
	"context"
	"fmt"

	"git_commit_assistant/internal/model"

	openrouter "github.com/revrost/go-openrouter"
)

func LLM_message(data string, credentials model.CredentialsFile) (string, error) {
	fmt.Println(credentials)

	client := openrouter.NewClient(
		credentials.Key,
		openrouter.WithXTitle("Git Commit Assistant"),
		openrouter.WithHTTPReferer("https://gitcommitassistant.com"),
	)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openrouter.ChatCompletionRequest{
			Model: credentials.Model,
			Messages: []openrouter.ChatCompletionMessage{
				openrouter.UserMessage(data),
			},
		},
	)

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("Error : %s ", err.Error())
	}

	response := resp.Choices[0].Message.Content

	return response.Text, err
}
