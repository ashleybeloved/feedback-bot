package handlers

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func Start(ctx *th.Context, update telego.Update) error {
	adminUsername := os.Getenv("TELEGRAM_USERNAME")

	msg := tu.Message(
		tu.ID(update.Message.From.ID),
		"Hi! You can sent any message for @"+adminUsername+" with this bot, just type or sent anything you want",
	)

	ctx.Bot().SendMessage(ctx, msg)

	return nil
}

func AnyMessage(ctx *th.Context, update telego.Update) error {
	adminId, err := strconv.Atoi(os.Getenv("TELEGRAM_ID"))
	if err != nil {
		return err
	}

	// User Information

	userId := update.Message.From.ID
	username := update.Message.From.Username
	firstname := update.Message.From.FirstName
	lastname := update.Message.From.LastName
	langCode := update.Message.From.LanguageCode

	if username != "" {
		username = "@" + username
	} else {
		username = "no username"
	}

	// Forward Original Message to Admin and Info About User

	forwardMsg := telego.ForwardMessageParams{
		ChatID:     tu.ID(int64(adminId)),
		FromChatID: update.Message.Chat.ChatID(),
		MessageID:  update.Message.MessageID,
	}

	keyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Reply").WithCallbackData(fmt.Sprintf("reply:%v", userId)),
			tu.InlineKeyboardButton("Ban").WithCallbackData(fmt.Sprintf("ban:%v", userId)),
		),
	)

	infoMsg := tu.Message(
		tu.ID(int64(adminId)),
		fmt.Sprintf("ID: %v\nUsername: %v\nName: %v %v\nLanguage Code: %v", userId, username, firstname, lastname, langCode),
	).WithReplyMarkup(keyboard)

	ctx.Bot().ForwardMessage(ctx, &forwardMsg)

	ctx.Bot().SendMessage(ctx, infoMsg)

	// Message for user

	msg := tu.Message(
		update.Message.Chat.ChatID(),
		"Message successfully sended!",
	)

	ctx.Bot().SendMessage(ctx, msg)

	return nil
}
