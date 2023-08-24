package spider

import (
	"fmt"
	"testing"
)

func TestParseLonghis(t *testing.T) {
	cases := []struct {
		Code    string
		WantMap Smap
	}{
		{"6200",
			Smap{
				"数年来高値":  "=3090/【値】",
				"数年来高値月": "2021/12",
				"数年来安値":  "=740/【値】",
				"数年来安値月": "2020/4",
			},
		},
		{"4235",
			Smap{
				"数年来高値":  "=5950/【値】",
				"数年来高値月": "2022/9",
				"数年来安値":  "=673/【値】",
				"数年来安値月": "2020/3",
			},
		},
	}
	for _, c := range cases {
		gotMap := parseLonghisTest(c.Code)
		deepCheckSmap(t, gotMap, c.WantMap)
	}
}

func parseLonghisTest(code string) Smap {
	path1 := fmt.Sprintf("./testdata/longhis/%s_1.html", code)
	doc1 := getDocFromPath(path1)
	days := getDays(doc1)
	path2 := fmt.Sprintf("./testdata/longhis/%s_2.html", code)
	doc2 := getDocFromPath(path2)
	days2 := getDays(doc2)
	days = append(days, days2...)
	return processLonghis(days)
}
