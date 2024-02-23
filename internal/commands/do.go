package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const doMsg = "А когда не делали"

type DoCommand struct {
	bot     *tgbotapi.BotAPI
	command string
}

func NewDoCommand(bot *tgbotapi.BotAPI, command string) *DoCommand {
	return &DoCommand{
		bot:     bot,
		command: command,
	}
}

func (dc *DoCommand) Execute(chatId int64) {
	tgMsg := tgbotapi.NewMessage(chatId, doMsg)
	_, err := dc.bot.Send(tgMsg)
	if err != nil {
		log.Println(err)
	}
}

func (dc *DoCommand) GetCommandName() string {
	return dc.command
}
