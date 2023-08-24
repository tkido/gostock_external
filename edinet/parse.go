package edinet

import (
	"fmt"
	"io"
	"log"

	"github.com/tkido/gostock/my"
	"github.com/tkido/gostock/xbrl"
)

var contexts = map[string]bool{
	"CurrentYearConsolidatedInstant":            true,
	"CurrentYearConsolidatedDuration":           true,
	"CurrentYearNonConsolidatedInstant":         false,
	"CurrentYearNonConsolidatedDuration":        false,
	"CurrentYearInstant":                        true,
	"CurrentYearDuration":                       true,
	"CurrentYearInstant_NonConsolidatedMember":  false,
	"CurrentYearDuration_NonConsolidatedMember": false,
}

// Parse XBRL files from tdnet
func Parse(key string, reader io.Reader) (r Report, err error) {
	x := &xbrl.XBRL{}
	err = xbrl.UnmarshalXBRL(x, reader)
	if err != nil {
		return
	}

	r.FilingDate, err = getFilingDate(key, x)
	if err != nil {
		return
	}
	r.EndMonth, err = getEndMonth(key, x)
	if err != nil {
		return
	}

	m := my.FactMap{}
	cm := my.FactMap{}
	for _, f := range x.Facts {
		if consolidated, ok := contexts[f.ContextRef]; ok {
			fl, err := f.Float()
			if err != nil {
				// log.Println(err)
				continue
			}
			if consolidated {
				cm[f.Name] = fl
			} else {
				m[f.Name] = fl
			}
		}
	}
	if len(cm) >= len(m) {
		m = cm
	}

	if v, ok := m.AnyOf("NetSales", "NetSalesSummaryOfBusinessResults", "NetSalesOfCompletedConstructionContractsCNS"); ok {
		r.NetSales = v
	} else {
		// return r, fmt.Errorf("NetSales not found in %s", key)
		log.Printf("NetSales not found in %s", key)
	}
	if v, ok := m.AnyOf("GrossProfit", "GrossProfitOnCompletedConstructionContractsCNS"); ok {
		r.GrossProfitRatio = v / r.NetSales
	} else {
		// return r, fmt.Errorf("GrossProfit not found in %s", key)
		log.Printf("GrossProfit not found in %s", key)
	}
	if v, ok := m.AnyOf("OperatingIncome"); ok {
		r.OperatingProfitRatio = v / r.NetSales
	} else {
		// return r, fmt.Errorf("OperatingIncome not found in %s", key)
		log.Printf("OperatingIncome not found in %s", key)
	}
	if v, ok := m.AnyOf("OrdinaryIncome", "OrdinaryIncomeLossSummaryOfBusinessResults"); ok {
		r.OrdinaryProfitRatio = v / r.NetSales
	} else {
		// return r, fmt.Errorf("OrdinaryIncome not found in %s", key)
		log.Printf("OrdinaryIncome not found in %s", key)
	}
	if v, ok := m.AnyOf("ProfitLossAttributableToOwnersOfParentSummaryOfBusinessResults", "ProfitAttributableToOwnersOfParent", "NetIncome", "NetIncomeLossSummaryOfBusinessResults", "ProfitLoss"); ok {
		r.NetIncome = v
		r.NetProfitRatio = v / r.NetSales
	} else {
		// return r, fmt.Errorf("NetIncome not found in %s", key)
		log.Printf("NetIncome not found in %s", key)
	}
	// BreakupValue, NetCash, Accruals, FreeCashFlow
	for k, w := range weights["BreakupValue"] {
		if v, ok := m[k]; ok {
			r.BreakupValue += v * w
		}
	}
	for k, w := range weights["NetCash"] {
		if v, ok := m[k]; ok {
			r.NetCash += v * w
		}
	}
	for k, w := range weights["Accruals"] {
		if v, ok := m[k]; ok {
			r.Accruals += v * w
		}
	}
	for k, w := range weights["FreeCashFlow"] {
		if v, ok := m[k]; ok {
			r.FreeCashFlow += v * w
		}
	}
	return r, nil
}

var filingDateContextIDs = map[string]bool{
	"FilingDateInstant": true,
	"DocumentInfo":      true,
}

func getFilingDate(key string, x *xbrl.XBRL) (string, error) {
	for _, c := range x.Contexts {
		if filingDateContextIDs[c.ID] {
			return c.Instant, nil
		}
	}
	return "", fmt.Errorf("filingDate not found in %s", key)
}

func getEndMonth(key string, x *xbrl.XBRL) (string, error) {
	for _, c := range x.Contexts {
		if _, ok := contexts[c.ID]; ok {
			if c.End != "" {
				return c.End[:7], nil
			}
		}
	}
	return "", fmt.Errorf("endMonth not found in %s", key)
}
