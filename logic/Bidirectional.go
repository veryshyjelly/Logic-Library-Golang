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

func (self bidirectional) Eq(other bidirectional) bool {
	return self.left == other.left && self.right == other.right
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
	res := make(map[string]bool)
	for k, v := range self.left.Symbols() {
		res[k] = v
	}
	for k, v := range self.right.Symbols() {
		res[k] = v
	}
	return res
}
