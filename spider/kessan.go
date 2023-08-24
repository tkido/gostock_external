package spider

var kessanURLTmpl = "https://www.nikkei.com/markets/kigyo/money-schedule/kessan/?ResultFlag=3&kwd=%s"

var kessanHints = []Hint{
	{
		"決算",
		`#newpresSchedule > div.m-block.newpresSearchResults > div > div.m-artcle > table > tbody > tr > th`,
		pass,
	},
}
