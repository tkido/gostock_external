package spider

import (
	"fmt"
	"testing"
)

func TestParseKessanPage(t *testing.T) {
	cases := []struct {
		Code    string
		WantMap Smap
	}{
		{"6200", Smap{"決算": "2022/7/25"}},
	}
	for _, c := range cases {
		path := fmt.Sprintf("./testdata/kessan/%s.html", c.Code)
		doc := getDocFromPath(path)
		gotMap := parse(doc, kessanHints)
		deepCheckSmap(t, gotMap, c.WantMap)
	}
}
