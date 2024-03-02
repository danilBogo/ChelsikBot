package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const demqq = "Дем хрю хрю"

type DemqqCommand struct {
	bot     *tgbotapi.BotAPI
	command string
}

func NewDemqqCommand(bot *tgbotapi.BotAPI, command string) *DemqqCommand {
	return &DemqqCommand{
		bot:     bot,
		command: command,
	}
}

func (dc *DemqqCommand) Execute(update tgbotapi.Update) {
	tgMsg := tgbotapi.NewMessage(update.Message.Chat.ID, demqq)
	_, err := dc.bot.Send(tgMsg)
	if err != nil {
		log.Println(err)
	}
}

func (dc *DemqqCommand) GetCommandName() string {
	return dc.command
}
