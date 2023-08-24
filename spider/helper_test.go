package spider

import (
	"log"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func getDocFromPath(path string) *goquery.Document {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func deepCheckSmap(t *testing.T, gotMap, wantMap Smap) {
	for k, v := range wantMap {
		got, ok := gotMap[k]
		if !ok {
			t.Errorf("cannot got value of %v", k)
		} else {
			if got != v {
				t.Errorf("got %v want %v as value of %v", got, v, k)
			}
		}
	}
}
