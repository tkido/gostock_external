package edinet

import (
	"log"
	"path"
	"regexp"
	"sync"

	"github.com/tkido/gostock/config"
	"github.com/tkido/gostock/db"
	"github.com/tkido/gostock/ufo"
)

// Download はEdinetの有価証券報告書XBRLをダウンロードする
func Download(codes []string, check bool) (err error) {
	wg := sync.WaitGroup{}
	q := make(chan struct{}, config.LimitOfParallelDownload)
	for _, code := range codes {
		wg.Add(1)
		q <- struct{}{}
		go func(code string) {
			defer func() { <-q; wg.Done() }()
			err = download(code, check)
			if err != nil {
				log.Printf("%s: %v", code, err)
			}
		}(code)
	}
	wg.Wait()
	return nil
}

var (
	reYuuhou = regexp.MustCompile(`有価証券報告書`)
	reXbrl   = regexp.MustCompile(`(?:jpcrp|jpfr).*?\.xbrl$`)
)

func download(code string, check bool) (err error) {
	if check {
		updated, err := IsUpdatedIn3Days(code)
		if err != nil {
			return err
		}
		if !updated {
			return nil
		}
	}
	feed, err := ufo.GetFeed("edinet", code)
	if err != nil {
		return
	}
	dir := path.Join(code, "edinet")
	for _, entry := range feed.Entries {
		if reYuuhou.MatchString(entry.Title) {
			for _, link := range entry.Links {
				if reXbrl.MatchString(link.Href) {
					name := path.Base(link.Href)
					key := path.Join(dir, name)
					if db.Has(key) {
						continue
					}
					err = db.Download(key, link.Href)
					if err != nil {
						return
					}
					log.Println(link.Href)
				}
			}
		}
	}
	return
}
