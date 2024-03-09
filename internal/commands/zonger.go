package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const zongerMsg = "Дропни авик плиз"

type ZongerCommand struct {
	bot     *tgbotapi.BotAPI
	command string
}

func NewZongerCommand(bot *tgbotapi.BotAPI, command string) *ZongerCommand {
	return &ZongerCommand{
		bot:     bot,
		command: command,
	}
}

func (dc *ZongerCommand) Execute(update tgbotapi.Update) {
	tgMsg := tgbotapi.NewMessage(update.Message.Chat.ID, zongerMsg)
	_, err := dc.bot.Send(tgMsg)
	if err != nil {
		log.Println(err)
	}
}

func (dc *ZongerCommand) GetCommandName() string {
	return dc.command
}
