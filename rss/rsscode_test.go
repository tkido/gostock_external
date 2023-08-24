package rss

import (
	"reflect"
	"testing"
	"time"
)

func TestGetCommonData(t *testing.T) {
	date := time.Date(2018, 2, 8, 0, 0, 0, 0, time.Local)
	cases := []struct {
		Code string
		Want Smap
	}{
		{"6200", Smap{"ID": "6200", "更新": "2018/02/08"}},
		{"3085", Smap{"ID": "3085", "更新": "2018/02/08"}},
	}
	for _, c := range cases {
		got := getCommonData(c.Code, date)
		if !reflect.DeepEqual(got, c.Want) {
			t.Errorf("got %v want %v", got, c.Want)
		}
	}
}
func TestRssFormula(t *testing.T) {
	cases := []struct {
		Code string
		Args []string
		Want string
	}{
		{"6200", []string{"現在値"}, `=RssMarket(6200,"現在値")`},
		{"3085", []string{"信用買残", "【発行】*1000"}, `=IFERROR(RssMarket(3085,"信用買残")/【発行】*1000, "-")`},
	}
	for _, c := range cases {
		got := rssFormula(c.Code, c.Args...)
		want := c.Want
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
