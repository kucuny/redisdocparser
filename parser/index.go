package parser

import gq "github.com/PuerkitoBio/goquery"

type RedisIndex struct {
	indexParser
}

func NewRedisIndex() RedisIndex {
	return RedisIndex{}
}

func (r RedisIndex) Run() map[Group][]Index {
	doc, err := r.parse()

	if err != nil {
		return nil
	}

	groupInfo := make(map[string]Group)
	doc.Find(".container").Eq(1).Find("select").Find("option").Each(func(i int, s *gq.Selection) {
		code, _ := s.Attr("value")
		name := s.Text()

		if code != "" {
			groupInfo[code] = Group{Code: code, Name: name}
		}
	})

	eachGroupCmdList := make(map[Group][]Index)
	doc.Find(".container").Eq(2).Find("ul").Find("li").Each(func(i int, s *gq.Selection) {
		group, _ := s.Attr("data-group")
		cmdName, _ := s.Attr("data-name")
		url, _ := s.Find("a").Attr("href")
		summary := s.Find(".summary").Text()

		realGroup := groupInfo[group]

		eachGroupCmdList[realGroup] = append(eachGroupCmdList[realGroup], Index{Group: realGroup, CmdName: cmdName, URL: url, Summary: summary})
	})

	return eachGroupCmdList
}

func (r RedisIndex) parse() (*gq.Document, error) {
	return gq.NewDocument(RedisIndexURL)
}
