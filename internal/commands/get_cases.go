package commands

import (
	"ChelsikBot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

type CasesCommand struct {
	bot         *tgbotapi.BotAPI
	command     string
	skinManager *services.SkinManager
}

func NewCasesCommand(bot *tgbotapi.BotAPI, command string) *CasesCommand {
	return &CasesCommand{
		bot:         bot,
		command:     command,
		skinManager: services.NewSkinManager(),
	}
}

func (dc *CasesCommand) Execute(update tgbotapi.Update) {
	cases := dc.skinManager.GetCases()
	strBuilder := strings.Builder{}
	lastValueIndex := len(cases) - 1
	for id, c := range cases {
		strBuilder.WriteString(c.Name)
		if id != lastValueIndex {
			strBuilder.WriteString(" ; ")
		}
	}

	tgMsg := tgbotapi.NewMessage(update.Message.Chat.ID, strBuilder.String())
	_, err := dc.bot.Send(tgMsg)
	if err != nil {
		log.Println(err)
	}
}

func (dc *CasesCommand) GetCommandName() string {
	return dc.command
}
