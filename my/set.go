package my

// NilStruct is zero byte struct
type NilStruct struct{}

// Nil is singleton of NilStruct
var Nil = NilStruct{}

// SSet is set of string
type SSet map[string]NilStruct

// NewSSet make new SSet
func NewSSet(ss ...string) SSet {
	set := SSet{}
	for _, s := range ss {
		set[s] = Nil
	}
	return set
}

// Add adds new strings to Set
func (set SSet) Add(ss ...string) {
	for _, s := range ss {
		set[s] = Nil
	}
}

// And makes product set
func (set SSet) And(other SSet) SSet {
	new := NewSSet()
	for k := range set {
		if _, ok := other[k]; ok {
			new[k] = Nil
		}
	}
	return new
}

// Sub makes difference set
func (set SSet) Sub(other SSet) SSet {
	new := NewSSet()
	for k := range set {
		if _, ok := other[k]; !ok {
			new[k] = Nil
		}
	}
	return new
}

// Or makes sum of sets
func (set SSet) Or(other SSet) SSet {
	new := NewSSet()
	for k := range set {
		new[k] = Nil
	}
	for k := range other {
		new[k] = Nil
	}
	return new
}

// Has returns set has the string
func (set SSet) Has(s string) bool {
	_, ok := set[s]
	return ok
}
