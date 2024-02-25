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

func (sc *SmokeCommand) Execute(update tgbotapi.Update) {
	tgMsg := tgbotapi.NewMessage(update.Message.Chat.ID, smokeMsg)
	_, err := sc.bot.Send(tgMsg)
	if err != nil {
		log.Println(err)
	}
}

func (sc *SmokeCommand) GetCommandName() string {
	return sc.command
}
