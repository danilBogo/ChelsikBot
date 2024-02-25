package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const smokeMsg = "Дымооок. Пошел по комнате дымок"

type SmokeCommand struct {
	bot     *tgbotapi.BotAPI
	command string
}

func NewSmokeCommand(bot *tgbotapi.BotAPI, command string) *SmokeCommand {
	return &SmokeCommand{
		bot:     bot,
		command: command,
	}
}

func (sc *SmokeCommand) Execute(chatId int64) {
	tgMsg := tgbotapi.NewMessage(chatId, smokeMsg)
	_, err := sc.bot.Send(tgMsg)
	if err != nil {
		log.Println(err)
	}
}

func (sc *SmokeCommand) GetCommandName() string {
	return sc.command
}
