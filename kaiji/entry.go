package kaiji

import (
	"path"
	"sort"

	"github.com/tkido/gostock/db"

	"github.com/tkido/gostock/ufo"
)

func getFeeds(code string) (feeds []*ufo.Feed, err error) {
	ef, err := getFeed("edinet", code)
	if err != nil {
		return nil, err
	}
	tf, err := getFeed("tdnet", code)
	if err != nil {
		return nil, err
	}
	return []*ufo.Feed{ef, tf}, nil
}

func getFeed(tipe, code string) (feed *ufo.Feed, err error) {
	key := path.Join(code, "feed", tipe+".xml")
	if db.Has(key) {
		return ufo.ReadFeed(key)
	}
	return ufo.GetFeed(tipe, code)
}

func collectEntries(feeds []*ufo.Feed) []*ufo.Entry {
	es := make([]*ufo.Entry, 0, 128)
	for _, feed := range feeds {
		for _, entry := range feed.Entries {
			es = append(es, entry)
		}
	}
	sort.SliceStable(es, func(i, j int) bool {
		return es[i].Updated.After(es[j].Updated)
	})
	return es
}
