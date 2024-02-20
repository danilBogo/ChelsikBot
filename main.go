package main

import (
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

const (
	dailyMsg = `
@cap_carapka @Kovast @repuKouH @dgiud4578 @T_moon808 друзья, пора на чейли! Формат чейли:
1) Что сделал
2) Есть ли блокеры
3) Когда будет готова задача
`
	csMsg = `
@cap_carapka @Kovast @repuKouH @dgiud4578 @T_moon808 идем кс уроды
`
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "daily":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, dailyMsg)
				_, err = bot.Send(msg)
				if err != nil {
					log.Println(err)
				}
			case "cs":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, csMsg)
				_, err = bot.Send(msg)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}
