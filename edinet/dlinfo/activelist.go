package dlinfo

import (
	"sort"

	"github.com/tkido/gostock/my"
	"github.com/tkido/gostock/my/json"
)

// MakeActiveList update active list
func MakeActiveList(pActiveMap, pActiveList string) (err error) {
	m := map[string]bool{}
	err = json.Load(pActiveMap, &m)
	if err != nil {
		return
	}
	codes := []string{}
	for code, active := range m {
		if active {
			codes = append(codes, code)
		}
	}
	sort.Slice(codes, func(i, j int) bool {
		return codes[i] < codes[j]
	})
	err = my.WriteFile(pActiveList, codes)
	return
}
