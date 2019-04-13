package engine

import (
	"imooc.com/Google资深工程师深度讲解Go语言/爬虫/crawier/fetcher"
	"log"
)

// 简单的非单任务引擎
type SimpleEngine struct {}

func (e SimpleEngine) Run(seed ...Request)  {
	var requests  []Request // 维护一个requests队列
	for _, i := range seed {
		requests = append(requests, i)
	}

	for len(requests) > 0  {
		r := requests[0]
		requests = requests[1:]

		// 将解析url，获取body，包装ParseResult打包成一个work函数
		parseResult, err := work(r)
		if err != nil{
			continue
		}

		requests = append(requests, parseResult.Request...)

		// 打印itom
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func work(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher : Error fetching url %s : %v", r.Url, err )
		return ParseResult{}, err
	}

	return  r.ParserFunc(body), nil
}