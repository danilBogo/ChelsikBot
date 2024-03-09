package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const faceitMsg = "Я фейсит"

type FaceitCommand struct {
	bot     *tgbotapi.BotAPI
	command string
}

func NewFaceitCommand(bot *tgbotapi.BotAPI, command string) *FaceitCommand {
	return &FaceitCommand{
		bot:     bot,
		command: command,
	}
}

func (dc *FaceitCommand) Execute(update tgbotapi.Update) {
	tgMsg := tgbotapi.NewMessage(update.Message.Chat.ID, faceitMsg)
	_, err := dc.bot.Send(tgMsg)
	if err != nil {
		log.Println(err)
	}
}

func (dc *FaceitCommand) GetCommandName() string {
	return dc.command
}
