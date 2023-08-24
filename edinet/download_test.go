package edinet

import (
	"testing"
	"time"
)

func TestAddDate(t *testing.T) {
	start := time.Date(2009, 12, 1, 0, 0, 0, 0, time.UTC)
	got := start.AddDate(0, -3, 0)
	want := time.Date(2009, 9, 1, 0, 0, 0, 0, time.UTC)
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestAfter(t *testing.T) {
	t1 := time.Date(2009, 12, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2009, 12, 1, 0, 0, 0, 0, time.UTC)
	got := t1.After(t2)
	want := false
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestDownload(t *testing.T) {
	err := Download([]string{"3085"}, true)
	if err != nil {
		t.Error(err)
	}
}
