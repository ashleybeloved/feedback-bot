package middleware

import (
	"feedback_bot/utils"
	"os"
	"strconv"
	"strings"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func UserMiddleware(ctx *th.Context, update telego.Update) error {
	var userid int64

	if update.CallbackQuery != nil {
		userid = update.CallbackQuery.From.ID
	} else if update.Message != nil {
		userid = update.Message.From.ID
	} else {
		return ctx.Next(update)
	}

	if utils.BanListCache.IsBanned(userid) {
		msg := tu.Message(
			tu.ID(userid),
			"You are banned in bot",
		)

		ctx.Bot().SendMessage(ctx, msg)
		return nil
	}

	return ctx.Next(update)
}

func AdminMiddleware(ctx *th.Context, update telego.Update) error {
	var userid int64

	if update.CallbackQuery != nil {
		userid = update.CallbackQuery.From.ID
	} else if update.Message != nil {
		userid = update.Message.From.ID
	} else {
		return ctx.Next(update)
	}

	adminId, err := strconv.Atoi(os.Getenv("TELEGRAM_ID"))
	if err != nil {
		return err
	}

	if int64(adminId) != userid {
		if update.Message != nil && update.Message.Text != "" {
			if strings.HasPrefix(update.Message.Text, "/") {
				return nil
			}

			ctx.Next(update)
		}

		return nil
	}

	return ctx.Next(update)
}
