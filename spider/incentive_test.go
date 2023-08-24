package spider

import (
	"fmt"
	"testing"
)

func TestParseIncentivePage(t *testing.T) {
	cases := []struct {
		Code    string
		WantMap Smap
	}{
		{"3222", Smap{"優待": "2月・8月"}},
		{"4235", Smap{"優待": ""}},
		{"6200", Smap{"優待": "9月"}},
		{"9267", Smap{"優待": "6月20日"}},
	}
	for _, c := range cases {
		path := fmt.Sprintf("./testdata/incentive/%s.html", c.Code)
		doc := getDocFromPath(path)
		gotMap := parse(doc, incentiveHints)
		deepCheckSmap(t, gotMap, c.WantMap)
	}
}
