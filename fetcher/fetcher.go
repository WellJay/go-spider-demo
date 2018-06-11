package fetcher

import (
	"net/http"
	"golang.org/x/text/transform"
	"io/ioutil"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"errors"
	"golang.org/x/text/encoding/unicode"
	"github.com/gpmgo/gopm/modules/log"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("wrong status code")
	}

	e := determineEncoding(resp.Body)

	//gopm get -g -v golang.org/x/text
	//将其他编码方式统一转成utf8，方便获取数据
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

//gopm get -g -v golang.org/x/net/html
//识别流的编码
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Warn("fetcher error：%v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
