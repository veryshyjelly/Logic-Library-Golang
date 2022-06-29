package logic

import (
	"fmt"
	"strings"
)

type or struct {
	disjuncts []Sentence
}

func Or(disjuncts ...Sentence) *or {
	self := new(or)
	self.disjuncts = disjuncts
	return self
}

func (self or) Eq(other interface{}) bool {
	var x or
	switch other.(type) {
	case or:
		x = other.(or)
	case *or:
		x = *other.(*or)
	default:
		return false
	}
	if len(self.disjuncts) != len(x.disjuncts) {
		return false
	}
	for i, disjunct := range self.disjuncts {
		if !disjunct.Eq(x.disjuncts[i]) {
			return false
		}
	}
	return true
}

func (self or) String() string {
	result := make([]string, len(self.disjuncts))
	for k, disjunct := range self.disjuncts {
		result[k] = fmt.Sprint(disjunct)
	}
	return fmt.Sprintf("And(%v)", strings.Join(result, ", "))

}

func (self or) Evaluate(model map[string]bool) bool {
	for _, disjunct := range self.disjuncts {
		if disjunct.Evaluate(model) {
			return true
		}
	}
	return false
}

func (self or) Formula() string {
	if len(self.disjuncts) == 1 {
		return self.disjuncts[0].Formula()
	}
	elems := make([]string, 0)
	for _, disjunct := range self.disjuncts {
		elems = append(elems, Parenthesize(disjunct.Formula()))
	}
	return strings.Join(elems, " ∨ ")
}

func (self or) Symbols() map[string]bool {
	result := make(map[string]bool)
	for _, disjunct := range self.disjuncts {
		for s, value := range disjunct.Symbols() {
			result[s] = value
		}
	}
	return result
}
