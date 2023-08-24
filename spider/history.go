package spider

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tkido/gostock/statistics"
)

var historyURLTmpl = "https://finance.yahoo.co.jp/quote/%s/history"

// Day is data from a business day
type Day struct {
	Date                                                                      string
	Open, High, Low, Close, Volume, AdjustedClose, LastClose, Buy, Sell, Move float64
}

// Days is []Day
type Days []Day

func parseHistory(doc *goquery.Document) Smap {
	days := getDays(doc)
	days.adjustSplit()
	days.addLast()
	days.calcBuySellMove()

	const volumeTmpl = "=%d/【発行】*1000" // 平均出来高は最終的には発行済総株式数の千分率で表示される
	const deviationTmpl = "=1-%f/【値】"

	return Smap{
		"Vol": fmt.Sprintf("%f", days.volatility()),
		"SPR": fmt.Sprintf("%f", days.spr()),
		"RCI": fmt.Sprintf("%f", days.rci()),
		"日出":  fmt.Sprintf(volumeTmpl, days.volumePerDays(1)),
		"週出":  fmt.Sprintf(volumeTmpl, days.volumePerDays(5)),
		"月出":  fmt.Sprintf(volumeTmpl, days.volumePerDays(20)),
		"乖離":  fmt.Sprintf(deviationTmpl, days.average()),
	}
}

// td のコレクションを日毎のデータ（Day）に詰める
func getDays(doc *goquery.Document) Days {
	texts := []string{}
	dates := []string{}
	trs := doc.Find(`#root > main > div > div > div.XuqDlHPN > div:nth-child(3) > section._3DmkswWx._1c78LjU4._2P-X1dj1 > div > table > tbody > tr`)
	trs.Each(func(i int, tr *goquery.Selection) {
		tds := tr.Find(`td`)
		isSplit := false
		tds.Each(func(i int, td *goquery.Selection) {
			text := td.Text()
			isSplit = strings.HasPrefix(text, "分割")
			if !isSplit {
				texts = append(texts, text)
			}
		})
		if !isSplit {
			ths := tr.Find(`th`)
			ths.Each(func(i int, th *goquery.Selection) {
				dates = append(dates, th.Text())
			})
		}
	})
	// []string -> Days の変換
	days := Days{}
	var day Day
	for i, text := range texts {
		mod := i % 6
		num, _ := strconv.ParseFloat(rmComma(text), 64)
		switch mod {
		case 0:
			day.Date = dates[i/6]
			day.Open = num
		case 1:
			day.High = num
		case 2:
			day.Low = num
		case 3:
			day.Close = num
		case 4:
			day.Volume = num
		case 5:
			day.AdjustedClose = num
			days = append(days, day)
		default:
			log.Fatal("Must Not Happen!!")
		}
	}
	return days
}

// 株式分割or併合を補正
func (days Days) adjustSplit() {
	for i, day := range days {
		rate := day.Close / day.AdjustedClose
		day.Open /= rate
		day.High /= rate
		day.Low /= rate
		day.Close /= rate
		day.Volume *= rate
		days[i] = day
	}
}

// 前日終値を取得。最後のものだけ当日始値で代用。
// そのためだけにもう1ページ取得するには及ばないため。
func (days Days) addLast() {
	for i, day := range days {
		if i == len(days)-1 {
			day.LastClose = day.Open
		} else {
			day.LastClose = days[i+1].Close
		}
		days[i] = day
	}
}

// 買い・売りとその合計の値動き
func (days Days) calcBuySellMove() {
	for i, day := range days {
		var deltas []float64
		if day.Close > day.Open {
			deltas = []float64{day.LastClose - day.Open, day.Open - day.Low, day.Low - day.High, day.High - day.Close}
		} else {
			deltas = []float64{day.LastClose - day.Open, day.Open - day.High, day.High - day.Low, day.Low - day.Close}
		}
		for _, delta := range deltas {
			if delta < 0 {
				day.Buy -= delta
			} else {
				day.Sell += delta
			}
			day.Move = day.Buy + day.Sell
		}
		days[i] = day
	}
}

// Vol (volatility) の計算
func (days Days) volatility() float64 {
	var moveSum, closeSum float64
	for _, day := range days {
		moveSum += day.Move
		closeSum += day.Close
	}
	return moveSum / closeSum
}

// 終値のRCI (Rank Correlation Index)の計算
// 1.0が単調上昇-1.0が単調下落。
// 渡すときに日付が古い方を前にするためにReverseする。
func (days Days) rci() float64 {
	closes := make([]float64, len(days))
	for i, day := range days {
		closes[len(days)-1-i] = day.Close
	}
	return statistics.RciFloat(closes)
}

// SPR (Selling Pressure Ratio)の計算
func (days Days) spr() float64 {
	var buyP, sellP float64
	for _, day := range days {
		if day.Move != 0 {
			buyP += day.Volume * day.Buy / day.Move
			sellP += day.Volume * day.Sell / day.Move
		}
	}
	return sellP / buyP
}

// n日前までの日平均出来高
func (days Days) volumePerDays(n int) int {
	var sum float64
	for i, day := range days {
		if i == n {
			break
		}
		sum += day.Volume
	}
	return int(sum / float64(n))
}

// 平均値の計算
func (days Days) average() float64 {
	var sum float64
	for _, day := range days {
		sum += day.Close
	}
	return sum / float64(len(days))
}
