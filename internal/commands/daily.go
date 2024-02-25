package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const dailyMsg = `
%s друзья, пора на чейли! Формат чейли:
1) Что сделал
2) Есть ли блокеры
3) Когда будет готова задача
`

type DailyCommand struct {
	bot     *tgbotapi.BotAPI
	pings   string
	command string
}

func NewDailyCommand(bot *tgbotapi.BotAPI, pings, command string) *DailyCommand {
	return &DailyCommand{
		bot:     bot,
		pings:   pings,
		command: command,
	}
}

func (dc *DailyCommand) Execute(update tgbotapi.Update) {
	msg := fmt.Sprintf(dailyMsg, dc.pings)
	tgMsg := tgbotapi.NewMessage(update.Message.Chat.ID, msg)
	_, err := dc.bot.Send(tgMsg)
	if err != nil {
		log.Println(err)
	}
}

func (dc *DailyCommand) GetCommandName() string {
	return dc.command
}
