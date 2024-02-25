package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const titanicMsg = "Нет, не надо слов, не надо паники. Это мой последний дееень на Титанике"

type TitanicCommand struct {
	bot     *tgbotapi.BotAPI
	command string
}

func NewTitanicCommand(bot *tgbotapi.BotAPI, command string) *TitanicCommand {
	return &TitanicCommand{
		bot:     bot,
		command: command,
	}
}

func (tc *TitanicCommand) Execute(chatId int64) {
	tgMsg := tgbotapi.NewMessage(chatId, titanicMsg)
	_, err := tc.bot.Send(tgMsg)
	if err != nil {
		log.Println(err)
	}
}

func (tc *TitanicCommand) GetCommandName() string {
	return tc.command
}
