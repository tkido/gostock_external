package statistics

import (
	"testing"
)

func TestRci(t *testing.T) {
	cases := []struct {
		Want float64
		Arr  interface{}
	}{
		{-1, []int{4, 3, 2, 1}},
		{1, []float64{1.0, 2.0, 3.0, 4.0}},
	}
	for _, c := range cases {
		got, err := Rci(c.Arr)
		if err != nil {
			t.Error(err)
		}
		if got != c.Want {
			t.Errorf("got %v want %v", got, c.Want)
		}
	}
}

func TestRciInt(t *testing.T) {
	cases := []struct {
		Want float64
		Arr  []int
	}{
		{-1, []int{4, 3, 2, 1}},
		{1, []int{1, 2, 3, 4}},
	}
	for _, c := range cases {
		got := RciInt(c.Arr)
		if got != c.Want {
			t.Errorf("got %v want %v", got, c.Want)
		}
	}
}

func TestRciFolat(t *testing.T) {
	cases := []struct {
		Want float64
		Arr  []float64
	}{
		{-1, []float64{4.0, 3.0, 2.0, 1.0}},
		{1, []float64{1.0, 2.0, 3.0, 4.0}},
	}
	for _, c := range cases {
		got := RciFloat(c.Arr)
		if got != c.Want {
			t.Errorf("got %v want %v", got, c.Want)
		}
	}
}
