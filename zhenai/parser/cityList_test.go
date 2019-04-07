package parser

import (
	"testing"
	"io/ioutil"
	"fmt"
)

func TestParseCityList(t *testing.T) {
	contenes, err := ioutil.ReadFile("D:\\code\\Go\\itcast\\TheThirdPhase\\01mysqlDatabase\\src\\imooc.com\\Google资深工程师深度讲解Go语言\\爬虫\\crawier\\zhenai\\parser\\cityList_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contenes)

	const resultSize = 470
	// 验证得到的数量是不是470，页面上看是470
	if len(result.Request) != resultSize {
		fmt.Errorf("result should hava %d rqusets; but had %d", resultSize, len(result.Request))
	}
	if len(result.Items) != resultSize {
		fmt.Errorf("result should hava %d Items; but had %d", resultSize, len(result.Items))
	}

	// 验证得到的前三个URL能不能对上
	expectedUrl  := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/baicheng1",
		"http://www.zhenai.com/zhenghun/cangzhou",
	}
	for i, url := range expectedUrl {
		if result.Request[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, result.Request[i].Url)
		}
	}
	expectedCity  := []string{
		"阿坝","白城","沧州",
	}
	for i, city := range expectedCity {
		if result.Items[i].(string) != city {
			t.Errorf("expected city #%d: %s; but was %s", i, city, result.Items[i].(string))
		}
	}
}
