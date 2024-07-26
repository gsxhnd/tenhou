package main

import (
	"github.com/gsxhnd/tenhou/api/di"
	"github.com/urfave/cli/v2"
)

var apiCmd = &cli.Command{
	Name: "api",
	Action: func(ctx *cli.Context) error {
		app, err := di.InitApp()
		if err != nil {
			return err
		}
		return app.Run()
	},
}
