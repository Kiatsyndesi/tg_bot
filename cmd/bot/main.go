package main

import (
	"github.com/Kiatsyndesi/tg_bot/internal/app/router"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Panic(".env file not find")
	}

	token, found := os.LookupEnv("TOKEN")

	if !found {
		log.Panic("TOKEN not found in .env")
	}

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Panic(err)
	}

	//comment if you want to disable debug mode
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	routerHandler := router.NewRouter(bot)

	for update := range updates {
		routerHandler.HandleUpdate(update)
	}
}
