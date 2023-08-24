package ufo

import "time"

// Feed is Atom Feed
type Feed struct {
	ID    string `xml:"id"`
	Title string `xml:"title"`
	Link  Link   `xml:"link"`
	// Updated time.Time `xml:"updated"`
	Entries []*Entry `xml:"entry"`
}

// Updated Feed自体のxml:"updated"は呼んだ時に生成されてしまうようなので最新Entryのそれを使用する。
func (f *Feed) Updated() time.Time {
	return f.Entries[0].Updated
}

// Entry is Entry
type Entry struct {
	Title   string    `xml:"title"`
	Links   []Link    `xml:"link"`
	Updated time.Time `xml:"updated"`
}

// Link is Link
type Link struct {
	Type string `xml:"type,attr"`
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
}
