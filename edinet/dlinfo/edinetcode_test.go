package dlinfo

import (
	"os"
	"path/filepath"
	"testing"
)

// "EdinetcodeDlInfo.csv" を更新したらこのテストを走らせる
func TestMakeCodeMap(t *testing.T) {
	src := "./testdata/EdinetcodeDlInfo.csv"
	rst := filepath.Join(os.TempDir(), "edinetcodemap.json")
	err := MakeCodeMap(src, rst)
	if err != nil {
		t.Error(err)
	}
}
