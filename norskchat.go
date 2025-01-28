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

const IM_DONE_PRACTICING_NORSK = "JEGERFERDIG"
const OPENING_PROMPT = "I'd like to practice my Norwegian language skills with you by us speaking norsk with eachother conversationally. During discussion, please feel free to give me corrections to my grammar, word use, and idioms. Limit yourself to three corrections per each of my responses. Kan vi snakke norsk sammen?"

func beginNorskPractice() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print(string(ColorNarrator), STARLINE+"The norsk practice has begun! You need to start the conversation though... go ahead and introduce yourself. Type JEGERFERDIG to finish chatting and return to the main menu.\n")

	userIsDone := false
	chatHistory := openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{}),
		Model:    openai.F(openai.ChatModelGPT4oMini),
	}

	chatHistory.Messages.Value = append(chatHistory.Messages.Value, openai.UserMessage(OPENING_PROMPT))

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), chatHistory)

	if err != nil {
		panic(err.Error())
	}

	chatHistory.Messages.Value = append(chatHistory.Messages.Value, chatCompletion.Choices[0].Message)

	fmt.Print(string(ColorMachine), chatCompletion.Choices[0].Message.Content+"\n")

	for !userIsDone {
		fmt.Print(string(ColorUser))
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == IM_DONE_PRACTICING_NORSK {
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
