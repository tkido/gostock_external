package my

import (
	"testing"
)

func TestRound(t *testing.T) {
	cases := []struct {
		f    float64
		want string
	}{
		{1, "1.00"},
		{2000, "2.00K"},
		{3000000, "3.00M"},
		{4000000000, "4.00B"},
		{5000000000000, "5.00T"},
		{6000000000000000, "6.00Q"},
		{-1, "-1.00"},
		{-2000, "-2.00K"},
		{-3000000, "-3.00M"},
		{-4000000000, "-4.00B"},
		{-5000000000000, "-5.00T"},
		{-6000000000000000, "-6.00Q"},
		{12, "12.0"},
		{23000, "23.0K"},
		{34000000, "34.0M"},
		{45000000000, "45.0B"},
		{56000000000000, "56.0T"},
		{67000000000000000, "67.0Q"},
		{-12, "-12.0"},
		{-23000, "-23.0K"},
		{-34000000, "-34.0M"},
		{-45000000000, "-45.0B"},
		{-56000000000000, "-56.0T"},
		{-67000000000000000, "-67.0Q"},
		{123, "123"},
		{234000, "234K"},
		{345000000, "345M"},
		{456000000000, "456B"},
		{567000000000000, "567T"},
		{678000000000000000, "678Q"},
		{-123, "-123"},
		{-234000, "-234K"},
		{-345000000, "-345M"},
		{-456000000000, "-456B"},
		{-567000000000000, "-567T"},
		{-678000000000000000, "-678Q"},
	}
	for _, c := range cases {
		got := Round(c.f)
		want := c.want
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}