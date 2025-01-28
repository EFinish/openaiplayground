package main

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

const (
	COMMAND_BEGIN_CHAT          = "begin a new chat session"
	COMMAND_CHAT_NORSK_PRACTICE = "practice norsk"
	COMMAND_EXIT                = "exit"

	STARLINE = "**********\n"

	ColorReset = "\033[0m"
	ColorRed   = "\033[31m"
	ColorGreen = "\033[32m"
	ColorBlue  = "\033[34m"
	ColorPink  = "\033[35m"

	ColorError    = ColorRed
	ColorUser     = ColorBlue
	ColorMachine  = ColorGreen
	ColorNarrator = ColorPink
)

func main() {

	for {
		templates := &promptui.SelectTemplates{
			Active:   templateGenericActive,
			Inactive: templateGenericInactive,
		}

		prompt := promptui.Select{
			Label:     "Select one of the following commands:",
			Items:     []string{COMMAND_BEGIN_CHAT, COMMAND_CHAT_NORSK_PRACTICE, COMMAND_EXIT},
			Templates: templates,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {
		case COMMAND_BEGIN_CHAT:
			beginChat()
		case COMMAND_CHAT_NORSK_PRACTICE:
			beginNorskPractice()
		case COMMAND_EXIT:
			os.Exit(0)
		default:
			fmt.Println(string(ColorReset), "Invalid command!")
		}

	}
}
