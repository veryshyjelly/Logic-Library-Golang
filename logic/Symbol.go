package logic

import (
	"log"
)

type symbol struct {
	name string
}

func Symbol(name string) *symbol {
	self := new(symbol)
	self.name = name
	return self
}

func (self symbol) Eq(other symbol) bool {
	return self.name == other.name
}

func (self symbol) String() string {
	return self.name
}

func (self symbol) Evaluate(model map[string]bool) bool {
	if v, ok := model[self.name]; ok {
		return v
	} else {
		log.Fatalf("variable %s not in model\n", self.name)
	}
	return false
}

func (self symbol) Formula() string {
	return self.name
}

func (self symbol) Symbols() map[string]bool {
	return map[string]bool{self.name: true}
}
