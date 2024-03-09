package commands

import (
	"ChelsikBot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

type CollectionsCommand struct {
	bot         *tgbotapi.BotAPI
	command     string
	skinManager *services.SkinManager
}

func NewCollectionsCommand(bot *tgbotapi.BotAPI, command string) *CollectionsCommand {
	return &CollectionsCommand{
		bot:         bot,
		command:     command,
		skinManager: services.NewSkinManager(),
	}
}

func (cc *CollectionsCommand) Execute(update tgbotapi.Update) {
	strBuilder := strings.Builder{}
	lastValueIndex := len(cc.skinManager.CollectionsName) - 1
	for id, name := range cc.skinManager.CollectionsName {
		strBuilder.WriteString(name)
		if id != lastValueIndex {
			strBuilder.WriteString(" ; ")
		}
	}

	tgMsg := tgbotapi.NewMessage(update.Message.Chat.ID, strBuilder.String())
	_, err := cc.bot.Send(tgMsg)
	if err != nil {
		log.Println(err)
	}
}

func (cc *CollectionsCommand) GetCommandName() string {
	return cc.command
}
