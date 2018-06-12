package fetcher

import (
	"bufio"
	"errors"
	"github.com/gpmgo/gopm/modules/log"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
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

	bodyReader := bufio.NewReader(resp.Body)

	e := determineEncoding(bodyReader)

	//gopm get -g -v golang.org/x/text
	//将其他编码方式统一转成utf8，方便获取数据
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

//gopm get -g -v golang.org/x/net/html
//识别流的编码
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Warn("fetcher error：%v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
