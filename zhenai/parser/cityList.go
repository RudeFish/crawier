package parser

import (
	"imooc.com/Google资深工程师深度讲解Go语言/爬虫/crawier/engine"
	"regexp"
)

const cityList  = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]*)</a>`

func ParseCityList(contents []byte) engine.ParseResult  {
	// <a href="http://www.zhenai.com/zhenghun/aba" data-v-5e16505f>阿坝</a>
	// 小技巧 [^>]*> 匹配到下一个> , [^<]*< 匹配到下一个<
	re , _ := regexp.Compile(cityList)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	limit := 1
	for _, i := range matches {
		// 将城市名称放到resule
		result.Items = append(result.Items, string(i[2]))
		result.Request = append(result.Request, engine.Request{string(i[1]), ParseCity})
		//fmt.Printf("City: %-8s, URL: %s \n", i[2], i[1])
		// 控制爬取的页面数
		limit--
		if limit == 0 {
			break
		}
	}
	return result
	//fmt.Println(len(find))
}
