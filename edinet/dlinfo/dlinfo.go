package dlinfo

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// Record is a data per line in "EdinetcodeDlInfo.csv"
type Record struct {
	EdinetCode      string // 0
	PresenterType   string // 1
	Listing         bool   // 2
	Consolidated    bool   // 3
	Capital         int    // 4
	SettingDay      string // 5
	Presenter       string // 6
	PresenterEn     string // 7
	PresenterKana   string // 8
	Location        string // 9
	IndustryType    string // 10
	StockCode       string // 11
	CorporateNumber string // 12
}

func newRecord(ss []string) (Record, error) {
	if len(ss) != 13 {
		return Record{}, errors.New("invalid line")
	}
	listing := ss[2] == "上場"
	consolidated := ss[3] == "有"
	capital, err := strconv.Atoi(ss[4])
	if err != nil {
		capital = -1
	}
	rec := Record{ss[0], ss[1], listing, consolidated, capital, ss[5], ss[6], ss[7], ss[8], ss[9], ss[10], ss[11], ss[12]}
	return rec, nil
}

// ParseDlInfo parse "EdinetcodeDlInfo.csv" downloaded from Edinet
func ParseDlInfo(path string) ([]Record, error) {
	sjisF, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer sjisF.Close()
	f := transform.NewReader(sjisF, japanese.ShiftJIS.NewDecoder())
	r := csv.NewReader(f)
	r.Read() // drop first 2 lines
	r.Read() // they are unnecessary title lines
	rst := make([]Record, 0, 4096)
	for {
		ss, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else if strings.HasSuffix(err.Error(), "wrong number of fields") {
				// ignore known error
			} else {
				return nil, err
			}
		}
		rec, err := newRecord(ss)
		if err != nil {
			return nil, err
		}
		rst = append(rst, rec)
	}
	return rst, err
}
