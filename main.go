package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/tkido/gostock/db"
	"github.com/tkido/gostock/edinet"
	"github.com/tkido/gostock/edinet/dlinfo"
	"github.com/tkido/gostock/my"
	"github.com/tkido/gostock/my/json"
	"github.com/tkido/gostock/patrol"
	"github.com/tkido/gostock/rss"
)

var reCode = regexp.MustCompile(`^\d{4}$`)

const (
	pConfig     = "./data/conf.properties"
	pActiveList = "./data/activelist.txt"
	pActiveMap  = "./data/activemap.json"
	pDlInfo     = "./data/EdinetcodeDlInfo.csv"
	pEdinetCode = "./data/edinetcodemap.json"
	// pRireki     = "./data/rireki.txt"
	pRssTable = "./data/table.txt"

	// pDownloaded = "./data/downloaded.html"
	pRssResult = "./data/result.html"
)

func main() {
	log.Println("START")
	defer log.Println("END")
	defer db.Close()

	var target string
	flag.StringVar(&target, "t", "rss", "target task")
	flag.Parse()
	prepare()
	clean()
	err := run(target)
	if err != nil {
		log.Fatal(err)
	}
}

func clean() {
	os.Remove(pRssResult)
}

func prepare() {
	m := map[string]string{}
	err := json.Load(pEdinetCode, &m)
	if err != nil {
		log.Fatal(err)
	}
	edinet.EdinetCodeMap = m
}

func run(target string) (err error) {
	switch target {
	case "build":
		err = build()
	case "patrol":
		err = doPatrol()
	case "foo":
		err = doFoo()
	case "prepare":
		err = doPrepare()
	case "publish":
		err = doPublish()
	default:
		return fmt.Errorf("invalid target %s", target)
	}
	if err != nil {
		return
	}
	return
}

func doFoo() (err error) {
	fmt.Println("hogehoge!")
	return nil
}

func doPrepare() (err error) {
	codes, err := my.ReadlinesMatched(pRssTable, reCode)
	if err != nil {
		return
	}
	err = rss.Prepare(codes)
	return
}

func doPublish() (err error) {
	codes, err := my.ReadlinesMatched(pRssTable, reCode)
	if err != nil {
		return
	}
	rst, err := rss.Publish(codes)
	if err != nil {
		return
	}
	err = my.WriteFileForCopyPaste(pRssResult, rst)
	if err != nil {
		return
	}
	return
}

func doPatrol() (err error) {
	codes, err := my.ReadlinesMatched(pActiveList, reCode)
	if err != nil {
		return
	}
	goods := patrol.Patrol(codes)
	if len(goods) == 0 {
		return nil
	}
	// err = rireki.Parse(pRireki)
	// if err != nil {
	// 	return
	// }
	// rst, err := rss.Execute(goods)
	// if err != nil {
	// 	return
	// }
	// err = my.WriteFileForCopyPaste(pRssResult, rst)
	// if err != nil {
	// 	return
	// }
	return
}

func build() (err error) {
	err = dlinfo.MakeCodeMap(pDlInfo, pEdinetCode)
	if err != nil {
		return
	}
	prepare()
	err = dlinfo.UpdateActiveMap(pDlInfo, pActiveMap)
	if err != nil {
		return
	}
	err = dlinfo.MakeActiveList(pActiveMap, pActiveList)
	if err != nil {
		return
	}
	return
}
