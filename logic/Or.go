package logic

import (
	"Knowledge/logic/sentence"
	"fmt"
	"strings"
)

type or struct {
	disjuncts []sentence.Sentence
}

func Or(disjuncts ...sentence.Sentence) *or {
	self := new(or)
	self.disjuncts = disjuncts
	return self
}

func (self or) Eq(other or) bool {
	if len(self.disjuncts) != len(other.disjuncts) {
		return false
	}
	for k, v := range self.disjuncts {
		if v != other.disjuncts[k] {
			return false
		}
	}
	return true
}

func (self or) String() string {
	elems := make([]string, 0)
	for _, v := range self.disjuncts {
		elems = append(elems, fmt.Sprint(v))
	}
	return fmt.Sprintf("Or(%v)", strings.Join(elems, ", "))
}

func (self or) Evaluate(model map[string]bool) bool {
	res := false
	for _, disjunct := range self.disjuncts {
		res = res || disjunct.Evaluate(model)
	}
	return res
}

func (self or) Formula() string {
	if len(self.disjuncts) == 1 {
		return self.disjuncts[0].Formula()
	}
	elems := make([]string, 0)
	for _, disjunct := range self.disjuncts {
		elems = append(elems, sentence.Parenthesize(disjunct.Formula()))
	}
	return strings.Join(elems, " ∨ ")
}

func (self or) Symbols() map[string]bool {
	res := make(map[string]bool)
	for _, disjunct := range self.disjuncts {
		for k, v := range disjunct.Symbols() {
			res[k] = v
		}
	}
	return res
}
