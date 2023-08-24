package rss

import (
	"sort"
	"sync"

	"github.com/tkido/gostock/config"
)

// Smap is map[string]string
type Smap map[string]string

// Add other Smap
func (m Smap) Add(other map[string]string) {
	for k, v := range other {
		m[k] = v
	}
}

// Pair is (int, string) pair
type Pair struct {
	row  int
	code string
}

// Pairs is []Pair
type Pairs []Pair

// Map convert Pair to Pair
func (ps Pairs) Map(f func(Pair) Pair) Pairs {
	ch := make(chan Pair)
	q := make(chan struct{}, config.LimitOfParallelProcess)
	wg := sync.WaitGroup{}
	for _, p := range ps {
		wg.Add(1)
		go func(p Pair) {
			q <- struct{}{}
			defer func() { <-q; wg.Done() }()
			ch <- f(p)
		}(p)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	rsts := make(Pairs, 0, len(ps))
	for s := range ch {
		rsts = append(rsts, s)
	}
	sort.Slice(rsts, func(i, j int) bool {
		return rsts[i].row < rsts[j].row
	})
	return rsts
}
