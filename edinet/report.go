package edinet

import (
	"bytes"
	"log"
	"math"
	"path"
	"sort"
	"strconv"

	"github.com/tkido/gostock/db"
	"github.com/tkido/gostock/my"
)

// Report is Report data from XBRL file
type Report struct {
	EndMonth, FilingDate string
	NetSales, NetIncome,
	BreakupValue, NetCash, Accruals, FreeCashFlow,
	GrossProfitRatio, OperatingProfitRatio, OrdinaryProfitRatio, NetProfitRatio float64
}

// Reports for HTML method
type Reports []Report

// FairValue returns company's fair value
func (rs Reports) FairValue() float64 {
	if len(rs) == 0 {
		return 0
	}
	if len(rs) > 5 {
		rs = rs[:5]
	}
	size := len(rs)
	lastR := rs[0]
	headR := rs[len(rs)-1]
	var gRate float64 // GrowthRate
	if size == 1 || lastR.NetIncome < 0 || headR.NetIncome == 0 {
		gRate = 1
	} else {
		gRate = math.Pow(lastR.NetIncome/headR.NetIncome, 1/float64(size-1))
	}
	stock := math.Min(lastR.BreakupValue, lastR.NetCash)
	flow := lastR.NetIncome
	if flow < 0 {
		return stock + flow*15
	}
	var fgRate float64 //ForcastedGrowthRate
	if gRate < 1 {
		fgRate = gRate // 成長していない場合は成長しない前提で考える
	} else {
		fgRate = gRate/2 + 0.5 // 成長している場合は平均への回帰的に成長減速を織り込む
	}
	var rate float64
	// n年後の割引率を計算する1.1は年間割引率10%を意味する
	pv := func(n int) float64 {
		f := float64(n)
		return math.Pow(fgRate, f) / math.Pow(1.1, f)
	}
	for i := 1; i <= 5; i++ {
		rate += pv(i)
	}
	rate += pv(5) * 10 // 5年後以降は成長しないとして扱う。そこから割引率10%で永久保有≒現在価値は10倍。
	rate = math.Min(rate, 15)

	return stock + flow*rate
}

// HTML make html table
func (rs Reports) HTML() string {
	if len(rs) == 0 {
		return ""
	}
	t := my.NewTable("EDINET有価証券報告書", "numbers")
	t.Th("終了月", "解価", "NetC", "アク", "FCF", "純利", "売上", "粗率", "営率", "経率", "純率", "開示日")
	for _, r := range rs {
		t.Td(
			r.EndMonth,
			r.BreakupValue, r.NetCash, r.Accruals, r.FreeCashFlow,
			r.NetIncome, r.NetSales,
			round(r.GrossProfitRatio),
			round(r.OperatingProfitRatio),
			round(r.OrdinaryProfitRatio),
			round(r.NetProfitRatio),
			r.FilingDate,
		)
	}
	return t.String()
}

func round(f float64) string {
	s := strconv.FormatFloat(f*100, 'f', 0, 64)
	return s + "%"
}

// MakeReports makes Reports from Edinet XBRL files
func MakeReports(code string) (rs Reports, err error) {
	key := path.Join(code, "edinet")
	// 生データをParseしてMapに詰める。重複除去を含む。
	rm := map[string]Report{}
	iter := db.NewIterator(key)
	for iter.Next() {
		r, err := Parse(
			string(iter.Key()),
			bytes.NewBuffer(iter.Value()),
		)
		if err != nil {
			log.Println(err)
			continue
		}
		rm[r.EndMonth] = r
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		return nil, err
	}

	rs = Reports{}
	for _, r := range rm {
		rs = append(rs, r)
	}
	sort.Slice(rs, func(i, j int) bool {
		return rs[i].EndMonth > rs[j].EndMonth
	})
	return
}
