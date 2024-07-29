package main

import (
	"github.com/gsxhnd/tenhou/di"
	"github.com/urfave/cli/v2"
)

var apiCmd = &cli.Command{
	Name: "api",
	Action: func(ctx *cli.Context) error {
		app, err := di.InitApp()
		if err != nil {
			panic(err)
		}
		return app.Run()
	},
}
