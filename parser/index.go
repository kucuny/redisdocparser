package parser

import (
	gq "github.com/PuerkitoBio/goquery"
)

type RedisIndex struct {
	indexParser
}

func NewRedisIndex() RedisIndex {
	return RedisIndex{}
}

func (r RedisIndex) Run() []Index {
	doc, err := r.parse()

	if err != nil {
		return nil
	}

	var result []Index

	doc.Find(".container").Eq(2).Find("ul").Find("li").Each(func(i int, s *gq.Selection) {
		group, _ := s.Attr("data-group")
		cmdName, _ := s.Attr("data-name")
		url, _ := s.Find("a").Attr("href")
		summary := s.Find(".summary").Text()

		result = append(result, Index{Group: group, CmdName: cmdName, Url: url, Summary: summary})
	})

	return result
}

func (r RedisIndex) parse() (*gq.Document, error) {
	return gq.NewDocument(RedisIndexURL)
}
