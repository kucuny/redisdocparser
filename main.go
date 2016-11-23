package main

import (
	"fmt"
	"github.com/kucuny/redisdocparser/parser"
)

func main() {
	index := parser.NewRedisIndex()
	cmdList := index.Run()

	view := parser.NewRedisView()
	cmdDeatilList := view.Run(cmdList)

	fmt.Println(cmdDeatilList)
}
