package main

import (
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"

	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
)

func Inita() {
	// --------------------------在此下书写插件内容--------------------------
	en := control.Register("demo", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Help:             "help from demo",
	})
	en.OnCommand("demo", zero.AdminPermission).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text("回复"))
		})
	// --------------------------在此上书写插件内容--------------------------
}

func main() {
	// stub!
}
