package goblog

import "sort"

//Item ...
type Item interface{}

//By ...
type By func(i1, i2 Item) bool

//Sort ...
func (by By) Sort(s []Item) {
	sw := &sortWrapper{
		items: s,
		by:    by,
	}
	sort.Sort(sw)
}

type sortWrapper struct {
	items []Item
	by    func(i1, i2 Item) bool
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
