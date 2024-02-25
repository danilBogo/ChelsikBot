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

func (dc *DoCommand) Execute(update tgbotapi.Update) {
	tgMsg := tgbotapi.NewMessage(update.Message.Chat.ID, doMsg)
	_, err := dc.bot.Send(tgMsg)
	if err != nil {
		log.Println(err)
	}
}

func (dc *DoCommand) GetCommandName() string {
	return dc.command
}
