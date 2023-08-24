package kaiji

import "testing"

func TestTrimTitle(t *testing.T) {
	cases := []struct {
		want, title string
	}{
		{"株式上場準備の開始", "【99840】ソフトバンクグループ株式会社 株式上場準備の開始"},
		{"合併", "合併について"},
		{"合併", "合併に関するお知らせ"},
		{"合併", "合併のお知らせ"},
		{"子会社ベクターの業績予想の修正", "当社子会社（株式会社ベクター）の業績予想の修正について"},
		{"変更", "一部変更について"},
		{"報道", "一部報道について"},
		{"訂正", "一部訂正について"},
		{"優待の変更", "株主優待制度の変更について"},
		{"配当", "余剰金の配当"},
		{"配当", "剰余金の配当"},
		{"中間配当", "剰余金の配当（中間配当）に関するお知らせ"},
		{"期末配当", "剰余金の配当（期末）に関するお知らせ"},
		{"IFRS", "国際会計基準"},
		{"大量保有", "変更報告書（大量保有）"},
		{"訂正大量保有", "訂正報告書（大量保有）"},
		{"代表取締役の追加選定", "代表取締役の異動（追加選定）に関するお知らせ"},
		{"無担保社債の発行", "第41回および第42回無担保社債の発行に関するお知らせ"},
		{"1Q短信", "第1四半期決算短信"},
		{"2Q短信", "第2四半期決算短信"},
		{"3Q短信", "第3四半期決算短信"},
		{"ボーダフォン日本法人買収資金のリファイナンス", "ボーダフォン日本法人買収資金のリファイナンスについて（２）"},
		{"有報", "有価証券報告書－第37期(平成28年4月1日－平成29年3月31日)"},
		{"3Q決算ハイライト", "平成21年３月期 第３四半期決算ハイライト"},
		{"四報1Q", "四半期報告書－第37期第1四半期(平成28年4月1日－平成28年6月30日)"},
		{"大量保有", "【E02778】ソフトバンクグループ株式会社 変更報告書（大量保有）"},
		{"四報1Q", "【E02778】ソフトバンクグループ株式会社 四半期報告書 ‐ 第30期 第1四半期（平成21年4月1日 ‐ 平成21年6月30日）"},
		{"訂正発行登録", "【E02778】ソフトバンクグループ株式会社 訂正発行登録書"},
		{"内部統制", "【E02778】ソフトバンクグループ株式会社 内部統制報告書 ‐ 第29期（平成20年4月1日 ‐ 平成21年3月31日）"},
		{"確認", "【E02778】ソフトバンクグループ株式会社 確認書"},
		{"臨時", "【E02778】ソフトバンクグループ株式会社 臨時報告書"},
		{"子会社ヤフーによるジャパンネット銀行の子会社化", "【99840】ソフトバンクグループ株式会社 当社子会社（ヤフー株式会社）による株式会社ジャパンネット銀行の子会社化に関するお知らせ"},
		{"子会社ケンコーコム株券等に対する公開買付けの結果", "【47550】楽天株式会社 子会社であるケンコーコム株式会社株券等に対する公開買付けの結果に関するお知らせ"},
		{"Ebatesの株式の取得完了", "Ebates Inc.の株式の取得完了に関するお知らせ"},
		{"ViberMediaの株式の取得", "Viber Media Ltd.の株式の取得（子会社化）に関するお知らせ"},
		{"東証一部への上場市場変更", "東京証券取引所 市場第一部への上場市場変更に関するお知らせ"},
		{"SOの付与", "ストックオプション（新株予約権）の付与について"},
		{"訂正株式交換により発行する新株式数の確定", "（訂正）株式交換により発行する新株式数の確定に関するお知らせ"},
		{"訂正英国進出に向けたPlayHoldingsの買収", "（訂正）「英国進出に向けたPlay Holdings Limited 社の買収に関するお知らせ」"},
		{"IFRSに基づく連結財務諸表", "「国際財務報告基準(IFRS) に基づく連結財務諸表」（IFRS 任意報告書）についてのお知らせ"},
		{"連結子会社楽天証券に対する金融庁の行政処分", "連結子会社（楽天証券株式会社）に対する金融庁の行政処分について"},
		{"連結子会社みんなの就職の合併", "当社と連結子会社（みんなの就職株式会社）の合併について"},
		{"", ""},
	}
	for _, c := range cases {
		got := trimTitle(c.title)
		if got != c.want {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}