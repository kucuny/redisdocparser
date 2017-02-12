package parser

import (
	gq "github.com/PuerkitoBio/goquery"
)

const redisBaseURL string = "http://redis.io"
const RedisIndexURL string = redisBaseURL + "/commands"
const RedisViewURL string = redisBaseURL + "%s"

type Group struct {
	Code string
	Name string
}

type Index struct {
	Group   Group
	CmdName string
	URL     string
	Summary string
}

type redisVersion struct {
	Major    int
	Minor    int
	Revision int
}

type View struct {
	Group            Group
	CmdName          string
	Args             []string
	AvailableVersion redisVersion
	TimeComplexity   string
	URL              string
}

type indexParser interface {
	parse() (*gq.Document, error)
	Run() []Index
}

type viewParser interface {
	parse(url string) (*gq.Document, error)
	Run(cmdList []Index)
}
