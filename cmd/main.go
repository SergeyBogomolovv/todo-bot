package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/SergeyBogomolovv/notes-bot/internal/app"
	"github.com/SergeyBogomolovv/notes-bot/internal/config"
	"github.com/SergeyBogomolovv/notes-bot/pkg/bot"
	"github.com/SergeyBogomolovv/notes-bot/pkg/db"
)

func main() {
	cfg := config.New()
	bot := bot.MustNewBot(cfg.Token)
	db := db.MustNewPostgres(cfg.PostgresURL)

	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	app := app.New(log, bot, db)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go app.Start(ctx)
	<-ctx.Done()
	app.Stop()
}
