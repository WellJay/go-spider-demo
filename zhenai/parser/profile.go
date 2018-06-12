package parser

import (
	"regexp"
	"spider/engine"
	"spider/model"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)

//读取用户资料
func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}
	profile.Marriage = extractString(contents, marriageRe)

	return engine.ParseResult{
		Items: []interface{}{profile},
	}
}

func extractString(contexts []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contexts)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
