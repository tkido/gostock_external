package tdnet

import (
	"fmt"

	"github.com/tkido/gostock/my"
	"github.com/tkido/gostock/spider"
)

// Smap is map[string]string
type Smap map[string]string

// MakeUpdated は短信が更新された銘柄のデータ一覧のHTMLを返す
func MakeUpdated(dls []Downloaded) (html string, err error) {
	t := my.NewTable("新TDnet短信（金額および前年同期比）", "numbers")
	t.Th("code", "名称", "終了月", "Q", "売上", "営利", "経利", "純利", "売上", "営利", "経利", "純利", "開示日")
	format := `<a href="%s" target="_blank" title="%s">%s</a>`
	for _, dl := range dls {
		smap := spider.Get(dl.Code)
		name := smap["名称"]
		code := fmt.Sprintf(format, dl.URL, dl.Code, dl.Code)
		ps, ds, err := MakeReports(dl.Code)
		if err != nil {
			return "", err
		}
		if len(ds) == 0 {
			continue
		}
		m := map[int]PerReport{}
		for _, p := range ps {
			m[p.id()] = p
		}
		d := ds[0]
		if p, ok := m[d.id()]; ok {
			t.Td(code, name, d.EndMonth, d.Quater, d.NetSales, d.OperatingIncome, d.OrdinaryIncome, d.NetIncome, p.NetSales, p.OperatingIncome, p.OrdinaryIncome, p.NetIncome, d.FilingDate)
		} else {
			t.Td(code, name, d.EndMonth, d.Quater, d.NetSales, d.OperatingIncome, d.OrdinaryIncome, d.NetIncome, "", "", "", "", d.FilingDate)
		}
	}
	return fmt.Sprintf(template, t.String()), nil
}

var template = `<html lang="jp">
<head>
  <meta charset="utf-8">
  <meta name="robots" content="noindex,nofollow">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>短信更新銘柄一覧</title>
  <!-- Bootstrap core CSS -->
  <link href="https://xfomax.com/gostock/css/bootstrap.min.css" rel="stylesheet">
  <!-- Custom styles for this template -->
  <link href="https://xfomax.com/gostock/css/stock.css" rel="stylesheet">
</head>
<body>
%s
</body>
</html>
`
