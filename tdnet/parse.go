package tdnet

import (
	"errors"
	"fmt"
	"io"
	"path"
	"strconv"
	"strings"

	"github.com/tkido/gostock/my"
	"github.com/tkido/gostock/xbrl"
)

// Parse XBRL files from tdnet
func Parse(key string, reader io.Reader) (r Report, err error) {
	x := &xbrl.XBRL{}
	err = xbrl.UnmarshalXBRL(x, reader)
	if err != nil {
		return
	}
	r.Quater, err = getQuater(x)
	if err != nil {
		return
	}
	r.FilingDate, err = getFilingDate(key, x)
	if err != nil {
		return
	}
	var context string
	if strings.HasSuffix(key, ".xbrl") {
		context = getContext(key, r.Quater, x)
	} else {
		context = getContextInline(key, r.Quater)
	}
	r.Year, r.EndMonth, err = getYearEndMonth(context, x)
	if err != nil {
		return
	}
	// NetSales, OperatingIncome, OrdinaryIncome, NetIncome
	m := my.FactMap{}
	for _, f := range x.Facts {
		if f.ContextRef == context {
			fl, err := f.Float()
			if err != nil {
				// log.Println(err)
				continue
			}
			m[f.Name] = fl
		}
	}
	var ok bool
	r.NetSales, ok = m.AnyOf("NetSales", "OperatingRevenues")
	if !ok {
		return r, fmt.Errorf("NetSales not found in %s", key)
	}
	r.OperatingIncome, ok = m.AnyOf("OperatingIncome")
	if !ok {
		return r, fmt.Errorf("OperatingIncome not found in %s", key)
	}
	r.OrdinaryIncome, ok = m.AnyOf("OrdinaryIncome")
	if !ok {
		return r, fmt.Errorf("OrdinaryIncome not found in %s", key)
	}
	r.NetIncome, ok = m.AnyOf("NetIncome", "ProfitAttributableToOwnersOfParent")
	if !ok {
		return r, fmt.Errorf("NetIncome not found in %s", key)
	}
	return r, nil
}

func getQuater(x *xbrl.XBRL) (quater int, err error) {
	for _, f := range x.Facts {
		if f.Name == "QuarterlyPeriod" {
			quater, err = strconv.Atoi(f.Value)
			if err != nil {
				return
			}
			return
		}
	}
	return 4, nil
}

func getFilingDate(key string, x *xbrl.XBRL) (string, error) {
	for _, f := range x.Facts {
		if f.Name == "FilingDate" {
			return my.NormDate(f.Value), nil
		}
	}
	return "", fmt.Errorf("filingDate not found in %s", key)
}

func getContext(key string, quater int, x *xbrl.XBRL) (context string) {
	fileName := path.Base(key)
	isConsolidated := fileName[7:8] == "c"
	if quater == 4 {
		if isConsolidated {
			context = "CurrentYearConsolidatedDuration"
		} else {
			context = "CurrentYearNonConsolidatedDuration"
		}
	} else {
		for _, c := range x.Contexts {
			if c.ID == "CurrentQuarterConsolidatedDuration" || c.ID == "CurrentQuarterNonConsolidatedDuration" {
				context = c.ID
				return
			}
		}
		if isConsolidated {
			context = fmt.Sprintf("CurrentAccumulatedQ%dConsolidatedDuration", quater)
		} else {
			context = fmt.Sprintf("CurrentAccumulatedQ%dNonConsolidatedDuration", quater)
		}
	}
	return
}

func getContextInline(key string, quater int) (context string) {
	fileName := path.Base(key)
	isConsolidated := fileName[5:6] == "c"
	if quater == 4 {
		if isConsolidated {
			context = "CurrentYearDuration_ConsolidatedMember_ResultMember"
		} else {
			context = "CurrentYearDuration_NonConsolidatedMember_ResultMember"
		}
	} else {
		if isConsolidated {
			context = fmt.Sprintf("CurrentAccumulatedQ%dDuration_ConsolidatedMember_ResultMember", quater)
		} else {
			context = fmt.Sprintf("CurrentAccumulatedQ%dDuration_NonConsolidatedMember_ResultMember", quater)
		}
	}
	return
}

func getYearEndMonth(context string, x *xbrl.XBRL) (year int, endMonth string, err error) {
	for _, c := range x.Contexts {
		if c.ID == context {
			year, err = strconv.Atoi(c.Start[:4])
			if err != nil {
				return 0, "", err
			}
			endMonth = c.End[:7]
			return
		}
	}
	return 0, "", errors.New("year and endMonth not found")
}
