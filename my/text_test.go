package my

import (
	"fmt"
	"testing"

	"golang.org/x/text/unicode/norm"
)

func TestNFKC(t *testing.T) {
	got := norm.NFKC.String("０１２３４５６７８９ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ")
	want := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestRune(t *testing.T) {
	s := "テスト123"
	fmt.Println(len(s))
	fmt.Println(len([]rune(s)))
	for p, r := range s {
		fmt.Println(p)
		fmt.Println(r)
	}
}

func TestWidthCount(t *testing.T) {
	cases := []struct {
		String string
		Width  int
	}{
		{"12345", 5},
		{"あいうえお", 10},
		{"12345あいうえお", 15},
	}
	for _, c := range cases {
		got := Width(c.String)
		want := c.Width
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestTruncateWidth(t *testing.T) {
	cases := []struct {
		String string
		Max    int
		Want   string
	}{
		{"12345", 5, "12345"},
		{"12345", 10, "12345"},
		{"12345", 3, "123"},
		{"12345", 0, ""},
		{"あいうえお", 0, ""},
		{"あいうえお", 1, ""},
		{"あいうえお", 2, "あ"},
		{"あいうえお", 3, "あ"},
		{"あいうえお", 4, "あい"},
		{"あいうえお", 10, "あいうえお"},
		{"あいうえお", 11, "あいうえお"},
		{"あいうえお", 12, "あいうえお"},
		{"12345あいうえお", 15, "12345あいうえお"},
		{"1あ2い3う4え5お", 15, "1あ2い3う4え5お"},
		{"1あ2い3う4え5お", 20, "1あ2い3う4え5お"},
		{"1あ2い3う4え5お", 2, "1"},
		{"1あ2い3う4え5お", 3, "1あ"},
		{"1あ2い3う4え5お", 4, "1あ2"},
	}
	for _, c := range cases {
		got := TruncateWidth(c.String, c.Max)
		want := c.Want
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
