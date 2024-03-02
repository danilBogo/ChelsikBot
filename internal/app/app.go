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
	bot             *tgbotapi.BotAPI
	commands        []Command
	telegramManager *services.TelegramManager
	voiceManager    *services.VoiceManager
}

func NewApp() *App {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(wd)

	err = godotenv.Load("../.env")
	if err != nil {
		log.Print("Error loading .env file")
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if len(token) == 0 {
		log.Fatal("TELEGRAM_BOT_TOKEN not found")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	pings := os.Getenv("PINGS")
	if len(pings) == 0 {
		log.Fatal("PINGS not found")
	}

	cmds := getCommands(bot, pings)

	telegramManager := services.NewTelegramManager(bot)

	voiceManager := services.NewVoiceManager()

	return &App{
		bot:             bot,
		commands:        cmds,
		telegramManager: telegramManager,
		voiceManager:    voiceManager,
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
				if !a.telegramManager.IsMuted(update, lastMessageTime) {
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
	//cmds = append(cmds, commands.NewSkinCommand(bot, "skin"))
	cmds = append(cmds, commands.NewUpdatesCommand(bot, "updates"))
	cmds = append(cmds, commands.NewGruntCommand(bot, "grunt"))
	cmds = append(cmds, commands.NewFivePorridgeSpoonfulsCommand(bot, "fiveporridgespoonfuls"))

	return cmds
}
