package app

import (
	"ChelsikBot/internal/commands"
	"ChelsikBot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Command interface {
	Execute(chatId int64)
	GetCommandName() string
}

type App struct {
	bot      *tgbotapi.BotAPI
	commands []Command
}

func NewApp() *App {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	cmds := getCommands(bot, os.Getenv("PINGS"))

	return &App{
		bot:      bot,
		commands: cmds,
	}
}

func (a *App) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := a.bot.GetUpdatesChan(u)

	lastMessageTime := make(map[int]time.Time)

	for update := range updates {
		if update.Message == nil || !update.Message.IsCommand() {
			continue
		}

		for _, command := range a.commands {
			if update.Message.Command() == command.GetCommandName() {
				if !services.IsMuted(a.bot, update, lastMessageTime) {
					command.Execute(update.Message.Chat.ID)
				}
			}
		}
	}
}

func getCommands(bot *tgbotapi.BotAPI, pings string) []Command {
	cmds := make([]Command, 0)

	cmds = append(cmds, commands.NewCsCommand(bot, pings, "cs"))
	cmds = append(cmds, commands.NewDailyCommand(bot, pings, "daily"))
	cmds = append(cmds, commands.NewDoCommand(bot, "do"))

	return cmds
}