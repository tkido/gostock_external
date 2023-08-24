package spider

import (
	"regexp"
	"strings"
)

var detailURLTmpl = "https://finance.yahoo.co.jp/quote/%s"

var re前比 = regexp.MustCompile(`（(.*)%）`)

var detailHints = []Hint{
	{
		"発行",
		`#referenc > div > ul > li:nth-child(2) > dl > dd > span._1fofaCjs._2aohzPlv._1DMRub9m > span > span._3rXWJKZF._11kV6f2G`,
		rmComma,
	},
	{
		"現値",
		`#root > main > div > div > div.XuqDlHPN > div:nth-child(3) > section._1zZriTjI._2l2sDX5w > div._1nb3c4wQ > header > div.nOmR5zWz > span > span > span`,
		rmComma,
	},
	{
		"前終",
		`#detail > section._2Yx3YP9V._3v4W38Hq > div > ul > li:nth-child(1) > dl > dd > span._1fofaCjs._2aohzPlv._1DMRub9m > span > span`,
		rmComma,
	},
	{
		"前比",
		`#root > main > div > div > div.XuqDlHPN > div:nth-child(3) > section._1zZriTjI._2l2sDX5w > div._1nb3c4wQ > div.PRD_bdfF > div._3PynB6qD > div > dl > dd > span > span._1-yujUee.RTJc6XMj._3BXIqAcg > span._3rXWJKZF`,
		pass,
	},
	{
		"出来",
		`#detail > section._2Yx3YP9V._3v4W38Hq > div > ul > li:nth-child(5) > dl > dd > span._1fofaCjs._2aohzPlv._1DMRub9m > span > span._3rXWJKZF._11kV6f2G`,
		rmComma,
	},
	{
		"買残",
		`#margin > div > ul > li:nth-child(1) > dl > dd > span._1fofaCjs._2aohzPlv > span > span._3rXWJKZF`,
		rmComma,
	},
	{
		"売残",
		`#margin > div > ul > li:nth-child(4) > dl > dd > span._1fofaCjs._2aohzPlv > span > span._3rXWJKZF`,
		rmComma,
	},
	{
		"買残週差",
		`#margin > div > ul > li:nth-child(2) > dl > dd > span._1fofaCjs._2aohzPlv > span > span._3rXWJKZF`,
		rmComma,
	},
	{
		"売残週差",
		`#margin > div > ul > li:nth-child(5) > dl > dd > span._1fofaCjs._2aohzPlv > span > span._3rXWJKZF`,
		rmComma,
	},
	{
		"年高",
		`#referenc > div > ul > li:nth-child(11) > dl > dd > span._1fofaCjs._2aohzPlv._1DMRub9m > span > span`,
		rmComma,
	},
	{
		"年安",
		`#referenc > div > ul > li:nth-child(12) > dl > dd > span._1fofaCjs._2aohzPlv._1DMRub9m > span > span`,
		rmComma,
	},
	{
		"年高日",
		`#referenc > div > ul > li:nth-child(11) > dl > dd > span._6wHOvL5`,
		func(s string) string { return strings.Trim(s, "(<!-- -->)\n") },
	},
	{
		"年安日",
		`#referenc > div > ul > li:nth-child(12) > dl > dd > span._6wHOvL5`,
		func(s string) string { return strings.Trim(s, "(<!-- -->)\n") },
	},
	{
		"利",
		`#referenc > div > ul > li:nth-child(3) > dl > dd > span._1fofaCjs._2aohzPlv._1DMRub9m > span > span._3rXWJKZF._11kV6f2G`,
		func(s string) string { return s + "%" },
	},
	{
		"PER",
		`#referenc > div > ul > li:nth-child(5) > dl > dd > a > span._1fofaCjs._2aohzPlv._1DMRub9m > span > span._3rXWJKZF._11kV6f2G`,
		func(s string) string {
			return strings.Trim(rmComma(s), "(連単)- ")
		},
	},
	{
		"PBR",
		`#referenc > div > ul > li:nth-child(6) > dl > dd > a > span._1fofaCjs._2aohzPlv._1DMRub9m > span > span._3rXWJKZF._11kV6f2G`,
		func(s string) string {
			return strings.Trim(rmComma(s), "(連単)- ")
		},
	},
}

var patrolHints = []Hint{
	{
		"発行",
		`#rfindex > div.chartFinance > div:nth-child(2) > dl > dd > strong`,
		rmComma,
	},
	{
		"現値",
		`#stockinf > div.stocksDtl.clearFix > div.forAddPortfolio > table > tbody > tr > td:nth-child(3)`,
		rmComma,
	},
	{
		"年高",
		`#rfindex > div.chartFinance > div:nth-child(11) > dl > dd > strong`,
		rmComma,
	},
	{
		"PER",
		`#rfindex > div.chartFinance > div:nth-child(5) > dl > dd > strong`,
		func(s string) string {
			return strings.Trim(rmComma(s), "(連単)- ")
		},
	},
	{
		"PBR",
		`#rfindex > div.chartFinance > div:nth-child(6) > dl > dd > strong`,
		func(s string) string {
			return strings.Trim(rmComma(s), "(連単)- ")
		},
	},
}
