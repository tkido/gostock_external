package kaiji

import (
	"fmt"
	"log"

	"github.com/tkido/gostock/my"
)

// Table return table
func Table(code string) (table string, err error) {
	feeds, err := getFeeds(code)
	if err != nil {
		log.Fatal(err)
	}
	entries := collectEntries(feeds)
	if len(entries) == 0 {
		return "", nil
	}
	t := my.NewTable("開示（EDINETおよびTDnet）")
	t.Th("日付", "内容")

	const format = `<a href="%s" target="_blank" title="%s">%s</a>`
	for _, entry := range entries {
		for _, link := range entry.Links {
			if link.Rel == "alternate" {
				date := makeDate(link.Href)
				title := trimTitle(entry.Title)
				shortTitle := my.TruncateWidth(title, 54)
				if title != shortTitle {
					shortTitle += "…"
				}
				content := fmt.Sprintf(format, link.Href, title, shortTitle)
				t.Td(date, content)
			}
		}
	}
	return t.String(), nil
}
