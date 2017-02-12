package main

import (
	"fmt"
	"sync"

	"github.com/kucuny/redisdocparser/parser"
)

func main() {
	index := parser.NewRedisIndex()
	cmdList := index.Run()

	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	view := parser.NewRedisView()
	var cmdDetailList []parser.View
	for _, groupCmd := range cmdList {
		wg.Add(1)
		go func() {
			defer wg.Done()
			groupCmdViewList := view.Run(groupCmd)

			mu.Lock()
			defer mu.Unlock()
			cmdDetailList = append(cmdDetailList, groupCmdViewList...)
		}()
	}

	wg.Wait()
	fmt.Println(cmdDetailList)
}
