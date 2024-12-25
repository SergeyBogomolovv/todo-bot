package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func MustNewBot(token string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	return bot
}
