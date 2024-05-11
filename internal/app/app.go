package app

import (
	"ChelsikBot/internal/app/metrics"
	"ChelsikBot/internal/commands"
	"ChelsikBot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
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
	loadEnv()

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if len(token) == 0 {
		log.Fatal("TELEGRAM_BOT_TOKEN not found")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	cmds := getCommands(bot)
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

	lastMessageTime := make(map[int]*services.MuteInfo)

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	for update := range updates {
		go func(update tgbotapi.Update) {
			start := time.Now()
			metrics.TotalRequestsCounter.Inc()

			if update.Message == nil || !update.Message.IsCommand() {
				return
			}

			metrics.TotalCommandCounter.WithLabelValues(update.Message.Command()).Inc()
			metrics.TotalUserCommandCounter.WithLabelValues(update.Message.From.UserName, update.Message.Command()).Inc()

			for _, command := range a.commands {
				if update.Message.Command() == command.GetCommandName() {
					metrics.SuccessCommandCounter.WithLabelValues(update.Message.Command()).Inc()
					metrics.SuccessUserCommandCounter.WithLabelValues(update.Message.From.UserName, update.Message.Command()).Inc()
					if update.Message.Chat.IsPrivate() || !a.telegramManager.IsMuted(update, lastMessageTime) {
						command.Execute(update)
						break
					}
				}
			}

			duration := time.Since(start).Seconds()
			metrics.RequestDuration.WithLabelValues(update.Message.Command()).Observe(duration)
		}(update)
	}
}

func (a *App) Health() bool {
	_, err := a.bot.GetMe()
	return err == nil
}

func loadEnv() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(wd)

	err = godotenv.Load("../.env")
	if err != nil {
		log.Print("Error loading .env file")
	}
}

func getCommands(bot *tgbotapi.BotAPI) []Command {
	pings := os.Getenv("PINGS")
	if len(pings) == 0 {
		log.Fatal("PINGS not found")
	}

	cmds := make([]Command, 0)

	cmds = append(cmds, commands.NewCsCommand(bot, pings, "cs"))
	cmds = append(cmds, commands.NewDailyCommand(bot, pings, "daily"))
	cmds = append(cmds, commands.NewDoCommand(bot, "do"))
	cmds = append(cmds, commands.NewSmokeCommand(bot, "smoke"))
	cmds = append(cmds, commands.NewTitanicCommand(bot, "titanic"))
	cmds = append(cmds, commands.NewFuckYouCommand(bot, "fuckyou"))
	cmds = append(cmds, commands.NewUpdatesCommand(bot, "updates"))
	cmds = append(cmds, commands.NewGruntCommand(bot, "grunt"))
	cmds = append(cmds, commands.NewFivePorridgeSpoonfulsCommand(bot, "fiveporridgespoonfuls"))

	return cmds
}
