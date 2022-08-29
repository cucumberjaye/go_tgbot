package main

import (
	"github.com/cucumberjaye/go_tgbot/pkg/bot"
	"github.com/cucumberjaye/go_tgbot/pkg/database"
	"time"
)

func main() {
	if err := database.CreateTable(); err != nil {
		panic(err)
	}
	time.Sleep(1 * time.Minute)

	bot.TelegramBot()
}
