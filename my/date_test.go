package my

import "testing"

func TestNormDate(t *testing.T) {
	cases := []struct {
		given, want string
	}{
		{"平成２６年２月７日", "2014-02-07"},
		{"令和２年２月７日", "2020-02-07"},
		{`<span style="font-family: 'MS Gothic'">2020年1月27日</span>`, "2020-01-27"},
	}
	for _, c := range cases {
		got := NormDate(c.given)
		want := c.want
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
