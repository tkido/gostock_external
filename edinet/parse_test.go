package edinet

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	cases := []struct {
		Path string
		Report
	}{
		{
			"./testdata/3085/jpcrp030000-asr-001_E03513-000_2016-12-31_01_2017-03-31.xbrl",
			Report{EndMonth: "2016-12", FilingDate: "2017-03-31", NetSales: 2.3286e+10, NetIncome: 2.069e+09, BreakupValue: 9.7073e+09, NetCash: 1.1337e+10, Accruals: -6.242e+08, FreeCashFlow: 1.959e+09, GrossProfitRatio: 0.5298033152967448, OperatingProfitRatio: 0.14386326548140513, OrdinaryProfitRatio: 0.14828652409172893, NetProfitRatio: 0.08885167053164991},
		},
		{
			"./testdata/4235/jpcrp030000-asr-001_E01061-000_2014-03-31_01_2014-06-25.xbrl",
			Report{EndMonth: "2014-03", FilingDate: "2014-06-25", NetSales: 4.218995e+09, NetIncome: 3.99145e+08, BreakupValue: 1.7119923e+09, NetCash: 1.090977e+09, Accruals: -2.199334e+08, FreeCashFlow: 5.00503e+08, GrossProfitRatio: 0.2373079370798022, OperatingProfitRatio: 0.11764033851663726, OrdinaryProfitRatio: 0.14587028427386142, NetProfitRatio: 0.0946066539543185},
		},
		{
			"./testdata/3085/jpfr-asr-E03513-000-2009-12-31-01-2010-03-29.xbrl",
			Report{EndMonth: "2009-12", FilingDate: "2010-03-29", NetSales: 8.361485e+09, NetIncome: 5.00888e+08, BreakupValue: 1.0849085e+09, NetCash: 1.487705e+09, Accruals: -3.241976e+08, FreeCashFlow: 7.31654e+08, GrossProfitRatio: 0.582983525055657, OperatingProfitRatio: 0.11057844390081427, OrdinaryProfitRatio: 0.11533585242334346, NetProfitRatio: 0.05990419165973508},
		},
	}
	for _, c := range cases {
		fp, err := os.Open(c.Path)
		if err != nil {
			t.Fatal(err)
		}
		got, err := Parse(c.Path, fp)
		fp.Close()
		if err != nil {
			t.Fatal(err)
		}
		want := c.Report
		if got != want {
			t.Errorf("got %#v want %#v", got, want)
		}
	}
}
