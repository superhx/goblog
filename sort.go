package goblog

import "sort"

//Item ...
type Item interface{}

//By ...
type By func(i1, i2 interface{}) bool

//Sort ...
func (by By) Sort(s []interface{}) {
	sw := &sortWrapper{
		items: s,
		by:    by,
	}
	sort.Sort(sw)
}

type sortWrapper struct {
	items []interface{}
	by    func(i1, i2 interface{}) bool
}

func (s *sortWrapper) Len() int {
	return len(s.items)
}

func (s *sortWrapper) Swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
}

func (s *sortWrapper) Less(i, j int) bool {
	return s.by(s.items[i], s.items[j])
}
