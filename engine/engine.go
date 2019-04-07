package engine

import (
	"imooc.com/Google资深工程师深度讲解Go语言/爬虫/crawier/fetcher"
	"log"
)

func Run(seed ...Request)  {
	var requests  []Request // 维护一个requests队列
	for _, i := range seed {
		requests = append(requests, i)
	}

	for len(requests) > 0  {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher : Error fetching url %s : %v", r.Url, err )
			continue // 遇到错误打出错误日志并运行下一条
		}

		parseResult :=  r.ParserFunc(body)
		requests = append(requests, parseResult.Request...)

		// 打印itom
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

