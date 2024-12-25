package controller

import (
	"log/slog"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type controller struct {
	notes  NotesService
	bot    *tgbotapi.BotAPI
	log    *slog.Logger
	states map[int64]UserState
}

func New(log *slog.Logger, bot *tgbotapi.BotAPI, notes NotesService) *controller {
	return &controller{notes: notes, log: log, bot: bot, states: make(map[int64]UserState)}
}

func (c *controller) HandleUpdate(update tgbotapi.Update) {
	switch {
	case update.Message != nil && update.Message.IsCommand():
		c.handleCommand(update.Message)
	case update.Message != nil:
		c.handleMessage(update.Message)
	case update.CallbackQuery != nil:
		c.handleCallbackQuery(update.CallbackQuery)
	default:
		c.log.Debug("Unhandled update", "update", update)
	}
}

func (c *controller) handleCommand(message *tgbotapi.Message) {
	chatID := message.Chat.ID
	switch message.Command() {
	case startCMD:
		c.handleStart(message)
	case aboutCMD:
		c.handleAbout(message)
	case noteCMD:
		c.handleNote(message)
	case randCMD:
		c.handleRandomNote(message)
	default:
		c.bot.Send(tgbotapi.NewMessage(chatID, unknownCommand))
	}
}

func (c *controller) handleMessage(message *tgbotapi.Message) {
	chatID := message.Chat.ID
	text := message.Text
	state, exists := c.states[chatID]

	switch text {
	case cancelText:
		c.handleCancel(message)
	case noteLabel:
		c.handleNote(message)
	case randomNoteLabel:
		c.handleRandomNote(message)
	default:
		if !exists {
			c.bot.Send(tgbotapi.NewMessage(chatID, unknownMessage))
			return
		}
		switch state.State {
		case WaitingForTitle:
			c.handleNoteTitleInput(message)
		case WaitingForContent:
			c.handleNoteContentInput(message)
		}
	}
}

func (c *controller) handleCallbackQuery(query *tgbotapi.CallbackQuery) {
	data := query.Data
	switch {
	case data == cancelData:
		c.handleCancel(query.Message)
	case checkIsRemoveQuery(data):
		c.handleRemoveNote(query)
	}
}

func checkIsRemoveQuery(data string) bool {
	return strings.HasPrefix(data, removeNotePrefix)
}
