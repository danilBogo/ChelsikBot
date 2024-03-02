package app

import (
	"ChelsikBot/internal/commands"
	"ChelsikBot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	totalRequestsCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "bot_requests_total",
		Help: "Total requests number",
	})

	totalCommandCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "bot_requests_total_command",
		Help: "Number of total command requests",
	}, []string{"command"})

	totalUserCommandCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "bot_requests_total_user_command",
		Help: "Number of total user command requests",
	}, []string{"username", "command"})

	successCommandCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "bot_requests_success_command",
		Help: "Number of success command requests",
	}, []string{"command"})

	successUserCommandCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "bot_requests_success_user_command",
		Help: "Number of success user command requests",
	}, []string{"username", "command"})

	requestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "bot_request_duration_seconds",
		Help:    "Histogram of the bot request duration in seconds",
		Buckets: prometheus.DefBuckets,
	}, []string{"command"})
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

	prometheus.MustRegister(totalRequestsCounter)
	prometheus.MustRegister(totalCommandCounter)
	prometheus.MustRegister(totalUserCommandCounter)
	prometheus.MustRegister(successCommandCounter)
	prometheus.MustRegister(successUserCommandCounter)
	prometheus.MustRegister(requestDuration)

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
			totalRequestsCounter.Inc()

			if update.Message == nil || !update.Message.IsCommand() {
				return
			}

			totalCommandCounter.WithLabelValues(update.Message.Command()).Inc()
			totalUserCommandCounter.WithLabelValues(update.Message.From.UserName, update.Message.Command()).Inc()

			for _, command := range a.commands {
				if update.Message.Command() == command.GetCommandName() {
					successCommandCounter.WithLabelValues(update.Message.Command()).Inc()
					successUserCommandCounter.WithLabelValues(update.Message.From.UserName, update.Message.Command()).Inc()
					if update.Message.Chat.IsPrivate() || !a.telegramManager.IsMuted(update, lastMessageTime) {
						command.Execute(update)
						break
					}
				}
			}

			duration := time.Since(start).Seconds()
			requestDuration.WithLabelValues(update.Message.Command()).Observe(duration)
		}(update)
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
	cmds = append(cmds, commands.NewGruntCommand(bot, "grunt"))
	cmds = append(cmds, commands.NewFivePorridgeSpoonfulsCommand(bot, "fiveporridgespoonfuls"))
	cmds = append(cmds, commands.NewCasesCommand(bot, "cases"))

	return cmds
}
