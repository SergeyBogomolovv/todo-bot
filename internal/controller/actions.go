package controller

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *controller) handleStart(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, startMessage)
	msg.ReplyMarkup = mainMenu
	c.bot.Send(msg)
}

func (c *controller) handleAbout(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, aboutMessage)
	msg.ReplyMarkup = mainMenu
	c.bot.Send(msg)
}

func (c *controller) handleCancel(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Действий и так нет.")
	msg.ReplyMarkup = mainMenu
	_, exists := c.states[message.Chat.ID]
	if exists {
		msg.Text = "Действие отменено."
		delete(c.states, message.Chat.ID)
	}
	c.bot.Send(msg)
}
