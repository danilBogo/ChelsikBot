package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const niggersNotGays = `
Нееегры - не пидоры ни разу.
Я понял это сразу, я был не прав.
Нееегры - обычные ребята.
И трогать их не надо
`

type NiggersNotGaysCommand struct {
	bot     *tgbotapi.BotAPI
	command string
}

func NewNiggersNotGaysCommand(bot *tgbotapi.BotAPI, command string) *NiggersNotGaysCommand {
	return &NiggersNotGaysCommand{
		bot:     bot,
		command: command,
	}
}

func (dc *NiggersNotGaysCommand) Execute(update tgbotapi.Update) {
	tgMsg := tgbotapi.NewMessage(update.Message.Chat.ID, niggersNotGays)
	_, err := dc.bot.Send(tgMsg)
	if err != nil {
		log.Println(err)
	}
}

func (dc *NiggersNotGaysCommand) GetCommandName() string {
	return dc.command
}
