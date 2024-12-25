package controller

import (
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type NotesService interface{}

type controller struct {
	notes NotesService
	bot   *tgbotapi.BotAPI
	log   *slog.Logger
}

func New(log *slog.Logger, bot *tgbotapi.BotAPI, notes NotesService) *controller {
	return &controller{notes: notes, log: log, bot: bot}
}

func (c *controller) HandleUpdate(update tgbotapi.Update) {
	switch {
	case update.Message != nil:
		if update.Message.IsCommand() {
			c.handleCommand(update.Message)
		} else {
			c.handleMessage(update.Message)
		}
	case update.CallbackQuery != nil:
		c.handleCallbackQuery(update.CallbackQuery)
	default:
		c.log.Debug("Unhandled update", "update", update)
	}
}
