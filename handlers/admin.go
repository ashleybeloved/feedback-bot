package handlers

import (
	"feedback_bot/utils"
	"fmt"
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
