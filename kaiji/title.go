package kaiji

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/tkido/gostock/my"
	"golang.org/x/text/unicode/norm"
)

func makeDate(url string) string {
	buf := bytes.Buffer{}
	l := len(url)
	buf.WriteString(url[l-13 : l-9]) // year
	buf.WriteString("-")
	buf.WriteString(url[l-9 : l-7]) // month
	buf.WriteString("-")
	buf.WriteString(url[l-7 : l-5]) // date
	return buf.String()
}

var reTrimTitle = my.MustCompile(`
	^【.*?\s #システムで付加されている先頭から最初の空白まで。e.g.『【99840】ソフトバンクグループ株式会社 』
  | [\s、・　\-－‐「」]                  #不要な空白と記号を除外
  | に関する | である                     #不要
  | (?:について)?の?(?:お知らせ)?$        #定型語尾
  | 当社(?:による|の|と)?                 #他社の開示をすることはない
  | 株式会社                              #株が関わることなのはわかっている
  | Limited\s?社?
  | ,?\s?Inc\. | ,?\s?Ltd\.
  | 余剰金の | 剰余金の                   #必ず「配当」などが続くので不要
  | 平成\d{1,2}年\d{1,2}月期の?と?        #日付で十分
  | \d{4}年\d{1,2}月期の?と?              #日付で十分
  | 第\d*?(?:回|期)(?:および)?            #不要な情報
  | 〔.*?〕 | \[.*?\]                     #ほぼ無意味
  `)
var reParen = regexp.MustCompile(`\(.*?\)`)

func trimTitle(title string) string {
	s := norm.NFKC.String(title)                                       // 後で全角括弧内を処理するため最初に半角統一
	s = reTrimTitle.ReplaceAllString(s, "")                            // 単純な除外処理
	s = reConvertTitle.ReplaceAllStringFuncSubmatches(s, convertTitle) // 後方参照が必要な複雑な変換処理
	s = reConvertTitle.ReplaceAllStringFuncSubmatches(s, convertTitle) // 2回行う
	s = reParen.ReplaceAllString(s, "")                                // 複雑な変換処理でも残った括弧を消す
	s = reTrimTitle.ReplaceAllString(s, "")                            // 変換によって変わった可能性があるので再度除外処理
	s = strings.TrimRight(s, "書")                                      //何らかの書であるのは自明
	s = strings.TrimSuffix(s, "報告")                                    //報告であるのも自明
	return s
}

var reTidyCompanyName = regexp.MustCompile(`\(株\)|株式会社|＆|ホールディングス?|コーポレーション|カンパニー|グループ|本社|ジャパン$`)

var reConvertTitle = my.MustCompile(`
	子会社\((.*?)\)		                       # $1: 子会社名
  | 一部(変更|報道|訂正|取り下げ)                 # $2
  | 株主(優待)制度                                # $3
  | (国際会計基準|国際財務報告基準)               # $4
  | 異動\((.*?)\)                                 # $5
  | 配当\((.*?)\)                                 # $6
  | 第(\d)四半期                                  # $7
  | (変更|訂正)報告書\(大量保有\)                 # $8
  | (決算短信)                                    # $9
  | (有価証券報告書)                              # $10
  | (四半期報告書)                                # $11
  | (東京証券取引所市場第一部)                    # $12
  | (ストックオプション)                          # $13
  | \((訂正)\)                                    # $14
  `)

func convertTitle(dollar []string) string {
	if s := dollar[1]; s != "" {
		return fmt.Sprintf("子会社%s", reTidyCompanyName.ReplaceAllString(s, ""))
	} else if s := dollar[2]; s != "" {
		return s
	} else if s := dollar[3]; s != "" {
		return s
	} else if s := dollar[4]; s != "" {
		return "IFRS"
	} else if s := dollar[5]; s != "" {
		return s
	} else if s := dollar[6]; s != "" {
		if len(s) <= 6 {
			s += "配当"
		}
		return s
	} else if s := dollar[7]; s != "" {
		return fmt.Sprintf("%sQ", s)
	} else if s := dollar[8]; s != "" {
		if s == "訂正" {
			return "訂正大量保有"
		}
		return "大量保有"
	} else if s := dollar[9]; s != "" {
		return "短信"
	} else if s := dollar[10]; s != "" {
		return "有報"
	} else if s := dollar[11]; s != "" {
		return "四報"
	} else if s := dollar[12]; s != "" {
		return "東証一部"
	} else if s := dollar[13]; s != "" {
		return "SO"
	} else if s := dollar[14]; s != "" {
		return s
	}
	return ""
}
