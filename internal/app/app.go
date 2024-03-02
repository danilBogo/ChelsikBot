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
	Execute(update tgbotapi.Update)
	GetCommandName() string
}

type App struct {
	bot      *tgbotapi.BotAPI
	commands []Command
}

func NewApp() *App {
	log.Println(os.Getwd())
	err := godotenv.Load("../.env")
	if err != nil {
		log.Print("Error loading .env file")
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
					command.Execute(update)
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
	cmds = append(cmds, commands.NewSmokeCommand(bot, "smoke"))
	cmds = append(cmds, commands.NewTitanicCommand(bot, "titanic"))
	cmds = append(cmds, commands.NewFuckYouCommand(bot, "fuckyou"))
	cmds = append(cmds, commands.NewDemqqCommand(bot, "demqq"))
	cmds = append(cmds, commands.NewNiggersGaysCommand(bot, "niggersgays"))
	cmds = append(cmds, commands.NewNiggersNotGaysCommand(bot, "niggersnotgays"))
	cmds = append(cmds, commands.NewSkinCommand(bot, "skin"))
	cmds = append(cmds, commands.NewUpdatesCommand(bot, "updates"))

	return cmds
}
