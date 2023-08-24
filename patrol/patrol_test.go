package patrol

import (
	"testing"
)

func TestPatrol(t *testing.T) {
	filtered := Patrol([]string{"3085", "4235", "2121"})
	t.Log(filtered)
}
