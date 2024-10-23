package commands

import (
	"ChelsikBot/internal/services"
	"bytes"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/uuid"
	"log"
)

const inviteRootDir = "../files/voices/invite"

type InviteCommand struct {
	bot     *tgbotapi.BotAPI
	command string
}

func NewInviteCommand(bot *tgbotapi.BotAPI, command string) *InviteCommand {
	return &InviteCommand{
		bot:     bot,
		command: command,
	}
}

func (ic *InviteCommand) Execute(update tgbotapi.Update) {
	fileBytes := services.GetRandomVoiceBytes(inviteRootDir)
	msg := tgbotapi.NewVoiceUpload(update.Message.Chat.ID, tgbotapi.FileReader{Name: uuid.New().String() + ".ogg", Reader: bytes.NewReader(fileBytes), Size: -1})
	_, err := ic.bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func (ic *InviteCommand) GetCommandName() string {
	return ic.command
}
