package main

import (
	"github.com/779789571/gin-vue-admin/server/cmd"
	"github.com/779789571/gin-vue-admin/server/core"
	"github.com/779789571/gin-vue-admin/server/global"
	"github.com/779789571/gin-vue-admin/server/initialize"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"os"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.GVA_VP = core.Viper()      // 初始化Viper
	global.GVA_LOG = core.Zap()       // 初始化zap日志库
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.GVA_DB != nil {
		initialize.RegisterTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	//core.RunWindowsServer()
	app := cli.NewApp()
	app.Name = "geranium"
	app.Usage = "convince，convenient，construction"
	app.Commands = []*cli.Command{&cmd.Web, &cmd.Scan} //&cmd.Scan
	app.Flags = append(app.Flags, cmd.Web.Flags...)
	app.Flags = append(app.Flags, cmd.Scan.Flags...)
	//app.Flags = append(app.Flags) //cmd.Scan.Flags...
	err := app.Run(os.Args)
	if err != nil {
		global.GVA_LOG.Error("app start error", zap.Any("err", err))
	}
}
