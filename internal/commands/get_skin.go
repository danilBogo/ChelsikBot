package commands

import (
	"ChelsikBot/internal/services"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const (
	invalidArguments = "%s еблуша введи полное или частичное название кейсы"
	drop             = `
@%s
%s
Уровень редкости: %s
`
	dropWithPhase = `
@%s
%s
Уровень редкости: %s
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
		tgMsg := tgbotapi.NewMessage(update.Message.Chat.ID, invalidArguments)
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
		caption = fmt.Sprintf(drop, update.Message.From.UserName, skin.Name, skin.Rarity)
	} else {
		caption = fmt.Sprintf(dropWithPhase, update.Message.From.UserName, skin.Name, skin.Rarity, skin.Phase)
	}

	photoConfig := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, skin.Image)
	photoConfig.Caption = caption

	_, err = dc.bot.Send(photoConfig)
	if err != nil {
		log.Fatal(err)
	}
}

func (dc *SkinCommand) GetCommandName() string {
	return dc.command
}
