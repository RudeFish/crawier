package parser

import (
	"imooc.com/Google资深工程师深度讲解Go语言/爬虫/crawier/engine"
	"imooc.com/Google资深工程师深度讲解Go语言/爬虫/crawier/model"
	"regexp"
)

var re = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^<]*)</div>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}

	profile.Name = name
	// 提取人员信息
	match := re.FindAllSubmatch(contents, -1)
	if match != nil {
		profile.Marriage = string(match[0][1])
		//fmt.Printf("婚姻状况 ： %s\n" ,match[0][1])

		profile.Age = string(match[1][1])
		//fmt.Printf("年龄 ： %s\n" ,match[0][2])
		profile.Xizuo = string(match[2][1])
		profile.Height = string(match[3][1])
		profile.Weight = string(match[4][1])
		profile.Education = string(match[5][1])
		profile.Incomr = string(match[6][1])
		//profile.Occupation = string(match[7][1])

	}

	// 将信息放到Item中返回出去
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}