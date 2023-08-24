package spider

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

// e.g. https://finance.yahoo.co.jp/quote/6200.T/history?from=20160101&to=20211231&timeFrame=m&page=2
const urlFormat = `https://finance.yahoo.co.jp/quote/%s/history?from=%d0101&to=%d1231&timeFrame=m&page=%d`

func getLonghis(code string) Smap {
	toYear := time.Now().Year()
	fromYear := toYear - 4
	url1 := fmt.Sprintf(urlFormat, code, fromYear, toYear, 1)
	doc1, err := getDocFromURL(url1)
	if err != nil {
		log.Println(err)
		return Smap{}
	}
	days := getDays(doc1)
	url2 := fmt.Sprintf(urlFormat, code, fromYear, toYear, 2)
	doc2, err := getDocFromURL(url2)
	if err != nil {
		log.Println(err)
		return Smap{}
	}
	days2 := getDays(doc2)
	days = append(days, days2...)

	return processLonghis(days)
}

func processLonghis(days Days) Smap {
	smap := Smap{}
	if len(days) == 0 {
		return smap
	}
	const limit = 36 // months 3年来高（安）値 36MH or 36ML
	if len(days) > 36 {
		days = days[:limit]
	}
	days.adjustSplit()

	const format = `=%.0f/【値】`
	sort.SliceStable(days, func(i, j int) bool {
		return days[i].Low <= days[j].Low
	})
	lowest := days[0]
	smap["数年来安値"] = fmt.Sprintf(format, lowest.Low)
	smap["数年来安値月"] = replace(lowest.Date)

	sort.SliceStable(days, func(i, j int) bool {
		return days[i].High >= days[j].High
	})
	highest := days[0]
	smap["数年来高値"] = fmt.Sprintf(format, highest.High)
	smap["数年来高値月"] = replace(highest.Date)

	return smap
}

func replace(s string) string {
	s = strings.Replace(s, "年", "/", 1)
	ss := strings.Split(s, "月")
	return ss[0]
}
