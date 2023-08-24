package dlinfo

import (
	"log"
	"sync"

	"github.com/tkido/gostock/config"
	"github.com/tkido/gostock/edinet"
	"github.com/tkido/gostock/my/json"
)

// Pair is Pair
type Pair struct {
	Code   string
	Active bool
}

// UpdateActiveMap update active map
func UpdateActiveMap(dlInfo, activeMap string) (err error) {
	rs, err := ParseDlInfo(dlInfo)
	if err != nil {
		return
	}
	m := map[string]bool{}
	err = json.Load(activeMap, &m)
	if err != nil {
		return
	}

	wg := sync.WaitGroup{}
	q := make(chan struct{}, config.LimitOfParallelProcess)
	ch := make(chan Pair)
	for _, r := range rs {
		if r.StockCode == "" || r.EdinetCode == "" {
			continue
		}
		code := r.StockCode[:4]
		a, ok := m[code]
		if ok && (a == r.Listing) {
			continue
		}
		wg.Add(1)
		go func(code string) {
			q <- struct{}{}
			defer func() { <-q; wg.Done() }()
			log.Printf("check: %s", code)
			isActive, err := edinet.IsUpdatedIn6Months(code)
			if err != nil {
				log.Printf("%s: %v", code, err)
			}
			ch <- Pair{code, isActive}
		}(code)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for p := range ch {
		m[p.Code] = p.Active
	}
	err = json.Save(activeMap, m)
	if err != nil {
		return
	}
	return
}
