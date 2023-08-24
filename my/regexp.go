package my

import (
	"bytes"
	"regexp"
)

// Regexp can use comment and ReplaceAllStringFuncSubmatches
type Regexp struct {
	*regexp.Regexp
}

// comment to delete
var reReComment = regexp.MustCompile(`(?m)(\s+)|(\#.*$)`)

// MustCompile can use comment
func MustCompile(raw string) *Regexp {
	s := reReComment.ReplaceAllString(raw, "")
	return &Regexp{regexp.MustCompile(s)}
}

// ReplaceAllStringFuncSubmatches 後方参照を使う関数でReplaceAllするdollar[i]が$iを表す
func (re *Regexp) ReplaceAllStringFuncSubmatches(src string, repl func([]string) string) string {
	buf := bytes.Buffer{}
	cursor := 0
	sms := re.FindAllStringSubmatchIndex(src, -1)
	for _, sm := range sms {
		if sm[0] == -1 {
			continue
		}
		buf.WriteString(src[cursor:sm[0]])
		cursor = sm[1]
		dollar := make([]string, 0, len(sm))
		for i := 0; i < len(sm)/2; i++ {
			if sm[2*i] != -1 {
				dollar = append(dollar, src[sm[2*i]:sm[2*i+1]])
			} else {
				dollar = append(dollar, "")
			}
		}
		buf.WriteString(repl(dollar))
	}
	buf.WriteString(src[cursor:len(src)])
	return buf.String()
}
