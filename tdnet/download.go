package tdnet

import (
	"log"
	"path"
	"regexp"
	"sync"

	"github.com/tkido/gostock/config"
	"github.com/tkido/gostock/db"
	"github.com/tkido/gostock/ufo"
)

// Downloaded is Downloaded
type Downloaded struct {
	Code string
	URL  string
}

// Download is Download
func Download(codes []string, check bool) (dls []Downloaded, err error) {
	wg := sync.WaitGroup{}
	q := make(chan struct{}, config.LimitOfParallelDownload)
	ch := make(chan Downloaded)
	go func() {
		for s := range ch {
			dls = append(dls, s)
		}
	}()
	for _, code := range codes {
		wg.Add(1)
		q <- struct{}{}
		go func(code string) {
			defer func() { <-q; wg.Done() }()
			err := download(code, ch, check)
			if err != nil {
				log.Println(err)
			}
		}(code)
	}
	wg.Wait()
	return
}

var (
	reTanshin = regexp.MustCompile(`決算短信`)
	reXbrl    = regexp.MustCompile(`(?:tdnet|tse)-..edjpsm.*?(?:\.xbrl|-ixbrl\.htm)$`)
)

func download(code string, ch chan Downloaded, check bool) (err error) {
	if check {
		updated, err := IsUpdated(code)
		if err != nil {
			return err
		}
		if !updated {
			return nil
		}
	}
	feed, err := ufo.GetFeed("tdnet", code)
	if err != nil {
		return
	}
	dir := path.Join(code, "tdnet")
	for _, entry := range feed.Entries {
		if reTanshin.MatchString(entry.Title) {
			for _, link := range entry.Links {
				if reXbrl.MatchString(link.Href) {
					name := path.Base(link.Href)
					key := path.Join(dir, name)
					if db.Has(key) {
						continue
					}
					err = db.Download(key, link.Href)
					if err != nil {
						log.Println(err)
						continue
					}
					log.Println(link.Href)
					ch <- Downloaded{code, link.Href}
				}
			}
		}
	}
	return nil
}
