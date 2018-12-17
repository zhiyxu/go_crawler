package parser

import (
	"regexp"

	"github.com/zhiyxu/golearn/project/crawler-concurrent/engine"
)

var cityRe = regexp.MustCompile(
	`<a href="(http://m.zhenai.com/u/[\d]+#seo)"[^>]+><div[^>]+><div class="f-nickname"[^>]+>[\s]*([^<\s]+)`)

func ParseCity(contents []byte) engine.ParseResult {

	matches := cityRe.FindAllStringSubmatch(string(contents), -1)

	result := engine.ParseResult{}
	for _, match := range matches {
		result.Items = append(
			result.Items, "User: "+match[2])
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        match[1],
				ParserFunc: ParseProfile,
			})
	}

	return result

}
