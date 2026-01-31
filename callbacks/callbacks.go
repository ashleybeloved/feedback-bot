package callbacks

import (
	"feedback_bot/middleware"
	"feedback_bot/utils"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func Reply(ctx *th.Context, query telego.CallbackQuery) error {
	data := strings.Split(query.Data, ":")
	id, err := strconv.Atoi(data[1])
	if err != nil {
		return err
	}

	middleware.State = fmt.Sprintf("await_reply:%v", id)

	ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(query.From.ID), fmt.Sprintf("Message for reply %v (\"cancel\" to cancel):", id)))

	return ctx.Bot().AnswerCallbackQuery(ctx, tu.CallbackQuery(query.ID))
}

func Ban(ctx *th.Context, query telego.CallbackQuery) error {
	data := strings.Split(query.Data, ":")
	id, err := strconv.Atoi(data[1])
	if err != nil {
		return err
	}

	adminId, err := strconv.Atoi(os.Getenv("TELEGRAM_ID"))
	if err != nil {
		return err
	}

	if int64(adminId) == int64(id) {
		ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(query.From.ID), "You can't ban yourself bro"))
		return nil
	}

	utils.BanListCache.BanUser(int64(id))

	ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(query.From.ID), fmt.Sprintf("%v successfully blocked in bot", id)))

	return ctx.Bot().AnswerCallbackQuery(ctx, tu.CallbackQuery(query.ID))
}
