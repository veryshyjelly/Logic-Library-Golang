package main

import (
	. "Knowledge/logic"
	"Knowledge/modelCheck"
	"fmt"
)

func main() {
	people := []string{"Gilderoy", "Pomona", "Minerva", "Horace"}
	houses := []string{"Gryffindor", "Hufflepuff", "Ravenclaw", "Slytherin"}

	symbols := make([]Sentence, 0, len(people)*len(houses))
	knowledge := And()

	for _, person := range people {
		for _, house := range houses {
			symbols = append(symbols, Symbol(fmt.Sprintf("%v%v", person, house)))
		}
	}

	for _, person := range people {
		knowledge.Add(Or(
			Symbol(fmt.Sprintf("%vGryffindor", person)),
			Symbol(fmt.Sprintf("%vHufflepuff", person)),
			Symbol(fmt.Sprintf("%vRavenclaw", person)),
			Symbol(fmt.Sprintf("%vSlytherin", person)),
		))
	}

	for _, person := range people {
		for _, h1 := range houses {
			for _, h2 := range houses {
				if h1 != h2 {
					knowledge.Add(
						Implication(Symbol(fmt.Sprintf("%v%v", person, h1)), Not(Symbol(fmt.Sprintf("%v%v", person, h2)))),
					)
				}
			}
		}
	}

	for _, house := range houses {
		for _, p1 := range people {
			for _, p2 := range people {
				if p1 != p2 {
					knowledge.Add(
						Implication(Symbol(fmt.Sprintf("%v%v", p1, house)), Not(Symbol(fmt.Sprintf("%v%v", p2, house)))),
					)
				}
			}
		}
	}

	knowledge.Add(
		Or(Symbol("GilderoyGryffindor"), Symbol("GilderoyRavenclaw")),
	)

	knowledge.Add(
		Not(Symbol("PomonaSlytherin")),
	)

	knowledge.Add(
		Symbol("MinervaGryffindor"),
	)

	for _, s := range symbols {
		if modelCheck.ModelCheck(knowledge, s) {
			fmt.Println(s)
		}
	}
}
