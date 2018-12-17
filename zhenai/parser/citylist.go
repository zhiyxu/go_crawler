package parser

import (
	"regexp"

	"github.com/zhiyxu/golearn/project/crawler-concurrent/engine"
)

const cityListRe = `{linkContent:"([^"征婚]+)",linkURL:"(http://m.zhenai.com/zhenghun/[0-9a-z]+)"}`

func ParseCityList(
	contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllStringSubmatch(string(contents), -1)

	// TODO: what's the style?
	result := engine.ParseResult{}
	for _, match := range matches {
		result.Items = append(
			result.Items, "City: "+match[1])
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        match[2],
				ParserFunc: ParseCity,
			})
	}

	return result
}
