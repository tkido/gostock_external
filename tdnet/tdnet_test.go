package tdnet

import (
	"testing"
)

func TestIsUpdated(t *testing.T) {
	codes := []string{"4235", "6200", "3850"}
	for _, code := range codes {
		_, err := IsUpdated(code)
		if err != nil {
			t.Error(err)
		}
	}
}
