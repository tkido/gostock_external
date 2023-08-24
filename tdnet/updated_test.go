package tdnet

import (
	"testing"

	"github.com/tkido/gostock/my"
)

func TestMakeUpdated(t *testing.T) {
	dls := []Downloaded{
		{"9795", "http://tkido.com/stock/9795.html"},
		{"6200", "http://tkido.com/stock/6200.html"},
		{"3909", "http://tkido.com/stock/6200.html"},
		{"9263", "http://tkido.com/stock/6200.html"},
	}
	html, err := MakeUpdated(dls)
	if err != nil {
		t.Error(err)
	}
	err = my.WriteFile("testdata/updated.html", html)
	if err != nil {
		t.Error(err)
	}
}
