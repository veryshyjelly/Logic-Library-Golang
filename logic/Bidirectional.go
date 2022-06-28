package logic

import (
	"fmt"
)

type bidirectional struct {
	left  Sentence
	right Sentence
}

func Bidirectional(left, right Sentence) *bidirectional {
	self := new(bidirectional)
	self.left, self.right = left, right
	return self
}

func (self bidirectional) Eq(other interface{}) bool {
	switch other.(type) {
	case bidirectional:
		return self.left.Eq(other.(bidirectional).left) && self.right.Eq(other.(bidirectional).right)
	case *bidirectional:
		return self.left.Eq(other.(*bidirectional).left) && self.right.Eq(other.(*bidirectional).right)
	default:
		return false
	}
}

func (self bidirectional) String() string {
	return fmt.Sprintf("Biconditional(%v, %v)", self.left, self.right)
}

func (self bidirectional) Evaluate(model map[string]bool) bool {
	return (self.left.Evaluate(model) && self.right.Evaluate(model)) || (!self.left.Evaluate(model) && !self.right.Evaluate(model))
}

func (self bidirectional) Formula() string {
	left := Parenthesize(fmt.Sprint(self.left))
	right := Parenthesize(fmt.Sprint(self.right))
	return left + " <=> " + right
}

func (self bidirectional) Symbols() map[string]bool {
	return Union(self.left.Symbols(), self.right.Symbols())
}
