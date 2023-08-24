package edinet

// import (
// 	"testing"
// )

// func TestMakeReports(t *testing.T) {
// 	cases := []struct {
// 		Code      string
// 		FairValue float64
// 	}{
// 		{
// 			"3085",
// 			3.883841062804401e+10,
// 		},
// 	}
// 	for _, c := range cases {
// 		rs, err := MakeReports("./testdata/", c.Code)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		got := rs.FairValue()
// 		want := c.FairValue
// 		if got != want {
// 			t.Errorf("got %v want %v", got, want)
// 		}
// 	}
// }
