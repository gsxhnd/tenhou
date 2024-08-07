package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

var rootCmd = cli.NewApp()

func init() {
	rootCmd.Usage = ""
	rootCmd.Commands = []*cli.Command{
		downloadDayCmd,
		downloadRecentCmd,
		fullHtml2DB,
		recentHtml2DB,
		apiCmd,
	}
}

func main() {
	err := rootCmd.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
