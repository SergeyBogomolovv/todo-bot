package controller

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *controller) handleCallbackQuery(query *tgbotapi.CallbackQuery) {
	callback := tgbotapi.NewCallback(query.ID, query.Data)
	c.bot.Request(callback)

	msg := tgbotapi.NewMessage(query.Message.Chat.ID, query.Data)
	c.bot.Send(msg)
}
