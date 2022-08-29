package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/cucumberjaye/go_tgbot/pkg/database"
	"os"
	"reflect"
)

func TelegramBot() {
	if err := database.InitLastPost(); err != nil {
		panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
			switch update.Message.Text {
			case "/start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi, i'm a HabrAboutGo bot, i can search new posts about go")
				bot.Send(msg)
			case "/go":
				msgs, err := database.GetUrls()
				if err != nil {
					panic(err)
				}
				if len(msgs) > 0 {
					for _, el := range msgs {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, el)
						bot.Send(msg)
					}
				}
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "incorrect command")
				bot.Send(msg)
			}
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Use the command")
			bot.Send(msg)
		}
	}
}
