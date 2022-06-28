package logic

import (
	"fmt"
)

type implication struct {
	antecedent Sentence
	consequent Sentence
}

func Implication(antecedent, consequent Sentence) *implication {
	self := new(implication)
	self.antecedent, self.consequent = antecedent, consequent
	return self
}

func (self implication) Eq(other implication) bool {
	return self.antecedent == self.antecedent && self.consequent == other.consequent
}

func (self implication) String() string {
	return fmt.Sprintf("Implication(%v, %v)", self.antecedent, self.consequent)
}

func (self implication) Evaluate(model map[string]bool) bool {
	return (!self.antecedent.Evaluate(model)) || self.consequent.Evaluate(model)
}

func (self implication) Formula() string {
	antecedent := Parenthesize(self.antecedent.Formula())
	consequent := Parenthesize(self.consequent.Formula())
	return antecedent + " => " + consequent
}

func (self implication) Symbols() map[string]bool {
	res := make(map[string]bool)
	for k, v := range self.antecedent.Symbols() {
		res[k] = v
	}
	for k, v := range self.consequent.Symbols() {
		res[k] = v
	}
	return res
}
