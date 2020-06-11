package cs112

import "fmt"

// Integer Sets

type IntSet map[int]struct{}

func (s IntSet) Add(v int) {
	s[v] = struct{}{}
}

func (s IntSet) Remove(v int) {
	delete(s, v)
}

func (s IntSet) Has(v int) bool {
	_, ok := s[v]; return ok
}

func (s IntSet) String() string {
    if len(s) == 0 {return "intSet{}"}
    result := "{"
    for key, _ := range s {result += fmt.Sprintf("%d", key) + ", "}
    return result[:len(result)-2] + "}"
}

func (s IntSet) LoadSlice(a []int) {
	for _, v := range a { s.Add(v) }
}

func (s IntSet) ToSlice() []int {
	result := make([]int, 0)
	for elem, _ := range s { result = append(result, elem) }
	return result
}

func (s1 IntSet) Union(s2 IntSet) IntSet {
	result := make(IntSet)
	for elem, _ := range s1 { result.Add(elem) }
	for elem, _ := range s2 { result.Add(elem) }
	return result
}

func IntSetFromSlice(a []int) IntSet {
	s := make(IntSet)
	s.LoadSlice(a)
	return s
}


// Integer Sets

type StringSet map[string]struct{}

func (s StringSet) Add(v string) {
	s[v] = struct{}{}
}

func (s StringSet) Remove(v string) {
	delete(s, v)
}

func (s StringSet) Has(v string) bool {
	_, ok := s[v]; return ok
}

func (s StringSet) String() string {
    if len(s) == 0 {return "stringSet{}"}
    result := "{"
    for key, _ := range s {result += key + ", "}
    return result[:len(result)-2] + "}"
}

func (s StringSet) LoadSlice(a []string) {
	for _, v := range a { s.Add(v) }
}

func (s StringSet) ToSlice() []string {
	result := make([]string, 0)
	for elem, _ := range s { result = append(result, elem) }
	return result
}

func (s1 StringSet) Union(s2 StringSet) StringSet {
	result := make(StringSet)
	for elem, _ := range s1 { result.Add(elem) }
	for elem, _ := range s2 { result.Add(elem) }
	return result
}

func StringSetFromSlice(a []string) StringSet {
	s := make(StringSet)
	s.LoadSlice(a)
	return s
}
