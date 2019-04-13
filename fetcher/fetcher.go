package fetcher

import (
	"net/http"
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"fmt"
	"log"
	"golang.org/x/text/encoding/unicode"
	"time"
)

var rateLimiter = time.Tick(10 * time.Millisecond)
// 获取页面body信息
func Fetch(url string) ([]byte, error)  {
	<- rateLimiter
	client := &http.Client{}
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	//增加header选项
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Safari/537.36")
	//reqest.Header.Add("Referer", "http://album.zhenai.com/u/1486293757")
	//reqest.Header.Add("Host", "www.zhenai.com")
	//reqest.Header.Add("Cookie", "flashRegisterSwitch=1; sid=hZOzHECElKiJniji18sW; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1554303290,1554547384,1554629216,1554734717; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1554735190; __channelId=905821%2C0")

	if err != nil {
		panic(err)
	}
	//处理返回结果
	resp, err := client.Do(reqest)



	//resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 判断返回码是否为200
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	// 如果编码格式不为UTF-8,进行转换
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)

}

// 用charset.DetermineEncoding判断编码格式，改写
func determineEncoding(r *bufio.Reader) encoding.Encoding  { // 对charset.DetermineEncoding进行包装
	bytes, err := r.Peek(1024) // 截取前1024字节
	if err != nil{
		log.Printf("Fetcher error : %v", err)
		return 	unicode.UTF8 // 遇到问题返回默认值UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}