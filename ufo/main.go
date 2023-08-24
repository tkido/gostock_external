package ufo

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/tkido/gostock/db"
	"github.com/tkido/gostock/my"
)

// ReadFeed read ATOM feed from local file.
// used for test from other package mainly.
func ReadFeed(key string) (f *Feed, err error) {
	b, err := db.Get(key)
	if err != nil {
		return
	}
	f = &Feed{}
	err = xml.Unmarshal(b, f)
	return
}

// GetFeed is get feed xml from ufo
func GetFeed(tipe, code string) (f *Feed, err error) {
	const tmpl = "http://resource.ufocatch.com/atom/%s/query/%s"
	url := fmt.Sprintf(tmpl, tipe, code)
	resp, err := my.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	key := path.Join(code, "feed", tipe+".xml")
	err = db.Put(key, b)
	if err != nil {
		return
	}
	f = &Feed{}
	err = xml.Unmarshal(b, f)
	return
}
