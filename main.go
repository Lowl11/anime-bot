package main

import (
	"anime-bot/db"
	"anime-bot/parser"
	"anime-bot/telegram"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jasonlvhit/gocron"
)

// Configuration - structure of json file with settings
type Configuration struct {
	ChatID     int64  `json:"chat_id"`
	BotToken   string `json:"bot_token"`
	DbUser     string `json:"db_user"`
	DbPassword string `json:"db_password"`
	DbName     string `json:"db_name"`
}

// Config - structure of json file with settings
var Config Configuration

func main() {
	readConfig()
	startBotActivity()
}

func startBotActivity() {
	collectAnimeAndSend()                              // first parsing
	gocron.Every(30).Minutes().Do(collectAnimeAndSend) // after 30 minutes

	<-gocron.Start()
}

func collectAnimeAndSend() {
	fmt.Println("")
	log.Printf("*** Job started ***\n\n")

	bot := telegram.GetBot(Config.BotToken)

	// message which will be sended
	var messageText string
	var sendMessage bool

	// getting anime list
	anistarAnimeList := parser.ParseAnistar()

	db.InitializeDatabase(Config.DbUser, Config.DbPassword, Config.DbName)
	comparedAnimeList := db.CompareAnimeList(anistarAnimeList)
	if len(comparedAnimeList) > 0 {
		messageText = telegram.PrepareMessage(comparedAnimeList)
		sendMessage = true
	}

	if sendMessage {
		telegram.SendMessage(bot, Config.ChatID, messageText)
	}

	fmt.Println("")
	log.Printf("*** Job Successfully done ***\n\n")
}

func readConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Panic("Required config.json file. Error message:", err.Error())
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		log.Panic("Couldn't read parse configuration file to object. Error message:", err.Error())
	}
}
