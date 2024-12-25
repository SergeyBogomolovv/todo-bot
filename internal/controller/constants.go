package controller

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	startCMD = "start"
	aboutCMD = "about"
	noteCMD  = "note"
	randCMD  = "random"
	allCMD   = "all"
)

const (
	startMessage            = "Привет! Я бот для ведения заметок. Что хочешь сделать?"
	aboutMessage            = "Ты можешь оставлять заметки, а я буду их хранить. Приятного пользования!"
	unknownMessage          = "Неизвестная команда. Используйте /note для создания заметки."
	unknownCommand          = "Такой команды не существует."
	noteTitleMessage        = "Введите название заметки:"
	emptyTitleMessage       = "Название заметки не может быть пустым. Попробуйте ещё раз."
	noteContentMessage      = "Введите текст заметки:"
	emptyContentMessage     = "Заметка не может быть пустой. Попробуйте ещё раз."
	failedToSaveNoteMessage = "Не удалось сохранить заметку. Попробуйте позже."
	noteSavedMessage        = "Заметка успешно сохранена!"
	noNotesMessage          = "Заметок пока нет. Создайте заметку с помощью команды /note."
	failedToGetNotesMessage = "Не удалось получить заметки. Попробуйте позже."
	failedToGetNoteMessage  = "Не удалось получить заметку. Попробуйте позже."
	noteRemovedMessage      = "Заметка удалена."
	failedToRemoveMessage   = "Не удалось удалить заметку. Попробуйте позже."
)

const (
	cancelText      = "Отмена"
	cancelData      = "cancel"
	noteLabel       = "Новая заметка"
	randomNoteLabel = "Случайная заметка"
	allNotesLabel   = "Все заметки"
)

const (
	removeNotePrefix = "remove:"
)

var (
	cancelButton = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(cancelText),
		),
	)
	mainMenu = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(noteLabel),
			tgbotapi.NewKeyboardButton(randomNoteLabel),
			tgbotapi.NewKeyboardButton(allNotesLabel),
		),
	)
)

func newDeleteKeyboard(noteID int64) tgbotapi.InlineKeyboardMarkup {
	data := fmt.Sprintf("%s%d", removeNotePrefix, noteID)
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Удалить", data),
		),
	)
}
