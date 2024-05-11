package commands

import (
	"ChelsikBot/internal/services"
	"bytes"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/uuid"
	"log"
)

const mentionRootDir = "../files/voices/mention"

type MentionCommand struct {
	bot     *tgbotapi.BotAPI
	pings   string
	command string
}

func NewMentionCommand(bot *tgbotapi.BotAPI, pings, command string) *MentionCommand {
	return &MentionCommand{
		bot:     bot,
		pings:   pings,
		command: command,
	}
}

func (mc *MentionCommand) Execute(update tgbotapi.Update) {
	fileBytes := services.GetRandomVoiceBytes(mentionRootDir)
	tgMsg := tgbotapi.NewMessage(update.Message.Chat.ID, mc.pings)
	_, err := mc.bot.Send(tgMsg)
	if err != nil {
		log.Println(err)
	}

	msg := tgbotapi.NewVoiceUpload(update.Message.Chat.ID, tgbotapi.FileReader{Name: uuid.New().String() + ".ogg", Reader: bytes.NewReader(fileBytes), Size: -1})
	_, err = mc.bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func (mc *MentionCommand) GetCommandName() string {
	return mc.command
}
