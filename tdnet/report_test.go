package tdnet

import (
	"testing"
)

func TestDoukihi(t *testing.T) {
	cases := []struct {
		current  float64
		previous float64
		want     string
	}{
		{1.0, 1.0, "0%"},
		{2.0, 1.0, "100%"},
		{0.5, 1.0, "-50%"},
		{1.0, -1.0, `+黒転`},
		{-1.0, -1.0, `-赤字`},
		{-1.0, 1.0, `-赤転`},
		{-2.0, -1.0, `-赤拡`},
		{-1.0, -2.0, `-赤縮`},
	}
	for _, c := range cases {
		got := Doukihi{c.current, c.previous}.String()
		want := c.want
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestMakeReports(t *testing.T) {
	_, _, err := MakeReports("3085")
	if err != nil {
		t.Fatal(err)
	}
}
