package spider

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseProfilePage(t *testing.T) {
	cases := []struct {
		Code    string
		WantMap Smap
	}{
		{"6200", Smap{"名称": "インソース", "特色": "【特色】企業等の人事部向けに講師派遣型研修、公開講座を運営。人事や営業サポートシステムも展開", "事業": "【連結事業】講師派遣型研修51、公開講座24、ＩＴサービス14、他11(2021.9)", "分類": "サービス", "設立": "2002", "上場": "2016", "決期": "9月", "従連": "570", "従単": "352", "齢": "31.4", "収": "472", "代表": "舟橋孝之", "市": "東P", "市記号": "T"}},
		{"4235", Smap{"名称": "ウルトラファブリックス", "特色": "【特色】湿式合成皮革で先駆。１７年２月米国販社のウルトラファブリックス買収、持株会社下に製販統合", "事業": "【連結事業】家具用29、自動車用40、航空機用7、他24【海外】96(2021.12)", "分類": "化学", "設立": "1966", "上場": "2003", "決期": "12月", "従連": "317", "従単": "187", "齢": "", "収": "", "代表": "吉村昇", "市": "東S", "市記号": "T"}},
	}
	for _, c := range cases {
		path := fmt.Sprintf("./testdata/profile/%s.html", c.Code)
		doc := getDocFromPath(path)
		gotMap := parse(doc, profileHints)
		deepCheckSmap(t, gotMap, c.WantMap)
	}
}

func TestTrimSpace(t *testing.T) {
	s := "　 "
	if "" != strings.TrimSpace(s) {
		t.Error()
	}
}
