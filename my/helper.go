package my

// FactMap is map of Fact Data
type FactMap map[string]float64

// AnyOf returns the first data found
func (m FactMap) AnyOf(ss ...string) (f float64, ok bool) {
	for _, s := range ss {
		if f, ok = m[s]; ok {
			return
		}
	}
	return
}
