package template

import "github.com/tkido/gostock/my"

var reVar = my.MustCompile(`\{\{(.*?)\}\}`)

// Execute is Execute
func Execute(t string, m map[string]string) string {
	return reVar.ReplaceAllStringFuncSubmatches(
		t,
		func(d []string) string {
			return m[d[1]]
		},
	)
}
