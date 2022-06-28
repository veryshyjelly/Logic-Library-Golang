package logic

import (
	"Knowledge/logic/sentence"
	"fmt"
	"strings"
)

type and struct {
	conjuncts []sentence.Sentence
}

func And(conjuncts ...sentence.Sentence) *and {
	self := new(and)
	self.conjuncts = conjuncts
	return self
}

func (self and) eq(other and) bool {
	if len(self.conjuncts) != len(other.conjuncts) {
		return false
	}
	for k, v := range self.conjuncts {
		if v != other.conjuncts[k] {
			return false
		}
	}
	return true
}

func (self and) String() string {
	elems := make([]string, 0)
	for _, conjunct := range self.conjuncts {
		elems = append(elems, fmt.Sprint(conjunct))
	}
	return fmt.Sprintf("Add(%v)", strings.Join(elems, ", "))
}

func (self *and) Add(conjunct sentence.Sentence) {
	self.conjuncts = append(self.conjuncts, conjunct)
}

func (self *and) Evaluate(model map[string]bool) bool {
	res := true
	for _, conjunct := range self.conjuncts {
		res = res && conjunct.Evaluate(model)
	}
	return res
}

func (self and) Formula() string {
	if len(self.conjuncts) == 1 {
		return self.conjuncts[0].Formula()
	}
	elems := make([]string, 0)
	for _, conjunct := range self.conjuncts {
		elems = append(elems, sentence.Parenthesize(conjunct.Formula()))
	}
	return strings.Join(elems, " ∧ ")
}

func (self and) Symbols() map[string]bool {
	res := make(map[string]bool)
	for _, conjunct := range self.conjuncts {
		for k, v := range conjunct.Symbols() {
			res[k] = v
		}
	}
	return res
}
