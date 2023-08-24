package dlinfo

import (
	"io/ioutil"
	"os"
	"testing"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func TestShifJis(t *testing.T) {
	sjisF, err := os.Open("./testdata/shift-jis.txt")
	if err != nil {
		t.Error(err)
	}
	defer sjisF.Close()
	f := transform.NewReader(sjisF, japanese.ShiftJIS.NewDecoder())
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
	}
	got := string(bs)
	want := "Shift-JISで書かれた日本語のテスト。"
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
