package spider

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/tkido/gostock/my"
)

func appendData(data []Smap, code, tmpl string, hints []Hint) []Smap {
	url := fmt.Sprintf(tmpl, code)
	doc, err := getDocFromURL(url)
	if err != nil {
		log.Println(url)
		log.Println(err)
	} else {
		data = append(data, parse(doc, hints))
	}
	return data
}

// Patrol gets data for patrolfrom Web
func Patrol(code string) (Smap, error) {
	url := fmt.Sprintf(detailURLTmpl, code)
	doc, err := getDocFromURL(url)
	if err != nil {
		return nil, err
	}
	return parse(doc, patrolHints), nil
}

// Get is get data from Web
func Get(code string) Smap {
	smap := map[string]string{}
	targets := []Smap{}
	targets = appendData(targets, code, detailURLTmpl, detailHints)
	// targets = appendData(targets, code, consolidateURLTmpl, consolidateHints)
	targets = appendData(targets, code, profileURLTmpl, profileHints)
	targets = appendData(targets, code, kessanURLTmpl, kessanHints)
	targets = appendData(targets, code, incentiveURLTmpl, incentiveHints)
	targets = append(targets, getLonghis(code))

	doc, err := getDocFromURL(fmt.Sprintf(historyURLTmpl, code))
	if err != nil {
		log.Println(err)
	} else {
		targets = append(targets, parseHistory(doc))
	}

	for _, target := range targets {
		for k, v := range target {
			smap[k] = v
		}
	}
	return smap
}

// Incentive gets incentive data for from Web
func Incentive(code string) (Smap, error) {
	url := fmt.Sprintf(incentiveURLTmpl, code)
	doc, err := getDocFromURL(url)
	if err != nil {
		return nil, err
	}
	return parse(doc, incentiveHints), nil
}

func GetDocFromURL(url string) (doc *goquery.Document, err error) {
	return getDocFromURL(url)
}

func getDocFromURL(url string) (doc *goquery.Document, err error) {
	resp, err := my.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	doc, err = goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return
	}
	return
}

func parse(doc *goquery.Document, hints []Hint) Smap {
	smap := Smap{}
	for _, hint := range hints {
		s := doc.Find(hint.selector).Text()
		if hint.sanitize != nil {
			s = hint.sanitize(s)
		}
		smap[hint.label] = s
	}
	return smap
}
