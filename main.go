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
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code ", resp.StatusCode)
		return
	}

	e := determineEncoding(resp.Body)

	//gopm get -g -v golang.org/x/net/html
	//将其他编码方式统一转成utf8，方便获取数据
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	printCityList(all)
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
