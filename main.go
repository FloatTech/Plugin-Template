package main

import (
	"fmt"
	"unsafe"

	ctrl "github.com/fumiama/ZeroBot-Hook/control"
	zero "github.com/fumiama/ZeroBot-Hook/hook"
	"github.com/fumiama/ZeroBot-Hook/hook/message"
)

func Inita() {
	// -------------在此下书写插件内容-------------
	en := ctrl.Register("demo", &ctrl.Options{
		DisableOnDefault: false,
		Help:             "help from demo",
	})
	en.OnCommand("demo", zero.AdminPermission).SetBlock(true).SecondPriority().
		Handle(func(ctx *zero.Ctx) {
			fmt.Println("msg recv.")
			ctx.SendChain(message.Text("回复"))
		})
	// -------------在此上书写插件内容-------------
}

// 以下勿动
// Hook 改变本插件的环境变量以加载插件
func Hook(botconf interface{}, apicallers interface{}, hooknew interface{},
	matlist interface{}, matlock interface{}, defen interface{},
	reg interface{}, del interface{},
	sndgrpmsg interface{}, sndprivmsg interface{}, getmsg interface{},
	parsectx interface{},
	custnode interface{}, pasemsg interface{}, parsemsgfromarr interface{},
) {
	zero.Hook(botconf, apicallers, hooknew, matlist, matlock, defen)
	rd := getdata(&reg)
	dd := getdata(&del)
	ctrl.Register = *(*(func(service string, o *ctrl.Options) *zero.Engine))(unsafe.Pointer(&rd))
	ctrl.Delete = *(*(func(service string)))(unsafe.Pointer(&dd))
	zero.HookCtx(sndgrpmsg, sndgrpmsg, getmsg, parsectx)
	message.HookMsg(custnode, pasemsg, parsemsgfromarr)
	IsHooked = true
	// fmt.Printf("[plugin]set reg: %x, del: %x\n", ctrl.Register, ctrl.Delete)
}

// IsHooked 已经 hook 则不再重复 hook
var IsHooked bool

// 没有方法的interface
type eface struct {
	_type uintptr
	data  unsafe.Pointer
}

func getdata(ptr *interface{}) unsafe.Pointer {
	return (*eface)(unsafe.Pointer(ptr)).data
}
