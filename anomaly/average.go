package anomaly

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/tkido/gostock/config"
	"github.com/tkido/gostock/my/json"
	"github.com/tkido/gostock/spider"
)

func Culculation() (err error) {
	data := map[string][]float64{}
	closes := map[int]float64{2009: 10546.44}
	days := []Day{}
	err = json.Load("./testdata/avarage.json", &days)
	for _, day := range days {
		t := day.Time
		closes[t.Year()] = day.Value
		value := day.Value / closes[t.Year()-1]
		// fmt.Println(t)
		// fmt.Println(value)
		label := fmt.Sprintf("%02d/%02d", t.Month(), t.Day())
		if _, ok := data[label]; !ok {
			data[label] = []float64{}
		}
		data[label] = append(data[label], value)
	}
	// fmt.Println(data)
	rsts := map[string]float64{}
	for k, v := range data {
		sum := 0.0
		for _, f := range v {
			sum += f
		}
		rst := sum / float64(len(v))
		fmt.Printf("%s\t%d\t%f\n", k, len(v), rst)
		rsts[k] = rst
	}
	for k, v := range rsts {
		fmt.Printf("%s\t%f\n", k, v)
	}
	return
}

type Day struct {
	Time  time.Time
	Value float64
}

func Avarage() (err error) {
	wg := sync.WaitGroup{}
	q := make(chan struct{}, config.LimitOfParallelProcess)
	ch := make(chan Day)
	for i := 1; i <= 13; i++ {
		wg.Add(1)
		go func(i int) {
			q <- struct{}{}
			defer func() { <-q; wg.Done() }()
			err = Download(i, ch)
			if err != nil {
				return
			}
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	days := []Day{}
	err = json.Load("./testdata/avarage.json", &days)
	if err != nil {
		return err
	}
	for day := range ch {
		days = append(days, day)
	}
	sort.SliceStable(days, func(i, j int) bool {
		return days[i].Time.Before(days[j].Time)
	})
	err = json.Save("./testdata/avarage.json", days)
	if err != nil {
		return
	}
	return
}

func Download(index int, ch chan Day) (err error) {
	const tmpl = `https://info.finance.yahoo.co.jp/history/?code=998407.O&sy=2019&sm=1&sd=1&ey=2020&em=12&ed=31&tm=d&p=%d`
	url := fmt.Sprintf(tmpl, index)
	doc, err := spider.GetDocFromURL(url)
	if err != nil {
		return
	}
	tds := doc.Find(`#main > div.padT12.marB10.clearFix > table > tbody > tr > td`)

	const format = "2006年1月2日"
	var day Day
	tds.Each(func(i int, td *goquery.Selection) {
		mod := i % 5
		switch mod {
		case 0:
			day = Day{}
			t, err := time.Parse(format, td.Text())
			if err != nil {
				return
			}
			day.Time = t
		case 4:
			s := strings.Replace(td.Text(), ",", "", -1)
			v, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return
			}
			day.Value = v
			ch <- day
		default:
			// pass
		}
	})
	return nil
}
