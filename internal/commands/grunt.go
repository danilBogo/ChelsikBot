package commands

import (
	"ChelsikBot/internal/services"
	"bytes"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const gruntRootDir = "../files/voices/grunt"

type GruntCommand struct {
	bot     *tgbotapi.BotAPI
	command string
}

func NewGruntCommand(bot *tgbotapi.BotAPI, command string) *GruntCommand {
	return &GruntCommand{
		bot:     bot,
		command: command,
	}
}

func (dc *GruntCommand) Execute(update tgbotapi.Update) {
	fileBytes := services.GetRandomVoiceBytes(gruntRootDir)
	msg := tgbotapi.NewVoiceUpload(update.Message.Chat.ID, tgbotapi.FileReader{Name: "voice_message.ogg", Reader: bytes.NewReader(fileBytes), Size: -1})
	_, err := dc.bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func (dc *GruntCommand) GetCommandName() string {
	return dc.command
}
