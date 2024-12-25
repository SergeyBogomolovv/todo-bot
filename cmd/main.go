package main

import (
	"context"
	"hello-bot/internal/app"
	"hello-bot/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	cfg := config.New()
	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		panic(err)
	}

	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	app := app.New(bot, log)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go app.Start(ctx)
	<-ctx.Done()
	app.Stop()
}
