package spider

// Smap is map[string]string
type Smap map[string]string

// Sanitizer is Sanitizer
type Sanitizer func(string) string

// Hint is hint of parse
type Hint struct {
	label    string    // データの名前・ラベル
	selector string    // Chromeのdev toolで取得できるものと通常同じ
	sanitize Sanitizer // goqueryから取得した後の処理
}
