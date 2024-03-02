package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const niggersGays = "Негры пидарасы"

type NiggersGaysCommand struct {
	bot     *tgbotapi.BotAPI
	command string
}

func NewNiggersGaysCommand(bot *tgbotapi.BotAPI, command string) *NiggersGaysCommand {
	return &NiggersGaysCommand{
		bot:     bot,
		command: command,
	}
}

func (dc *NiggersGaysCommand) Execute(update tgbotapi.Update) {
	tgMsg := tgbotapi.NewMessage(update.Message.Chat.ID, niggersGays)
	_, err := dc.bot.Send(tgMsg)
	if err != nil {
		log.Println(err)
	}
}

func (dc *NiggersGaysCommand) GetCommandName() string {
	return dc.command
}
