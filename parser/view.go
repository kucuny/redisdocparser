package parser

import (
	"fmt"
	"regexp"
	"strconv"

	gq "github.com/PuerkitoBio/goquery"
)

const RegAvailableVersion string = `(?P<major>\d+)\.(?P<minor>\d+)\.(?P<revision>\d+)\.`
const RegTimeComplexity string = `(O\(\w+\))`

type RedisView struct {
	viewParser
}

func NewRedisView() RedisView {
	return RedisView{}
}

func (r RedisView) Run(cmdList []Index) []View {
	var result []View

	for k, cmd := range cmdList {
		url := fmt.Sprintf(RedisViewURL, cmd.URL)
		doc, err := r.parse(url)

		if err != nil {
			fmt.Println(err)
		}

		doc.Find(".site-content .text").Each(func(i int, s *gq.Selection) {
			var args []string
			cmdName := s.Find(".command .name").Text()
			tmpArgs := s.Find(".command .arg")
			for j := range tmpArgs.Nodes {
				args = append(args, tmpArgs.Eq(j).Text())
			}
			availableVersion := r.parseAvailableVersion(s.Find(".metadata p strong").Eq(0).Text())
			timeComplexity := r.parseTimeComplexity(s.Find(".metadata p").Eq(1).Text())

			temp := View{
				Group:            cmd.Group,
				CmdName:          cmdName,
				AvailableVersion: availableVersion,
				TimeComplexity:   timeComplexity,
				Args:             args,
				URL:              url,
			}
			result = append(result, temp)
		})

		if k > 2 {
			return result
		}
	}

	return result
}

func (r RedisView) parse(url string) (*gq.Document, error) {
	return gq.NewDocument(url)
}

func (r RedisView) parseAvailableVersion(version string) redisVersion {
	re := regexp.MustCompile(RegAvailableVersion)
	parsedVersion := re.FindStringSubmatch(version)

	major, _ := strconv.Atoi(parsedVersion[1])
	minor, _ := strconv.Atoi(parsedVersion[2])
	revision, _ := strconv.Atoi(parsedVersion[3])

	return redisVersion{Major: major, Minor: minor, Revision: revision}
}

func (r RedisView) parseTimeComplexity(tc string) string {
	re := regexp.MustCompile(RegTimeComplexity)
	return re.FindString(tc)
}

// func (r RedisView) generate(viewDetail []View) {
// 	var resultEachVersion map[int]map[int]map[int][]View
// 	var resultEachGroup map[string][]View

// 	for _, cmd := range viewDetail {
// 		if _, ok := resultEachVersion[cmd.AvailableVersion.Major]; !ok {
// 			resultEachVersion = make(map[int][int][]View)
// 		}

// 		if _, ok := resultEachVersion[cmd.AvailableVersion.Major][cmd.AvailableVersion.Minor]; !ok {
// 			resultEachVersion[cmd.AvailableVersion.Major] = make([int][]View)
// 		}

// 		if major, ok := resultEachVersion[cmd.AvailableVersion.Major]; ok {
// 			if minor, ok := resultEachVersion[major][cmd.AvailableVersion.Minor]; ok {
// 				if revision, ok := resultEachVersion[major][minor][cmd.AvailableVersion.Revision]; ok {
// 					resultEachVersion[major][minor][revision] = append(resultEachVersion[major][minor][revision], cmd)
// 				} else {
// 					resultEachVersion[major][minor] = make(map[int][]View)
// 					resultEachVersion[major][minor][revision] = cmd
// 				}
// 			}
// 		}
// 	}
// }
