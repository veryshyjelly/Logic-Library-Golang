package main

import (
	. "Knowledge/logic"
	"Knowledge/modelCheck"
	"fmt"
)

func main() {
	mustard := Symbol("ColMustard")
	plum := Symbol("ProfPlum")
	scarlet := Symbol("MrsScarlet")
	characters := []Sentence{mustard, plum, scarlet}

	ballroom := Symbol("ballroom")
	kitchen := Symbol("kitchen")
	library := Symbol("library")
	rooms := []Sentence{ballroom, kitchen, library}

	knife := Symbol("knife")
	revolver := Symbol("revolver")
	wrench := Symbol("wrench")
	weapons := []Sentence{knife, revolver, wrench}

	var symbols []Sentence
	symbols = append(symbols, characters...)
	symbols = append(symbols, rooms...)
	symbols = append(symbols, weapons...)

	knowledge := And(
		Or(mustard, plum, scarlet),
		Or(ballroom, kitchen, library),
		Or(knife, revolver, wrench),
	)
	knowledge.Add(And(
		Not(mustard), Not(kitchen), Not(revolver),
	))
	knowledge.Add(Or(
		Not(scarlet), Not(library), Not(wrench),
	))
	knowledge.Add(Not(plum))
	knowledge.Add(Not(ballroom))

	checkKnowledge(knowledge, symbols)
}

func checkKnowledge(knowledge Sentence, symbols []Sentence) {
	for _, s := range symbols {
		if modelCheck.ModelCheck(knowledge, s) {
			fmt.Printf("\033[32m%v: YES\n\033[0m", s)
		} else if !modelCheck.ModelCheck(knowledge, Not(s)) {
			fmt.Printf("%v: MAYBE\n", s)
		}
	}
}
