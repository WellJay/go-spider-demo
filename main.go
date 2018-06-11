package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"golang.org/x/text/transform"
	"io"
	"golang.org/x/net/html/charset"
	"bufio"
	"golang.org/x/text/encoding"
	"regexp"
)

func main() {
	//resp := httpRequestUtf8("http://www.zhenai.com/zhenghun")
	//printCityList(resp)
	printUserByCityUrl("http://www.zhenai.com/zhenghun/aba")
}

func httpRequestUtf8(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code ", resp.StatusCode)
		return nil
	}

	e := determineEncoding(resp.Body)

	//gopm get -g -v golang.org/x/text
	//将其他编码方式统一转成utf8，方便获取数据
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	return all;
}

//gopm get -g -v golang.org/x/net/html
//识别流的编码
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("City：%s，URL：%s\n", m[2], m[1])
	}
	fmt.Printf("found: %d\n", len(matches))
}

//根据城市读取信息
func printUserByCityUrl(cityUrl string) {
	resp := httpRequestUtf8(cityUrl)
	re := regexp.MustCompile(`http://photo.*\.zastatic.com.*\..{3}`)
	matches := re.FindAllSubmatch(resp, -1)
	for _, m := range matches {
		fmt.Printf("headPhoto：%s \n", m)
	}
}