package rss

import (
	"fmt"
	"log"
	"time"
)

var fixedData = Smap{
	"値":  `=IF(【現値】=" ", 【前終】, 【現値】)`,
	"値幅": `=VLOOKUP(【基準】, 値幅制限表, 2, TRUE)`,
	"時価": `=【値】*【発行】/100000000`,
	"益":  `=IF(【PER】=0, 0, 1/【PER】`,
	"性":  `=IF(【益】=0, 0, 【利】/【益】`,
	"率":  `=IF(【企価】=0, 0, 【値】/【株価】)`,
	"株価": `=IF(【企価】="", 0, 【企価】/【発行】)`,
	"Vf": `=(【年高】-1)/【Vol】`,
}

func getCommonData(code string, date time.Time) Smap {
	return Smap{
		"ID": code,
		"更新": date.Format("2006/01/02"),
	}
}

// lable, divColum
var urTable = [7][2]string{
	{"出来", "【発行】*1000"},
	{"買残", "【発行】*1000"},
	{"買残週差", "【発行】*1000"},
	{"売残", "【発行】*1000"},
	{"売残週差", "【発行】*1000"},
	{"年高", "【値】"},
	{"年安", "【値】"},
}

func getUnRealTimeData(m Smap) Smap {
	const format = "=%s/%s"
	new := Smap{}
	for _, arr := range urTable {
		s, div := arr[0], arr[1]
		if m[s] == "-" {
			new[s] = "-"
			continue
		}
		new[s] = fmt.Sprintf(format, m[s], div)
	}
	return new
}

// lable, label for RSS, divColum
var rssTable = [][]string{
	{"特売", "特別売気配フラグ"},
	{"最売", "最良売気配値"},
	{"最売数", "最良売気配数量"},
	{"特買", "特別買気配フラグ"},
	{"最買", "最良買気配値"},
	{"最買数", "最良買気配数量"},
	{"前終", "前日終値"},
	{"基準", "当日基準値"},
	{"前比", "前日比率"},
	{"出来", "出来高", "【発行】*1000"},
	{"落日", "配当落日"},
	{"買残", "信用買残", "【発行】*1000"},
	{"買残週差", "信用買残前週比", "【発行】*1000"},
	{"売残", "信用売残", "【発行】*1000"},
	{"売残週差", "信用売残前週比", "【発行】*1000"},
	{"年高", "年初来高値", "【値】"},
	{"年高日", "年初来高値日付"},
	{"年安", "年初来安値", "【値】"},
	{"年安日", "年初来安値日付"},
	{"利", "配当", "【値】"},
	{"PER", "PER"},
	{"PBR", "PBR"},
	{"現値", "現在値"},
}

func getRssData(code string) Smap {
	smap := Smap{}
	for _, arr := range rssTable {
		smap[arr[0]] = rssFormula(code, arr[1:]...)
	}
	return smap
}

func rssFormula(code string, ss ...string) string {
	var id, div string
	switch len(ss) {
	case 1:
		id, div = ss[0], ""
	case 2:
		id, div = ss[0], "/"+ss[1]
	default:
		log.Fatal("rssFormula() invalid args")
	}
	// RSS関数のみの部分
	s := fmt.Sprintf(`RssMarket(%s,"%s")%s`, code, id, div)
	switch id {
	case "配当落日", "年初来高値日付", "年初来安値日付":
		return fmt.Sprintf("=DATEVALUE(%s)", s)
	}
	if len(ss) == 2 {
		return fmt.Sprintf(`=IFERROR(%s, "-")`, s)
	} else {
		return fmt.Sprintf("=%s", s)
	}
}
