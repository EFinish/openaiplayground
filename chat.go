package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func beginChat() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print(string(ColorNarrator), STARLINE+"The chat has begun! You need to start the conversation though... go ahead and introduce yourself.:\n")
	fmt.Print(string(ColorUser))
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(input),
		}),
		Model: openai.F(openai.ChatModelGPT4o),
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Print(string(ColorMachine), chatCompletion.Choices[0].Message.Content+"\n")
}
