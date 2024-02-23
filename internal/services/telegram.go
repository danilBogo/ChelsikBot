package services

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"time"
)

func IsMuted(bot *tgbotapi.BotAPI, update tgbotapi.Update, lastMessageTime map[int]time.Time) bool {
	userID := update.Message.From.ID
	lastMsgTime, ok := lastMessageTime[userID]
	if ok && time.Now().Sub(lastMsgTime) < time.Second*5 {
		lastMessageTime[userID] = time.Now()
		msg := fmt.Sprintf("@%s заебал флудить урод", update.Message.From.UserName)
		tgMsg := tgbotapi.NewMessage(update.Message.Chat.ID, msg)
		_, err := bot.Send(tgMsg)
		if err != nil {
			log.Println(err)
		}
		return true
	} else {
		lastMessageTime[userID] = time.Now()
		return false
	}
}
