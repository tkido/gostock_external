package rss

import "time"

var now = time.Now()
var todayStr = now.Format("2006/01/02")

var colOrder = []string{
	"ID", "短縮名", "R", "値",
	"特売", "最売", "最売数", "特買", "最買", "最買数",
	"現値", "前終", "基準", "値幅", "前比",
	"出来", "日出", "週出", "月出",
	"買残", "買残週差", "売残", "売残週差",
	"年高", "年高日", "年安", "年安日", "乖離",
	"Vol", "Vf", "SPR", "RCI",
	"数年来高値", "数年来高値月", "数年来安値", "数年来安値月",
	"利", "益", "性", "ROE", "自",
	"PER", "PBR", "市", "時価",
	"決算", "優待", "落日", "発行", "分類",
	"代表", "設立", "上場", "決期",
	"従連", "従単", "齢", "収",
	"率", "株価", "企価", "更新",
}

var colMap = map[string]string{}

func init() {
	cols := []string{}
	for _, s := range []string{"", "A", "B"} {
		for _, rune := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
			cols = append(cols, s+string(rune))
		}
	}
	for i, label := range colOrder {
		colMap[label] = cols[i]
	}
}
