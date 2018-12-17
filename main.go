package main

import (
	"fmt"

	"github.com/zhiyxu/golearn/project/crawler-concurrent/scheduler"

	"github.com/zhiyxu/golearn/project/crawler-concurrent/engine"
	"github.com/zhiyxu/golearn/project/crawler-concurrent/zhenai/parser"

	"github.com/zhiyxu/golearn/project/crawler-concurrent/fetcher"
)

func main() {

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	//fetchContent("http://album.zhenai.com/u/22843760")

}

func fetchContent(url string) {

	contents, err := fetcher.Fetch(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", contents)
}
