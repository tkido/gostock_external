package dlinfo

import "testing"

func TestMakeActiveList(t *testing.T) {
	const (
		src = "./testdata/activemap.json"
		rst = "./testdata/activelist.txt"
	)
	err := MakeActiveList(src, rst)
	if err != nil {
		t.Error(err)
	}
}
