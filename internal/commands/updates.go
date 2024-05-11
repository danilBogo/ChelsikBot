package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const updatesMsg = `
12.05.2024
New:
- removed obsolete commands
Updated:
- add command /mention
- add command /ton
`

type UpdatesCommand struct {
	bot     *tgbotapi.BotAPI
	command string
}

func NewUpdatesCommand(bot *tgbotapi.BotAPI, command string) *UpdatesCommand {
	return &UpdatesCommand{
		bot:     bot,
		command: command,
	}
}

func (dc *UpdatesCommand) Execute(update tgbotapi.Update) {
	tgMsg := tgbotapi.NewMessage(update.Message.Chat.ID, updatesMsg)
	_, err := dc.bot.Send(tgMsg)
	if err != nil {
		log.Println(err)
	}
}

func (dc *UpdatesCommand) GetCommandName() string {
	return dc.command
}
