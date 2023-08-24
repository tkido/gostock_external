package tdnet

import (
	"bytes"
	"log"
	"path"
	"sort"
	"strconv"

	"github.com/tkido/gostock/db"
	"github.com/tkido/gostock/my"
)

// Doukihi is 前年同期比
type Doukihi struct {
	current, previous float64
}

// Float is Float
func (d Doukihi) Float() float64 {
	if d.previous <= 0 {
		return -1
	}
	return d.current/d.previous - 1
}

func (d Doukihi) String() string {
	// 前年黒字
	if d.previous > 0 {
		if d.current < 0 {
			return `-赤転`
		}
		// 通常比率
		f := d.current/d.previous*100 - 100
		s := strconv.FormatFloat(f, 'f', 0, 64)
		return s + "%"
	}
	// 前年赤字
	if d.current < 0 {
		if d.current < d.previous {
			return `-赤拡`
		} else if d.current > d.previous {
			return `-赤縮`
		}
		return `-赤字`
	}
	return `+黒転`
}

// Report is Report data from XBRL file
type Report struct {
	Year, Quater                                         int
	FilingDate, EndMonth                                 string
	NetSales, OperatingIncome, OrdinaryIncome, NetIncome float64
}

func (r Report) id() int {
	return r.Year*10 + r.Quater
}

// 差分を取るべき生ReportのID
func (r Report) prevID() (id int, ok bool) {
	if r.Quater == 1 {
		return 0, false //1Qの場合はない
	}
	return r.Year*10 + r.Quater - 1, true
}

// 前年同期の差分ReportのID
func (r Report) lastID() int {
	return (r.Year-1)*10 + r.Quater
}

// Reports for HTML method
type Reports []Report

// PerReport は前年同期比のReport
type PerReport struct {
	Year, Quater                                         int
	FilingDate, EndMonth                                 string
	NetSales, OperatingIncome, OrdinaryIncome, NetIncome Doukihi
}

func (p PerReport) id() int {
	return p.Year*10 + p.Quater
}

// PerReports for HTML method
type PerReports []PerReport

// MakeReports makes report from XBRL files from TDnet
func MakeReports(code string) (ps PerReports, ds Reports, err error) {
	key := path.Join(code, "tdnet")
	rm := map[int]Report{}
	// 生データをParseしてRawMapに詰める
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
		rm[r.id()] = r
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		return nil, nil, err
	}
	// 前Qとの差分Deltaを計算してMapに詰める
	dm := map[int]Report{}
	for id, r := range rm {
		prevID, ok := r.prevID()
		if !ok {
			dm[id] = r
			continue
		}
		prev, ok := rm[prevID]
		if !ok {
			continue
		}
		delta := Report{
			r.Year,
			r.Quater,
			r.FilingDate,
			r.EndMonth,
			r.NetSales - prev.NetSales,
			r.OperatingIncome - prev.OperatingIncome,
			r.OrdinaryIncome - prev.OrdinaryIncome,
			r.NetIncome - prev.NetIncome,
		}
		dm[id] = delta
	}
	// 前年同期比Percentageを計算
	for _, d := range dm {
		ds = append(ds, d)
		last, ok := dm[d.lastID()]
		if !ok {
			continue
		}
		per := PerReport{
			d.Year,
			d.Quater,
			d.FilingDate,
			d.EndMonth,
			Doukihi{d.NetSales, last.NetSales},
			Doukihi{d.OperatingIncome, last.OperatingIncome},
			Doukihi{d.OrdinaryIncome, last.OrdinaryIncome},
			Doukihi{d.NetIncome, last.NetIncome},
		}
		ps = append(ps, per)
	}
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].id() > ds[j].id()
	})
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].id() > ps[j].id()
	})
	return
}

// MakeTable makes html table from reports
func MakeTable(ps PerReports, ds Reports) (table string) {
	if len(ds) == 0 {
		return
	}
	t := my.NewTable("TDnet短信（金額および前年同期比）", "numbers")
	t.Th("終了月", "Q", "売上", "営利", "経利", "純利", "売上", "営利", "経利", "純利", "開示日")
	m := map[int]PerReport{}
	for _, p := range ps {
		m[p.id()] = p
	}
	for _, d := range ds {
		if p, ok := m[d.id()]; ok {
			t.Td(d.EndMonth, d.Quater, d.NetSales, d.OperatingIncome, d.OrdinaryIncome, d.NetIncome, p.NetSales, p.OperatingIncome, p.OrdinaryIncome, p.NetIncome, d.FilingDate)
			continue
		}
		t.Td(d.EndMonth, d.Quater, d.NetSales, d.OperatingIncome, d.OrdinaryIncome, d.NetIncome, "", "", "", "", d.FilingDate)
	}
	return t.String()
}
