package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const updatesMsg = `
02.03.2024
New:
- added commands /demqq, /niggersgays, /niggersnotgays, /skin, /updates, /grunt, /fiveporridgespoonfuls, /cases
Updated:
- changed response when sending /fuckyou to reply to bot message
- added automatic push to docker hub via github workflow
- remove mute in private conversation
03.03.2024
New:
- added metrics bot_requests_total, bot_requests_total_command, bot_requests_total_user_command, bot_requests_success_command, bot_requests_success_user_command, bot_request_duration_seconds
- added prometheus
- added graphana
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
