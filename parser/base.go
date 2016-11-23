package parser

import (
	gq "github.com/PuerkitoBio/goquery"
)

const redisBaseURL string = "http://redis.io"
const RedisIndexURL string = redisBaseURL + "/commands"
const RedisViewURL string = redisBaseURL + "%s"

type Index struct {
	Group   string
	CmdName string
	Url     string
	Summary string
}

type redisVersion struct {
	Major    int
	Minor    int
	Revision int
}

type View struct {
	Group            string
	CmdName          string
	Args             []string
	AvailableVersion redisVersion
	TimeComplexity   string
	Url              string
}

type indexParser interface {
	parse() (*gq.Document, error)
	Run() []Index
}

type viewParser interface {
	parse(url string) (*gq.Document, error)
	Run(cmdList []Index)
}
