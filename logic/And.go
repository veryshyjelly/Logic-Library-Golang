package logic

import (
	"fmt"
	"strings"
)

type and struct {
	conjuncts []Sentence
}

func And(conjuncts ...Sentence) *and {
	self := new(and)
	self.conjuncts = conjuncts
	return self
}

func (self and) Eq(other interface{}) bool {
	var x and
	switch other.(type) {
	case and:
		x = other.(and)
	case *and:
		x = *other.(*and)
	default:
		return false
	}
	if len(self.conjuncts) != len(x.conjuncts) {
		return false
	}
	for i, conjunct := range self.conjuncts {
		if !conjunct.Eq(x.conjuncts[i]) {
			return false
		}
	}
	return true
}

func (self and) String() string {
	result := fmt.Sprint(self.conjuncts)
	return fmt.Sprintf("And(%v)", result)[1 : len(result)-1]
}

func (self *and) Add(conjunct Sentence) {
	self.conjuncts = append(self.conjuncts, conjunct)
}

func (self and) Evaluate(model map[string]bool) bool {
	for _, conjunct := range self.conjuncts {
		if !conjunct.Evaluate(model) {
			return false
		}
	}
	return true
}

func (self and) Formula() string {
	if len(self.conjuncts) == 1 {
		return self.conjuncts[0].Formula()
	}
	elems := make([]string, 0)
	for _, conjunct := range self.conjuncts {
		elems = append(elems, Parenthesize(conjunct.Formula()))
	}
	return strings.Join(elems, " ∧ ")
}

func (self and) Symbols() map[string]bool {
	result := make(map[string]bool)
	for _, conjunct := range self.conjuncts {
		for s, value := range conjunct.Symbols() {
			result[s] = value
		}
	}
	return result
}
