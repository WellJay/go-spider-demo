package parser

import (
	"../../engine"
	"regexp"
)

const userInfoRe = `http://photo.*\.zastatic.com.*\..{3}`

//根据城市读取信息
func ParseUserInfo(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(userInfoRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, m)
		result.Requests = append(result.Requests, engine.Request{
			Url:        "",
			ParserFunc: engine.NilParse,
		})
	}
	return result
}
