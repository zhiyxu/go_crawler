package parser

import (
	"regexp"
	"strconv"

	"github.com/zhiyxu/golearn/project/crawler-concurrent/model"

	"github.com/zhiyxu/golearn/project/crawler-concurrent/engine"
)

//<div class="des f-cl" data-v-3c42fade>兰州 | 34岁 | 大专 | 未婚 | 164cm | 3001-5000元</div>
//var IDRe = regexp.MustCompile(`<div[^>]*>ID：([\d]+)</div>`)
//var GenderRe = regexp.MustCompile(`"genderString":"(.+)士"`)
//var LocationRe = regexp.MustCompile(
//	`<div[^>]*>([^|]*) | [\d]*岁 | [^|]* | [^|]* | [\d]*cm | [^<]*</div>`)
//var AgeRe = regexp.MustCompile(
//	`<div[^>]*>[^|]* | ([\d]*)岁 | [^|]* | [^|]* | [\d]*cm | [^<]*</div>`)
//var EducationRe = regexp.MustCompile(
//	`<div[^>]*>[^|]* | [\d]*岁 | ([^|]*) | [^|]* | [\d]*cm | [^<]*</div>`)
//var MarriageRe = regexp.MustCompile(
//	`<div[^>]*>[^|]* | [\d]*岁 | [^|]* | ([^|]*) | [\d]*cm | [^<]*</div>`)
//var HeightRe = regexp.MustCompile(
//	`<div[^>]*>[^|]* | [\d]*岁 | [^|]* | [^|]* | ([\d]*)cm | [^<]*</div>`)
//var IncomeRe = regexp.MustCompile(
//	`<div[^>]*>[^|]* | [\d]*岁 | [^|]* | [^|]* | [\d]*cm | ([^<]*)</div>`)

var IDRe = regexp.MustCompile(`<p[^>]+>ID:([\d]+)</p>`)
var GenderRe = regexp.MustCompile(`"genderString":"(.+)士"`)
var LocationRe = regexp.MustCompile(
	`个人资料[^择偶条件]*<span class="nick_name">工作地</span>[\s]*<span>([^<]+)</span>`)
var AgeRe = regexp.MustCompile(
	`个人资料[^择偶条件]*<span class="nick_name">年龄</span>[\s]*<span>([\d]*)岁</span>`)
var EducationRe = regexp.MustCompile(
	`个人资料[^择偶条件]*<span class="nick_name">学历</span>[\s]*<span>([^<]+)</span>`)
var MarriageRe = regexp.MustCompile(
	`个人资料[^择偶条件]*<span class="nick_name">婚姻状况</span>[\s]*<span>([^<]+)</span>`)
var HeightRe = regexp.MustCompile(
	`个人资料[^择偶条件]*<span class="nick_name">身高</span>[\s]*<span>([\d]*)CM</span>`)
var IncomeRe = regexp.MustCompile(
	`个人资料[^择偶条件]*<span class="nick_name">月收入</span>[\s]*<span>([^<]+)</span>`)

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}

	profile.ID = extractInt(contents, IDRe)
	profile.Gender = extractString(contents, GenderRe)
	profile.Location = extractString(contents, LocationRe)
	profile.Age = extractInt(contents, AgeRe)
	profile.Education = extractString(contents, EducationRe)
	profile.Marriage = extractString(contents, MarriageRe)
	profile.Height = extractInt(contents, HeightRe)
	profile.Income = extractString(contents, IncomeRe)

	result := engine.ParseResult{
		//TODO: what's the usage
		//TODO: no more requests
		Items: []interface{}{profile},
	}

	return result
}

func extractInt(contents []byte, re *regexp.Regexp) int {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		content, err := strconv.Atoi(string(match[1]))
		if err == nil {
			return content
		}
	}

	return -1
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
