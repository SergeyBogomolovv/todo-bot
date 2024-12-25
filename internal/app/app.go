package app

import (
	"context"
	"log/slog"

	"github.com/SergeyBogomolovv/notes-bot/internal/controller"
	"github.com/SergeyBogomolovv/notes-bot/internal/service"
	"github.com/SergeyBogomolovv/notes-bot/internal/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
)

type Controller interface {
	HandleUpdate(tgbotapi.Update)
}

type application struct {
	bot        *tgbotapi.BotAPI
	log        *slog.Logger
	controller Controller
}

func New(log *slog.Logger, bot *tgbotapi.BotAPI, db *sqlx.DB) *application {
	storage := storage.New(db)
	service := service.New(log, storage)
	controller := controller.New(log, bot, service)

	return &application{
		bot:        bot,
		log:        log,
		controller: controller,
	}
}

func (a *application) Start(ctx context.Context) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := a.bot.GetUpdatesChan(u)

	a.log.Info("Bot started", "bot", a.bot.Self.UserName)

	for {
		select {
		case <-ctx.Done():
			return
		case update := <-updates:
			a.controller.HandleUpdate(update)
		}
	}
}

func (a *application) Stop() {
	a.bot.StopReceivingUpdates()
	a.log.Info("Bot stopped")
}
