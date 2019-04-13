package main

import (
	"imooc.com/Google资深工程师深度讲解Go语言/爬虫/crawier/engine"
	"imooc.com/Google资深工程师深度讲解Go语言/爬虫/crawier/zhenai/parser"
	"imooc.com/Google资深工程师深度讲解Go语言/爬虫/crawier/scheduler"
)

func main(){
	e := engine.ConcurrentEngine{Scheduler: &scheduler.SimpleScheduler{}, WorkCount: 100}
	e.Run(engine.Request{
			"http://www.zhenai.com/zhenghun",
		parser.ParseCityList,
	})
}