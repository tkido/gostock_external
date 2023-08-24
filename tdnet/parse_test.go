package tdnet

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
			"./testdata/3105/tse-qcedjpsm-31050-20170808331050-ixbrl.htm",
			Report{
				2017,
				1,
				"2017-08-08",
				"2017-06",
				113182000000,
				-708000000,
				1250000000,
				16450000000,
			},
		},
		{
			"./testdata/3085/tdnet-anedjpsm-30850-20090114061118.xbrl",
			Report{
				2008,
				4,
				"2009-02-10",
				"2008-12",
				7997000000,
				818000000,
				856000000,
				417000000,
			},
		},
		{
			"./testdata/3085/tdnet-qnedjpsm-30850-20090414053380.xbrl",
			Report{
				2009,
				1,
				"2009-04-30",
				"2009-03",
				2025000000,
				234000000,
				255000000,
				147000000,
			},
		},
		{
			"./testdata/3085/tdnet-acedjpsm-30850-20130127071019.xbrl",
			Report{
				2012,
				4,
				"2013-02-07",
				"2012-12",
				12797000000,
				1962000000,
				2013000000,
				1111000000,
			},
		},
		{
			"./testdata/3085/tdnet-qcedjpsm-30850-20131019039715.xbrl",
			Report{
				2013,
				3,
				"2013-10-29",
				"2013-09",
				10801000000,
				1620000000,
				1636000000,
				966000000,
			},
		},
		{
			"./testdata/3085/tse-acedjpsm-30850-20140124091132-ixbrl.htm",
			Report{
				2013,
				4,
				"2014-02-07",
				"2013-12",
				14986000000,
				2323000000,
				2359000000,
				1353000000,
			},
		},
		{
			"./testdata/6200/tse-acedjpsm-62000-20171110362000-ixbrl.htm",
			Report{
				2016,
				4,
				"2017-11-10",
				"2017-09",
				3585000000,
				592000000,
				608000000,
				412000000,
			},
		},
		{
			"./testdata/6200/tse-qcedjpsm-62000-20160722455582-ixbrl.htm",
			Report{
				2015,
				3,
				"2016-07-29",
				"2016-06",
				2097000000,
				359000000,
				357000000,
				224000000,
			},
		},
	}
	for _, c := range cases {
		fp, err := os.Open(c.Path)
		if err != nil {
			t.Fatal(err)
		}
		got, err := Parse(c.Path, fp)
		if err != nil {
			t.Error(err)
			return
		}
		want := c.Report
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
