package parser

import (
	"imooc.com/Google资深工程师深度讲解Go语言/爬虫/crawier/engine"
	"regexp"
)

// 获取city页面的用户名称及URL
const cityRe  = `<a href="(http://album.zhenai.com/u/[0-9]*)" target="_blank">([^<]*)</a>`

func ParseCity(content []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matchs := re.FindAllSubmatch(content, -1)

	results := engine.ParseResult{}
	for _, m := range matchs {
		results.Items = append(results.Items, "User " + string(m[2]))
		//results.Request = append(results.Request, engine.Request{string(m[1]), engine.NilParser})
		// 将用户名从此处传到ParseProfile中，不改变原来结构需用到匿名函数
		results.Request = append(results.Request, engine.Request{Url: string(m[1]),
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, string(m[2]))
				},
		})


	}

	return results
}
