package template

import (
	"testing"
)

func Test(t *testing.T) {
	m := map[string]string{
		"1": "壱",
		"2": "弐",
	}
	tmpl := `1は漢字で{{1}}、\n2は漢字で{{2}}と書きます。`
	got := Execute(tmpl, m)
	want := `1は漢字で壱、\n2は漢字で弐と書きます。`
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
