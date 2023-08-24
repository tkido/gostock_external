package spider

import (
	"reflect"
	"testing"
)

// TestSmap 順番が違ってもmapのDeepEqualが有効かの確認
func TestSmap(t *testing.T) {
	want := Smap{"現値": "3080", "発行": "6800000", "前終": "3205"}
	got := Smap{"発行": "6800000", "前終": "3205", "現値": "3080"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v want %v", got, want)
	}
}
