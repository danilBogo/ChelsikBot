package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const updatesMsg = `
03.03.2024
New:
- added metrics bot_requests_total, bot_requests_total_command, bot_requests_total_user_command, bot_requests_success_command, bot_requests_success_user_command, bot_request_duration_seconds
- added prometheus
- added graphana

05.03.2024
New:
- added pattern, float

10.03.2024
New:
- added commands /faceit /zonger /collections
Updated:
- changed mute behaviour
- added collections to command /skin
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
