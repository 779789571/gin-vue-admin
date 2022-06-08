package cmd

import (
	"github.com/779789571/gin-vue-admin/server/collect"
	"github.com/779789571/gin-vue-admin/server/core"

	"github.com/urfave/cli/v2"
)

var Web = cli.Command{
	Name:        "web",
	Usage:       "Startup a web Service",
	Description: "Startup a web Service",
	Action:      core.RunWindowsServer,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "host, H",
			Value: "0.0.0.0",
			Usage: "web listen address",
		},
		&cli.IntFlag{
			Name:  "post, p",
			Value: 8000,
			Usage: "web listen port",
		},
	},
}

var Scan = cli.Command{
	Name:        "scan",
	Usage:       "Start to scan ",
	Description: "start to scan",
	Action:      collect.Scanner,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "plugins",
			Aliases: []string{"p"},
			Value:   "",
			Usage:   "scan plugins: online,vulner,gather all",
		},
		&cli.IntFlag{
			Name:    "time",
			Aliases: []string{"t"},
			Value:   24,
			Usage:   "scan interval(hour)",
		},
		&cli.StringFlag{
			Name:    "proxy",
			Aliases: []string{"px"},
			Value:   "",
			Usage:   "web proxy",
		},
	},
}
