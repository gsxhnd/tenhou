package main

import (
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
	"github.com/urfave/cli/v2"
)

var downloadRecentCmd = &cli.Command{
	Name:  "download_recent",
	Usage: "convert tenhou day data to csv",
	Action: func(ctx *cli.Context) error {
		queue, _ := queue.New(1, &queue.InMemoryQueueStorage{MaxSize: 10000})
		c := colly.NewCollector()
		c.Limit(&colly.LimitRule{
			Parallelism: 10,
			Delay:       1 * time.Second,
			RandomDelay: 5 * time.Second,
		})

		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Download zip from: ", r.URL)
		})

		c.OnResponse(func(r *colly.Response) {
			p := strings.Split(r.Request.URL.Path, "/")
			name1 := p[len(p)-1]
			name2 := strings.Split(name1, ".")[0] + ".html"
			dstPath := path.Join("./data/tenhou_html", name2)
			r.Save(dstPath)
		})

		var startDate = time.Now().AddDate(0, 0, -20)
		var endDate = time.Now().AddDate(0, 0, -10)

		for i := startDate; i.Before(endDate); i = i.AddDate(0, 0, 1) {
			year := i.Year()
			date := i.Format("20060102")
			var url = fmt.Sprintf("https://tenhou.net/sc/raw/dat/%v/scc%v.html.gz", year, date)
			fmt.Println(url)
			// queue.AddURL(url)
		}

		queue.Run(c)
		return nil
	},
}
