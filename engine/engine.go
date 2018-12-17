package engine

import (
	"log"

	fetcher "github.com/zhiyxu/golearn/project/crawler-concurrent/fetcher"
)

type SimpleEngineer struct {
}

func (SimpleEngineer) Run(seeds ...Request) {
	var request []Request
	for _, r := range seeds {
		request = append(request, r)
	}

	for len(request) > 0 {
		r := request[0]
		request = request[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}

		request = append(request,
			parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Get Item: %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching Url: %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetch error "+
			"fetch url %s: %v",
			r.Url, err)

		return ParseResult{}, nil
	}

	return r.ParserFunc(body), nil
}
