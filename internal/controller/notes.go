package controller

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/SergeyBogomolovv/notes-bot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	ErrCantGetTitle = errors.New("can't get title")
)

func (c *controller) handleNote(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, noteTitleMessage)
	msg.ReplyMarkup = cancelButton
	c.states[message.Chat.ID] = UserState{State: WaitingForTitle}
	c.bot.Send(msg)
}

func (c *controller) handleNoteTitleInput(message *tgbotapi.Message) {
	chatID := message.Chat.ID
	title := message.Text

	if title == "" {
		c.bot.Send(tgbotapi.NewMessage(chatID, emptyTitleMessage))
		return
	}

	c.states[chatID] = UserState{State: WaitingForContent, Data: title}

	msg := tgbotapi.NewMessage(chatID, noteContentMessage)
	msg.ReplyMarkup = cancelButton
	c.bot.Send(msg)
}

func (c *controller) handleNoteContentInput(message *tgbotapi.Message) {
	chatID := message.Chat.ID
	content := message.Text

	if content == "" {
		c.bot.Send(tgbotapi.NewMessage(chatID, emptyContentMessage))
		return
	}

	title, err := getTitle(c.states[chatID])
	if err != nil {
		c.bot.Send(tgbotapi.NewMessage(chatID, failedToSaveNoteMessage))
		c.log.Error("Failed to save note", "error", err)
		return
	}

	_, err = c.notes.CreateNote(context.TODO(), title, content, message.From.ID)
	if err != nil {
		c.bot.Send(tgbotapi.NewMessage(chatID, failedToSaveNoteMessage))
		c.log.Error("Failed to save note", "error", err)
		return
	}

	delete(c.states, chatID)
	msg := tgbotapi.NewMessage(chatID, noteSavedMessage)
	msg.ReplyMarkup = mainMenu
	c.bot.Send(msg)
}

func (c *controller) handleRandomNote(message *tgbotapi.Message) {
	chatID := message.Chat.ID
	note, err := c.notes.PickRandomNote(context.TODO(), message.From.ID)
	if err != nil {
		if errors.Is(err, service.ErrNoNotes) {
			c.bot.Send(tgbotapi.NewMessage(chatID, noNotesMessage))
			return
		}
		c.bot.Send(tgbotapi.NewMessage(chatID, failedToGetNoteMessage))
		c.log.Error("Failed to get notes", "error", err)
		return
	}
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("%s\n\n%s", note.Title, note.Content))
	msg.ReplyMarkup = newDeleteKeyboard(note.ID)
	c.bot.Send(msg)
}

func (c *controller) handleRemoveNote(query *tgbotapi.CallbackQuery) {
	message := query.Message
	chatID := message.Chat.ID
	id, err := parseRemoveNoteId(query.Data)
	if err != nil {
		c.bot.Send(tgbotapi.NewMessage(chatID, failedToRemoveMessage))
		return
	}
	if err := c.notes.RemoveNote(context.TODO(), id); err != nil {
		if !errors.Is(err, service.ErrNoNotes) {
			c.log.Error("failed to remove note", "error", err)
		}
		c.bot.Send(tgbotapi.NewMessage(chatID, failedToRemoveMessage))
		return
	}

	c.bot.Send(tgbotapi.NewCallback(query.ID, noteRemovedMessage))

	msg := tgbotapi.NewMessage(chatID, noteRemovedMessage)
	c.bot.Send(msg)
	c.bot.Send(tgbotapi.NewDeleteMessage(chatID, message.MessageID))
}

func getTitle(state UserState) (string, error) {
	title, ok := state.Data.(string)
	if !ok {
		return "", ErrCantGetTitle
	}
	return title, nil
}

func parseRemoveNoteId(data string) (int64, error) {
	data, ok := strings.CutPrefix(data, removeNotePrefix)
	if !ok {
		return 0, errors.New("invalid command")
	}
	res, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return 0, err
	}
	return res, nil
}
