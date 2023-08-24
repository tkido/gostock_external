package patrol

import (
	"log"
	"strconv"
	"sync"

	"github.com/tkido/gostock/config"
	"github.com/tkido/gostock/spider"
	"github.com/tkido/gostock/tdnet"
)

// Patrol find codes that is good enough to give attention
func Patrol(codes []string) []string {
	wg := sync.WaitGroup{}
	q := make(chan struct{}, config.LimitOfParallelProcess)
	ch := make(chan string)
	for _, code := range codes {
		wg.Add(1)
		go func(code string) {
			q <- struct{}{}
			defer func() { <-q; wg.Done() }()
			isGood, err := isGood(code)
			if err != nil {
				log.Printf("%s: %s", code, err)
			}
			if isGood {
				ch <- code
			}
		}(code)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	rsts := []string{}
	for s := range ch {
		rsts = append(rsts, s)
	}
	return rsts
}

var labels = [5]string{"現値", "年高", "PER", "PBR", "発行"}

func parseFls(code string, m map[string]string) []float64 {
	fs := make([]float64, 5)
	for i, l := range labels {
		s, ok := m[l]
		if !ok {
			log.Printf("%s not found", l)
			return fs
		}
		f, err := strconv.ParseFloat(s, 64)
		if err != nil && s != "" && s != "-" {
			log.Printf("%s %s: %s", code, l, err)
		}
		fs[i] = f
	}
	return fs
}

func isGood(code string) (rst bool, err error) {
	log.Printf("patrol %s", code)
	score := 0
	m, err := spider.Patrol(code)
	if err != nil {
		return
	}
	fs := parseFls(code, m)
	// cur, high, per, pbr, out := fs[0], fs[1], fs[2], fs[3], fs[4]
	cur, high, _, _, _ := fs[0], fs[1], fs[2], fs[3], fs[4]

	if 0 < cur && 0 < high {
		if cur >= high*0.96 {
			score += 5
		}
		// } else if cur < high/2 {
		// 	score++
		// }
	}
	// if 0 < per && per <= 5 {
	// 	score++
	// } else if 0 < pbr && pbr <= 0.5 {
	// 	score++
	// }
	// if 0 < cur && 0 < out {
	// 	ers, err := edinet.MakeReports(config.EdinetRoot), code)
	// 	if err != nil {
	// 		return false, err
	// 	}
	// 	ratio := cur * out / ers.FairValue()
	// 	if 0.3 < ratio && ratio < 0.7 {
	// 		score++
	// 	}
	// }

	trs, _, err := tdnet.MakeReports(code)
	if err != nil {
		return
	}
	if len(trs) > 0 {
		if trs[0].NetSales.Float() >= 0.1 {
			score += 2
		}
		if trs[0].NetIncome.Float() >= 0.2 {
			score += 2
		}
	}
	if len(trs) > 4 {
		if trs[4].NetSales.Float() >= 0.1 {
			score++
		}
		if trs[4].NetIncome.Float() >= 0.2 {
			score++
		}
	}

	if score >= 9 {
		rst = true
	}
	return
}
