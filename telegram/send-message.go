package telegram

import (
	"anime-bot/parser"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// SendMessage - sending message to Telegram Chat
func SendMessage(bot *tgbotapi.BotAPI, chatID int64, message string) {
	messageConfig := tgbotapi.NewMessage(chatID, strings.TrimSpace(message))
	bot.Send(messageConfig)
}

// PrepareMessage - prepare message to send message to Telegram Chat
func PrepareMessage(animeList []parser.Anime) string {
	var message string
	if len(animeList) == 1 {
		message += "Вышло новое аниме!"
	} else if len(animeList) >= 2 {
		message += "Вышли новые аниме!"
	}
	message += "\n"

	for _, anime := range animeList {
		message += fmt.Sprintf("%s\nСерия: %s\nСсылка: %s\n\n", anime.Title, anime.Episode, anime.Url)
	}
	return message
}
