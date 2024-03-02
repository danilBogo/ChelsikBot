package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const skin = "⊂=====3 соси, пока что не сделал"

type SkinCommand struct {
	bot     *tgbotapi.BotAPI
	command string
}

func NewSkinCommand(bot *tgbotapi.BotAPI, command string) *SkinCommand {
	return &SkinCommand{
		bot:     bot,
		command: command,
	}
}

func (dc *SkinCommand) Execute(update tgbotapi.Update) {
	tgMsg := tgbotapi.NewMessage(update.Message.Chat.ID, skin)
	_, err := dc.bot.Send(tgMsg)
	if err != nil {
		log.Println(err)
	}
}

func (dc *SkinCommand) GetCommandName() string {
	return dc.command
}
