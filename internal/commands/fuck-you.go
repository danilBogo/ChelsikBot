package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const fuckYouMsg = "@%s пошел нахуй"
const fuckYouBotMsg = "@%s ахуел чурка блять? Сам пошел нахуй"

type FuckYouCommand struct {
	bot     *tgbotapi.BotAPI
	command string
}

func NewFuckYouCommand(bot *tgbotapi.BotAPI, command string) *FuckYouCommand {
	return &FuckYouCommand{
		bot:     bot,
		command: command,
	}
}

func (fyc *FuckYouCommand) Execute(update tgbotapi.Update) {
	var msg string
	if update.Message.ReplyToMessage != nil {
		if update.Message.ReplyToMessage.From.IsBot {
			msg = fmt.Sprintf(fuckYouBotMsg, update.Message.From.UserName)
		} else {
			msg = fmt.Sprintf(fuckYouMsg, update.Message.ReplyToMessage.From.UserName)
		}
	} else {
		msg = fmt.Sprintf(fuckYouMsg, update.Message.From.UserName)
	}

	tgMsg := tgbotapi.NewMessage(update.Message.Chat.ID, msg)
	_, err := fyc.bot.Send(tgMsg)
	if err != nil {
		log.Println(err)
	}
}

func (fyc *FuckYouCommand) GetCommandName() string {
	return fyc.command
}
