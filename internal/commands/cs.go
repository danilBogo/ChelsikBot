package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const csMsg = `
%s идем кс уроды
`

type CsCommand struct {
	bot     *tgbotapi.BotAPI
	pings   string
	command string
}

func NewCsCommand(bot *tgbotapi.BotAPI, pings, command string) *CsCommand {
	return &CsCommand{
		bot:     bot,
		pings:   pings,
		command: command,
	}
}

func (cc *CsCommand) Execute(update tgbotapi.Update) {
	msg := fmt.Sprintf(csMsg, cc.pings)
	tgMsg := tgbotapi.NewMessage(update.Message.Chat.ID, msg)
	_, err := cc.bot.Send(tgMsg)
	if err != nil {
		log.Println(err)
	}
}

func (cc *CsCommand) GetCommandName() string {
	return cc.command
}
