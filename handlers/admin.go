package handlers

import (
	"feedback_bot/utils"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func Ban(ctx *th.Context, update telego.Update) error {
	data := strings.Split(update.Message.Text, " ")
	if len(data) < 2 {
		msg := tu.Message(
			tu.ID(update.Message.From.ID),
			"Invalid format: /ban <ID>",
		)

		ctx.Bot().SendMessage(ctx, msg)
		return nil
	}

	id, err := strconv.Atoi(data[1])
	if err != nil {
		return err
	}

	adminId, err := strconv.Atoi(os.Getenv("TELEGRAM_ID"))
	if err != nil {
		return err
	}

	if int64(adminId) == int64(id) {
		ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(update.Message.From.ID), "You can't ban yourself bro"))
		return nil
	}

	utils.BanListCache.BanUser(int64(id))

	msg := tu.Message(
		tu.ID(update.Message.From.ID),
		fmt.Sprintf("User %v successfully banned", id),
	)

	ctx.Bot().SendMessage(ctx, msg)

	return nil
}

func Unban(ctx *th.Context, update telego.Update) error {
	data := strings.Split(update.Message.Text, " ")
	if len(data) < 2 {
		msg := tu.Message(
			tu.ID(update.Message.From.ID),
			"Invalid format: /unban <ID>",
		)

		ctx.Bot().SendMessage(ctx, msg)
		return nil
	}

	id, err := strconv.Atoi(data[1])
	if err != nil {
		return err
	}

	utils.BanListCache.UnbanUser(int64(id))

	msg := tu.Message(
		tu.ID(update.Message.From.ID),
		fmt.Sprintf("User %v successfully unbanned", id),
	)

	ctx.Bot().SendMessage(ctx, msg)

	return nil
}
