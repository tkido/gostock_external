package dlinfo

import "github.com/tkido/gostock/my/json"

// MakeCodeMap makes map from StockCode to EdinetCode
func MakeCodeMap(src, rst string) (err error) {
	rs, err := ParseDlInfo(src)
	if err != nil {
		return
	}
	m := make(map[string]string, 4096)
	for _, r := range rs {
		if r.StockCode == "" || r.EdinetCode == "" {
			continue
		}
		code := r.StockCode[:4]
		m[code] = r.EdinetCode
	}
	err = json.Save(rst, m)
	return
}
