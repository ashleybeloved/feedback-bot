package main

import (
	"context"
	"feedback_bot/callbacks"
	"feedback_bot/handlers"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

var BannedUsers = make(map[int64]bool)

func main() {
	// Load .env secrets

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	botToken := os.Getenv("BOT_TOKEN")

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(ctx, nil)

	bh, _ := th.NewBotHandler(bot, updates)
	defer func() { _ = bh.Stop() }()

	// Handlers and callbacks

	bh.Handle(handlers.Start, th.CommandEqual("start"))
	bh.Handle(handlers.AnyMessage, th.AnyMessage())

	bh.HandleCallbackQuery(callbacks.Reply, th.CallbackDataContains("reply"))

	log.Print("Bot started on @", bot.Username())

	_ = bh.Start()
}
