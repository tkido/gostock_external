package ufo

// func TestGetFeed(t *testing.T) {
// 	cases := []struct {
// 		Type, Code, Title string
// 	}{
// 		{"edinet", "6200", "有報キャッチャー - EDINET情報配信サービス"},
// 		{"edinetx", "6200", "有報キャッチャー - EDINET情報配信サービス"},
// 		{"tdnet", "6200", "有報キャッチャー - 適時開示情報配信サービス"},
// 		{"tdnetx", "6200", "有報キャッチャー - 適時開示情報配信サービス"},
// 	}
// 	for _, c := range cases {
// 		f, err := GetFeed(c.Type, c.Code)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		got := f.Title
// 		want := c.Title
// 		if got != want {
// 			t.Errorf("got %v want %v", got, want)
// 		}
// 	}
// }
