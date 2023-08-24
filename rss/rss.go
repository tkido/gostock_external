package rss

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/tkido/gostock/config"
	"github.com/tkido/gostock/db"
	"github.com/tkido/gostock/my"
	"github.com/tkido/gostock/page"
	"github.com/tkido/gostock/spider"
	"golang.org/x/text/width"
)

func Prepare(codes []string) (err error) {
	wg := sync.WaitGroup{}
	q := make(chan struct{}, config.LimitOfParallelDownload)
	for _, code := range codes {
		wg.Add(1)
		q <- struct{}{}
		go func(code string) {
			defer func() { <-q; wg.Done() }()
			err = prepare(code)
			if err != nil {
				log.Printf("%s: %v", code, err)
			}
		}(code)
	}
	wg.Wait()
	return nil
}

func prepare(code string) (err error) {
	if s, _ := db.GetString(code); s != "" {
		if strings.HasSuffix(s, todayStr) {
			log.Printf("%s is already prepared\n", code)
			return nil
		}
	}

	smap := Smap{}
	smap.Add(fixedData)
	smap.Add(getCommonData(code, now))
	smap.Add(spider.Get(code))

	// html作成
	htmlPath := filepath.Join(config.HTMLRoot, code+".html")
	page.Write(htmlPath, smap)

	smap["短縮名"] = trimName(smap["名称"])

	smap.Add(getUnRealTimeData(smap))
	if s, ok := smap["現値"]; !ok || s == "-" {
		smap["現値"] = " "
	}

	// R欄後半の特色および事業の文字列
	s := fmt.Sprintf(` %s.%s`, smap["特色"], escape(smap["事業"]))
	s = strings.ReplaceAll(s, "、", ",")
	s = strings.ReplaceAll(s, "。", ".")
	s = width.Narrow.String(s)

	smap["R"] = "x" + s
	save(code, "", smap)

	if strings.HasPrefix(smap["市"], "東") {
		// 楽天RSS用データ
		smap.Add(getRssData(code))
		smap["R"] = " " + s
	}
	save(code, "_rss", smap)

	return err
}

func save(code string, key string, smap Smap) (err error) {
	ss := make([]string, len(colOrder))
	for i, col := range colOrder {
		ss[i] = smap[col]
	}
	s := strings.Join(ss, "\t")
	err = db.PutString(code+key, s)
	if err != nil {
		return
	}
	return
}

// Publish is Publish
func Publish(codes []string) (rst string, err error) {
	pairs := make(Pairs, len(codes))
	for i, code := range codes {
		pairs[i] = Pair{i + config.Offset, code}
	}
	pairs = pairs.Map(publish)
	ss := make([]string, len(pairs))
	for i, pair := range pairs {
		ss[i] = pair.code
	}
	return strings.Join(ss, "\n"), nil
}

func publish(p Pair) Pair {
	key := p.code
	if config.RealTime && p.row <= 300 {
		key += "_rss"
	}
	s, err := db.GetString(key)
	if err != nil {
		log.Println(err)
	}
	// replace e.g.【企価】 -> BD3
	p.code = reColumn.ReplaceAllStringFuncSubmatches(
		s,
		func(dollar []string) string {
			return colMap[dollar[1]] + strconv.Itoa(p.row)
		},
	)
	return p
}

// func getEdinetData(code string) (m Smap, err error) {
// 	m = Smap{}
// 	rs, err := edinet.MakeReports(code)
// 	if err != nil {
// 		return
// 	}
// 	m["企価"] = fmt.Sprint(rs.FairValue())
// 	m["edinet"] = rs.HTML()
// 	return
// }

// func getTdnetData(code string) (m Smap, err error) {
// 	m = Smap{}
// 	ps, ds, err := tdnet.MakeReports(code)
// 	if err != nil {
// 		return
// 	}
// 	m["tdnet"] = tdnet.MakeTable(ps, ds)
// 	return
// }

// func getKaijiData(code string) (m Smap, err error) {
// 	m = Smap{}
// 	t, err := kaiji.Table(code)
// 	if err != nil {
// 		return
// 	}
// 	m["kaiji"] = t
// 	return
// }

var reColumn = my.MustCompile(`【(.*?)】`)

// 解説文字列内部の【】がエクセル用の変換に引っかからないようにする
// 【】 -> 《》
func escape(s string) string {
	s = strings.Replace(s, "【", "《", -1)
	s = strings.Replace(s, "】", "》", -1)
	return s
}

func trimName(s string) string {
	if utf8.RuneCountInString(s) > 10 {
		s = width.Narrow.String(s)
	}
	return s
}
