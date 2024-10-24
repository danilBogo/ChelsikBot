package commands

import (
	"ChelsikBot/internal/services"
	"bytes"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/uuid"
	"log"
)

const fivePorridgeSpoonfulsRootDir = "../files/voices/five_porridge_spoonfuls"

type FivePorridgeSpoonfulsCommand struct {
	bot     *tgbotapi.BotAPI
	command string
}

func NewFivePorridgeSpoonfulsCommand(bot *tgbotapi.BotAPI, command string) *FivePorridgeSpoonfulsCommand {
	return &FivePorridgeSpoonfulsCommand{
		bot:     bot,
		command: command,
	}
}

func (dc *FivePorridgeSpoonfulsCommand) Execute(update tgbotapi.Update) {
	fileBytes := services.GetRandomVoiceBytes(fivePorridgeSpoonfulsRootDir)
	msg := tgbotapi.NewVoiceUpload(update.Message.Chat.ID, tgbotapi.FileReader{Name: uuid.New().String() + ".ogg", Reader: bytes.NewReader(fileBytes), Size: -1})
	_, err := dc.bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func (dc *FivePorridgeSpoonfulsCommand) GetCommandName() string {
	return dc.command
}
