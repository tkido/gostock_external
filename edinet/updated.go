package edinet

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// IsUpdatedIn3Days 3日以内にEdinetで更新があったかを返す
func IsUpdatedIn3Days(code string) (bool, error) {
	return IsUpdated(code, "1")
}

// IsUpdatedIn6Months 6ヶ月以内にEdinetで更新があったかを返す
func IsUpdatedIn6Months(code string) (bool, error) {
	return IsUpdated(code, "3")
}

// IsUpdated 指定の期間でEdinetで更新があったかを返す
func IsUpdated(code, pfs string) (updated bool, err error) {
	const sessionKey = "1534319823365"
	edinetCode, ok := EdinetCodeMap[code]
	if !ok {
		return false, errors.New("edinet.IsUpdated: unknown stock code " + code)
	}
	const tmpl = `https://disclosure.edinet-fsa.go.jp/E01EW/BLMainController.jsp?uji.verb=W1E63021CXP002002DSPSch&uji.bean=ee.bean.parent.EECommonSearchBean&PID=W1E63021&TID=W1E63021&SESSIONKEY=%s&lgKbn=2&pkbn=0&skbn=1&dskb=&askb=&dflg=0&iflg=0&preId=1&sec=%s&scc=&shb=&snm=&spf1=1&spf2=1&iec=&icc=&inm=&spf3=1&fdc=&fnm=&spf4=1&spf5=1&cal=1&era=H&yer=&mon=&psr=1&pfs=%s&row=0&idx=0&str=&kbn=1&flg=&syoruiKanriNo=`
	url := fmt.Sprintf(tmpl, sessionKey, edinetCode, pfs)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36"
	req.Header.Set("User-Agent", userAgent)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return
	}
	updated = doc.Find("#errorDisplayArea").Text() == ""
	return
}
