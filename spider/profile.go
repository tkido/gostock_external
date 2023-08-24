package spider

import (
	"regexp"
	"strings"

	"golang.org/x/text/unicode/norm"
)

var profileURLTmpl = "https://finance.yahoo.co.jp/quote/%s/profile"

var profileHints = []Hint{
	{
		"名称",
		`#root > main > div > div > div.XuqDlHPN > div:nth-child(3) > section._1zZriTjI._2l2sDX5w > div._1nb3c4wQ > header > div.DL5lxuTC > h1`,
		trimName,
	},
	{
		"特色",
		`#profile > div > table > tbody > tr:nth-child(1) > td > p`,
		func(s string) string {
			if len(s) > 12 {
				return s[12:]
			}
			return s
		},
	},
	{
		"事業",
		`#profile > div > table > tbody > tr:nth-child(2) > td > p`,
		func(s string) string {
			if len(s) > 18 {
				return s[18:]
			}
			return s
		},
	},
	{
		"分類",
		`#profile > div > table > tbody > tr:nth-child(5) > td > a`,
		func(s string) string {
			if len(s) <= 2 {
				return s
			}
			s = strings.Replace(s, "・", "", 1)
			if strings.HasPrefix(s, "その他") {
				return strings.TrimPrefix(s, "その他")
			}
			s = strings.TrimSuffix(s, "製品")
			return strings.TrimSuffix(s, "業")
		},
	},
	{
		"設立",
		`#profile > div > table > tbody > tr:nth-child(8) > td > p`,
		func(s string) string {
			if len(s) > 4 {
				return s[:4]
			}
			return s
		},
	},
	{
		"上場",
		`#profile > div > table > tbody > tr:nth-child(10) > td > p`,
		func(s string) string {
			if len(s) > 4 {
				return s[:4]
			}
			return s
		},
	},
	{
		"決期",
		`#profile > div > table > tbody > tr:nth-child(11) > td > p`,
		func(s string) string { return strings.TrimSuffix(s, "末日") },
	},
	{
		"従連",
		`#profile > div > table > tbody > tr:nth-child(14) > td > p`,
		func(s string) string { return strings.TrimSuffix(s, "人") },
	},
	{
		"従単",
		`#profile > div > table > tbody > tr:nth-child(13) > td > p`,
		func(s string) string { return strings.TrimSuffix(s, "人") },
	},
	{
		"齢",
		`#profile > div > table > tbody > tr:nth-child(15) > td > p`,
		func(s string) string { return strings.TrimRight(s, "-歳") },
	},
	{
		"収",
		`#profile > div > table > tbody > tr:nth-child(16) > td > p`,
		func(s string) string {
			s = rmComma(strings.TrimRight(s, "-千円"))
			if s != "" {
				s = s[:len(s)-1]
			}
			return s
		},
	},
	{
		"代表",
		`#profile > div > table > tbody > tr:nth-child(7) > td > p`,
		func(s string) string {
			s = strings.Replace(s, "　", "", 1)
			s = strings.Replace(s, " [役員]", "", 1)
			return s
		},
	},
	{
		"市",
		`#profile > div > table > tbody > tr:nth-child(9) > td > p`,
		marketName,
	},
	{
		"市記号",
		`#profile > div > table > tbody > tr:nth-child(9) > td > p`,
		marketCode,
	},
}

var reTrimName = regexp.MustCompile(`\(株\)|・|　|＆|ホールディングス?|コーポレーション|カンパニー|グループ|本社|ジャパン$`)

func trimName(s string) string {
	s = reTrimName.ReplaceAllString(s, "")
	s = norm.NFKC.String(s)
	return s
}

func marketName(s string) string {
	ss := strings.Split(s, ",")
	if len(ss) > 1 {
		s = ss[0]
	}
	switch s {
	case "東証プライム":
		return "東P"
	case "東証スタンダード":
		return "東S"
	case "東証グロース":
		return "東G"
	case "名証プレミア":
		return "名P"
	case "名証メイン":
		return "名M"
	case "名証ネクスト":
		return "名N"
	case "札証":
		return "札"
	case "札証アンビシャス":
		return "札A"
	case "福証":
		return "福"
	case "福証Q-Board":
		return "福Q"
	default:
		return s
	}
}

func marketCode(s string) string {
	ss := strings.Split(s, ",")
	if len(ss) > 1 {
		s = ss[0]
	}
	switch {
	case strings.HasPrefix(s, "福"):
		return "F"
	case strings.HasPrefix(s, "札"):
		return "S"
	case strings.HasPrefix(s, "名"):
		return "N"
	default:
		return "T"
	}
}
