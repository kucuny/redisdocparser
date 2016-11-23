package main

import (
	"fmt"
	gq "github.com/PuerkitoBio/goquery"
)

func main() {
	indexDoc, err := gq.NewDocument("http://redis.io/commands")

	if err != nil {
		fmt.Println(err)
	}

	indexDoc.Find(".container").Eq(2).Find("ul").Find("li").Each(func(i int, s *gq.Selection) {
		group, _ := s.Attr("data-group")
		cmdName, _ := s.Attr("data-name")
		url, _ := s.Find("a").Attr("href")
		summary := s.Find(".summary").Text()
		fmt.Println(group, cmdName, url, summary)
	})

	viewDoc, err := gq.NewDocument("http://redis.io/commands/append")

	if err != nil {
		fmt.Println(err)
	}

	viewDoc.Find(".site-content .text").Each(func(i int, s *gq.Selection) {
		cmdName := s.Find(".command .name").Text()
		args := s.Find(".command .arg").Text()
		availableVersion := s.Find(".metadata p strong").Eq(0).Text()
		fmt.Println(cmdName, availableVersion, args)
	})
}
