package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"

	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

const IM_DONE_CHATTING = "IMDONE"

func beginChat() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print(string(ColorNarrator), STARLINE+"The chat has begun! You need to start the conversation though... go ahead and introduce yourself. Type IMDONE to finish chatting and return to the main menu.\n")

	userIsDone := false
	chatHistory := openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{}),
		Model:    openai.F(openai.ChatModelGPT4oMini),
	}

	for !userIsDone {
		fmt.Print(string(ColorUser))
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == IM_DONE_CHATTING {
			userIsDone = true
			break
		}

		chatHistory.Messages.Value = append(chatHistory.Messages.Value, openai.UserMessage(userInput))

		chatCompletion, err := client.Chat.Completions.New(context.TODO(), chatHistory)

		if err != nil {
			panic(err.Error())
		}

		chatHistory.Messages.Value = append(chatHistory.Messages.Value, chatCompletion.Choices[0].Message)

		fmt.Print(string(ColorMachine), chatCompletion.Choices[0].Message.Content+"\n")
	}
}
