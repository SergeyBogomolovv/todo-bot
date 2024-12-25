package app

import (
	"context"
	"hello-bot/internal/controller"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Controller interface {
	HandleUpdate(tgbotapi.Update)
}

type application struct {
	bot        *tgbotapi.BotAPI
	log        *slog.Logger
	controller Controller
}

func New(bot *tgbotapi.BotAPI, log *slog.Logger) *application {
	controller := controller.New(log, bot, nil)

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
