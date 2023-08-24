// Package tdnet はtdnetの更新を感知する。
package tdnet

import (
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var (
	t0, t1 string
	client = &http.Client{}
)

const tdnetURL = "https://www.release.tdnet.info/onsf/TDJFSearch/TDJFSearch"

func init() {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	const format = "20060102"
	t0 = yesterday.Format(format)
	t1 = now.Format(format)
}

// IsUpdated は一両日中にtdnetで開示があったかどうかを返す
func IsUpdated(code string) (bool, error) {
	values := url.Values{}
	values.Add("m", "0")      // fixed
	values.Add("q", code+"0") // query
	values.Add("t0", t0)      // start date
	values.Add("t1", t1)      // end date

	resp, err := client.PostForm(tdnetURL, values)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return false, err
	}
	s := doc.Find("#result").Text()

	return s != "", nil
}

// curl 'https://www.release.tdnet.info/onsf/TDJFSearch/TDJFSearch' -d 't0=20180121&t1=20180219&q=42350&m=0'
// のようなURLからtdnetの更新を感知する。
// 実際のデータはufo catcher web API から取る。
