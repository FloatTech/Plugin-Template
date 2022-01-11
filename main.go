package main

import (
	"fmt"

	ctrl "github.com/FloatTech/zbpctrl"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func Inita() {
	// --------------------------在此下书写插件内容--------------------------
	en := ctrl.Register("demo", &ctrl.Options{
		DisableOnDefault: false,
		Help:             "help from demo",
	})
	en.OnCommand("demo", zero.AdminPermission).SetBlock(true).SecondPriority().
		Handle(func(ctx *zero.Ctx) {
			fmt.Println("msg recv.")
			ctx.SendChain(message.Text("回复"))
		})
	// --------------------------在此上书写插件内容--------------------------
}

func main() {
	// stub!
}
