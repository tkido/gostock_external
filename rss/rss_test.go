package rss

import (
	"testing"
)

func TestPrepare(t *testing.T) {
	err := Prepare([]string{"6200", "4235"})
	if err != nil {
		t.Error(err)
	}
}

func TestPublish(t *testing.T) {
	_, err := Publish([]string{"6200", "4235"})
	if err != nil {
		t.Error(err)
	}
}
