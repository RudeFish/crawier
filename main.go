package main

import (
	"imooc.com/Google资深工程师深度讲解Go语言/爬虫/crawier/engine"
	"imooc.com/Google资深工程师深度讲解Go语言/爬虫/crawier/zhenai/parser"
)

func main(){
	engine.Run(engine.Request{
			"http://www.zhenai.com/zhenghun",
		parser.ParseCityList,
	})
}