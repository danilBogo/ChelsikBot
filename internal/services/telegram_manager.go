package services

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"time"
)

const muteDuration = 5

type TelegramManager struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramManager(bot *tgbotapi.BotAPI) *TelegramManager {
	return &TelegramManager{bot: bot}
}

type MuteInfo struct {
	lastMsgTime    time.Time
	warningMsgSent bool
}

func (tm *TelegramManager) IsMuted(update tgbotapi.Update, lastMessageTime map[int]*MuteInfo) bool {
	userID := update.Message.From.ID
	muteInfo, ok := lastMessageTime[userID]
	if ok && time.Now().Sub(muteInfo.lastMsgTime) < time.Second*5 {
		muteInfo.lastMsgTime = time.Now()

		deleteConfig := tgbotapi.DeleteMessageConfig{
			ChatID:    update.Message.Chat.ID,
			MessageID: update.Message.MessageID,
		}

		_, err := tm.bot.DeleteMessage(deleteConfig)
		if err != nil {
			log.Println(err)

			msg := fmt.Sprintf("@%s заебал флудить урод", update.Message.From.UserName)
			tgMsg := tgbotapi.NewMessage(update.Message.Chat.ID, msg)
			_, err := tm.bot.Send(tgMsg)
			if err != nil {
				log.Println(err)
			}

			return true
		}

		if !muteInfo.warningMsgSent {
			msg := fmt.Sprintf("@%s отъехал в мут на %d секунд", update.Message.From.UserName, muteDuration)
			tgMsg := tgbotapi.NewMessage(update.Message.Chat.ID, msg)
			_, err = tm.bot.Send(tgMsg)
			if err != nil {
				log.Println(err)
			}

			muteInfo.warningMsgSent = true
		}

		return true
	} else {
		if !ok {
			lastMessageTime[userID] = &MuteInfo{lastMsgTime: time.Now(), warningMsgSent: false}
		} else {
			muteInfo.lastMsgTime = time.Now()
			muteInfo.warningMsgSent = false
		}
		return false
	}
}
