package spider

import (
	"fmt"
	"testing"
)

func TestParseHistory(t *testing.T) {
	cases := []struct {
		Code    string
		WantMap Smap
	}{
		{"2884", Smap{"日出": "=18100/【発行】*1000", "週出": "=15820/【発行】*1000", "月出": "=42605/【発行】*1000", "Vol": "0.033543", "SPR": "1.355172", "RCI": "-0.828571", "乖離": "=1-491.900000/【値】"}},
		{"5285", Smap{"日出": "=9700/【発行】*1000", "週出": "=31380/【発行】*1000", "月出": "=21310/【発行】*1000", "Vol": "0.065422", "SPR": "0.804604", "RCI": "-0.060150", "乖離": "=1-326.725000/【値】"}},
	}
	for _, c := range cases {
		path := fmt.Sprintf("./testdata/history/%s.html", c.Code)
		doc := getDocFromPath(path)
		gotMap := parseHistory(doc)
		deepCheckSmap(t, gotMap, c.WantMap)
	}
}
