package my

import (
	"fmt"
	"regexp"
	"strconv"

	"golang.org/x/text/unicode/norm"
)

var reDateJp = regexp.MustCompile(`.*?(平成|令和|20)(\d{1,2})年(\d{1,2})月(\d{1,2})日`)

// NormDate convert e.g. 平成26年2月7日 => 2014-02-07
func NormDate(jp string) string {
	s := norm.NFKC.String(jp)
	sms := reDateJp.FindStringSubmatch(s)
	if len(sms) != 5 {
		return s
	}
	era := sms[1]
	year, _ := strconv.Atoi(sms[2])
	month, _ := strconv.Atoi(sms[3])
	day, _ := strconv.Atoi(sms[4])
	var ad int
	switch era {
	case "平成":
		ad = 1988 + year
	case "令和":
		ad = 2018 + year
	default:
		ad = 2000 + year
	}
	return fmt.Sprintf("%04d-%02d-%02d", ad, month, day)
}

// IsLeapYear returns
func IsLeapYear(year int) bool {
	switch {
	case year%400 == 0:
		return true
	case year%100 == 0:
		return false
	case year%4 == 0:
		return true
	default:
		return false
	}
}
