package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// GetBot - creating connection with Telegram bot
func GetBot(botToken string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account \"%s\"", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return bot
}
