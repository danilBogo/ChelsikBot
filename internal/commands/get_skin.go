package commands

import (
	"ChelsikBot/internal/services"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/uuid"
	"log"
)

const (
	invalidArguments = "@%s еблуша введи полное или частичное название кейса"
	drop             = `
%s
@%s
%s
%s
%s
%s
%s
`
	dropWithPhase = `
%s
@%s
%s
%s
%s
%s
%s
%s
`
)

type SkinCommand struct {
	bot         *tgbotapi.BotAPI
	command     string
	skinManager *services.SkinManager
}

func NewSkinCommand(bot *tgbotapi.BotAPI, command string) *SkinCommand {
	return &SkinCommand{
		bot:         bot,
		command:     command,
		skinManager: services.NewSkinManager(),
	}
}

func (dc *SkinCommand) Execute(update tgbotapi.Update) {
	caseName := update.Message.CommandArguments()
	if len(caseName) == 0 {
		tgMsg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf(invalidArguments, update.Message.From.UserName))
		_, err := dc.bot.Send(tgMsg)
		if err != nil {
			log.Println(err)
		}
		return
	}

	skin, err := dc.skinManager.GetSkin(caseName)
	if err != nil {
		log.Println(err)
		tgMsg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		_, err := dc.bot.Send(tgMsg)
		if err != nil {
			log.Println(err)
		}
		return
	}

	var caption string
	if skin.Phase == nil {
		caption = fmt.Sprintf(drop, skin.Rarity, update.Message.From.UserName, skin.Case, skin.Name, skin.Pattern, skin.Float, skin.Rarity)
	} else {
		caption = fmt.Sprintf(dropWithPhase, skin.Rarity, update.Message.From.UserName, skin.Case, skin.Name, skin.Phase, skin.Pattern, skin.Float, skin.Rarity)
	}

	file := tgbotapi.FileBytes{
		Name:  uuid.New().String() + ".png",
		Bytes: skin.Image,
	}

	photoConfig := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, file)
	photoConfig.Caption = caption

	_, err = dc.bot.Send(photoConfig)
	if err != nil {
		log.Fatal(err)
	}
}

func (dc *SkinCommand) GetCommandName() string {
	return dc.command
}
