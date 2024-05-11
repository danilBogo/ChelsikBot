package commands

import (
	"ChelsikBot/internal/services"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const tonMsg = "TON/USDT: %s"

type TonCommand struct {
	bot            *tgbotapi.BotAPI
	binanceManager *services.BinanceManager
	command        string
}

func NewTonCommand(bot *tgbotapi.BotAPI, binanceManager *services.BinanceManager, command string) *TonCommand {
	return &TonCommand{
		bot:            bot,
		binanceManager: binanceManager,
		command:        command,
	}
}

func (tc *TonCommand) Execute(update tgbotapi.Update) {
	ton := tc.binanceManager.GetTon()
	msg := fmt.Sprintf(tonMsg, ton.Price)
	tgMsg := tgbotapi.NewMessage(update.Message.Chat.ID, msg)
	_, err := tc.bot.Send(tgMsg)
	if err != nil {
		log.Println(err)
	}
}

func (tc *TonCommand) GetCommandName() string {
	return tc.command
}
