package engine

import (
	"../fetcher"
	"fmt"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests := requests[1:]

		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			fmt.Printf("Fetcher：error fetching url %d：%v", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("items:%s", item)
		}

	}
}
