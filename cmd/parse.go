package main

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
)

// type paifuList struct {
// 	south3 []paifu
// 	south4 []paifu
// 	east3  []paifu
// 	east4  []paifu
// }

type Paifu struct {
	Date     string `json:"date"`
	GameType string `json:"game_type"`
	LogID    string `json:"log_id"`
}

func ReadSingleFile(filePath string) ([]Paifu, error) {
	var datas = make([]Paifu, 0)
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		// ReadLine is a low-level line-reading primitive.
		// Most callers should use ReadBytes('\n') or ReadString('\n') instead or use a Scanner.
		bytes, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		line := string(bytes)

		lineContentList := strings.Split(line, "|")
		timeMin := strings.TrimSpace(lineContentList[0])

		doc, _ := htmlquery.Parse(strings.NewReader(strings.TrimSpace(lineContentList[3])))
		link := htmlquery.SelectAttr(htmlquery.FindOne(doc, "//a/@href"), "href")
		u, _ := url.Parse(link)
		logId := u.Query().Get("log")
		logIdSplit := strings.Split(logId, "-")

		logDate := strings.Replace(logIdSplit[0], "gm", "", -1)
		fmt.Println(fmt.Sprintf("%v%v", logDate, strings.Split(timeMin, ":")[1]))
		date, _ := time.Parse("200601021504", fmt.Sprintf("%v%v", logDate, strings.Split(timeMin, ":")[1]))

		datas = append(datas, Paifu{
			LogID:    logId,
			GameType: logIdSplit[1],
			Date:     date.Format(time.RFC3339),
		})

		// 	switch data.GameType {
		// 	case "00b1":
		// 		c.data.east3 = append(c.data.east3, data)
		// 	case "00f1":
		// 		c.data.east3 = append(c.data.south3, data)

		// 	case "00b9":
		// 		c.data.south3 = append(c.data.south3, data)
		// 	case "00f9":
		// 		c.data.south3 = append(c.data.south3, data)

		// 	case "00a1":
		// 		c.data.east4 = append(c.data.east4, data)
		// 	case "00e1":
		// 		c.data.east4 = append(c.data.east4, data)

		// 	case "00a9":
		// 		c.data.south4 = append(c.data.south4, data)
		// 	case "00e9":
		// 		c.data.south4 = append(c.data.south4, data)
		// 	}

	}
	return datas, nil
}
