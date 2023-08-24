package spider

import (
	"strings"
)

var incentiveURLTmpl = "https://finance.yahoo.co.jp/quote/%s/incentive"

var incentiveHints = []Hint{
	{
		"優待",
		`#root > main > div:nth-child(2) > div > div.XuqDlHPN > div:nth-child(3) > section._1naUMvAn > div > table > tbody > tr:nth-child(1) > td`,
		func(s string) string { return strings.Replace(s, "末日", "", -1) },
	},
}
