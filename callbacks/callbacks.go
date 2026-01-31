package callbacks

import (
	"feedback_bot/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func Reply(ctx *th.Context, query telego.CallbackQuery) error {
	// todo
	return nil
}

func Ban(ctx *th.Context, query telego.CallbackQuery) error {
	data := strings.Split(query.Data, ":")
	id, err := strconv.Atoi(data[1])
	if err != nil {
		return err
	}

	utils.BanListCache.BanUser(int64(id))

	ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(query.From.ID), fmt.Sprintf("%v successfully blocked in bot", id)))
	return nil
}
