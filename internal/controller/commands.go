package controller

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	startCMD = "start"
	aboutCMD = "about"
)

const (
	startMessage          = "Привет! Я бот для ведения заметок. Что хочешь сделать?"
	aboutMessage          = "Ты можешь оставлять заметки, а я буду их хранить. Приятного пользования!"
	unknownCommandMessage = "Такой команды не существует."
)

func (c *controller) handleCommand(message *tgbotapi.Message) {
	switch message.Command() {
	case startCMD:
		c.handleStart(message)
	case aboutCMD:
		c.handleAbout(message)
	default:
		c.handleUnknownCommand(message)
	}
}

func (c *controller) handleStart(message *tgbotapi.Message) {
	c.bot.Send(tgbotapi.NewMessage(message.Chat.ID, startMessage))
}

func (c *controller) handleAbout(message *tgbotapi.Message) {
	c.bot.Send(tgbotapi.NewMessage(message.Chat.ID, aboutMessage))
}

func (c *controller) handleUnknownCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, unknownCommandMessage)
	c.bot.Send(msg)
}
