package collect

import (
	"fmt"
	"github.com/779789571/gin-vue-admin/server/collect/online"
	"github.com/779789571/gin-vue-admin/server/core"
	"github.com/779789571/gin-vue-admin/server/global"
	"github.com/urfave/cli/v2"
	"strings"
)

func Scanner(ctx *cli.Context) error {
	//设置时间周期 小时
	var intervalHour float64
	var proxy string
	if ctx.IsSet("time") {
		intervalHour = ctx.Float64("time")
	}
	if ctx.IsSet("proxy") {
		proxy = strings.ToLower("proxy")
	}
	var plugins string
	if ctx.IsSet("plugins") || ctx.IsSet("p") {
		plugins = strings.ToLower(ctx.String("plugins"))
		switch plugins {
		case "online":
			print(core.Banner + "\n")
			global.GVA_LOG.Info(fmt.Sprintf("Repeat every %v hours\n", intervalHour))
			global.GVA_LOG.Info("Start searching..\n")

			online.SearchRunner(intervalHour, proxy)
		}
	}

	return nil
}
