package xbrl

import (
	"encoding/xml"
	"os"
	"testing"
)

func TestParseXbrl(t *testing.T) {
	_, err := ParseXbrl("./testdata/tdnet/6200/tse-qcedjpsm-62000-20160722455582-ixbrl.htm")
	if err != nil {
		t.Error(err)
	}
}

func ParseXbrl(path string) (*XBRL, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	x := &XBRL{}
	err = UnmarshalXBRL(x, fp)
	if err != nil {
		return nil, err
	}
	return x, nil
}

func TestFloat(t *testing.T) {
	f := &Fact{XMLName: xml.Name{Space: "http://www.xbrl.org/2008/inlineXBRL", Local: "nonFraction"}, Name: "OperatingIncome", Value: "22,472", ContextRef: "CurrentAccumulatedQ2Duration_ConsolidatedMember_ResultMember", UnitRef: "JPY", Decimals: "-6", Scale: "6", Sign: "-", Nil: false}
	got, err := f.Float()
	if err != nil {
		t.Fatal(err)
	}
	want := -22472000000.0
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
