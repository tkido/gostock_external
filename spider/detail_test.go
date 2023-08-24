package spider

import (
	"fmt"
	"testing"
)

func TestParseDetailPage(t *testing.T) {
	cases := []struct {
		Code    string
		WantMap Smap
	}{
		{"4235",
			Smap{
				"発行":   "7830600",
				"現値":   "4610",
				"前終":   "4715",
				"前比":   "-2.23",
				"出来":   "74600",
				"買残":   "136100",
				"売残":   "0",
				"買残週差": "+9100",
				"売残週差": "0",
				"年高":   "5950",
				"年安":   "1788",
				"年高日":  "22/09/06",
				"年安日":  "22/01/28",
				"利":    "0.74%",
				"PER":  "20.23",
				"PBR":  "2.49",
			},
		},
		{"6200",
			Smap{
				"発行":   "42621500",
				"現値":   "2595",
				"前終":   "2564",
				"前比":   "+1.21",
				"出来":   "322100",
				"買残":   "90400",
				"売残":   "214700",
				"買残週差": "+300",
				"売残週差": "-17700",
				"年高":   "2991",
				"年安":   "1638",
				"年高日":  "22/08/16",
				"年安日":  "22/01/28",
				"利":    "0.83%",
				"PER":  "49.66",
				"PBR":  "19.71",
			},
		},
	}
	for _, c := range cases {
		path := fmt.Sprintf("./testdata/detail/%s.html", c.Code)
		doc := getDocFromPath(path)
		gotMap := parse(doc, detailHints)
		deepCheckSmap(t, gotMap, c.WantMap)
	}
}
