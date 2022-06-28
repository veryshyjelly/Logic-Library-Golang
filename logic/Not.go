package logic

import (
	"fmt"
)

type not struct {
	operand Sentence
}

func Not(operand Sentence) *not {
	self := new(not)
	self.operand = operand
	return self
}

func (self not) Eq(other interface{}) bool {
	switch other.(type) {
	case not:
		return self.operand.Eq(other.(not).operand)
	case *not:
		return self.operand.Eq(other.(*not).operand)
	default:
		return false
	}
}

func (self not) String() string {
	return fmt.Sprintf("Not(%v)", self.operand)
}

func (self not) Evaluate(model map[string]bool) bool {
	return !self.operand.Evaluate(model)
}

func (self not) Formula() string {
	return "¬" + Parenthesize(self.operand.Formula())
}

func (self not) Symbols() map[string]bool {
	return self.operand.Symbols()
}
