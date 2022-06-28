package main

import (
	. "Knowledge/logic"
	"Knowledge/modelCheck"
	"fmt"
)

func main() {
	rain := Symbol("rain")
	hagrid := Symbol("hagrid")
	dumbldore := Symbol("dumbledore")

	//mySentence := And(rain, hagrid)
	//fmt.Println(mySentence.Formula())

	knowledge := And(
		Implication(Not(rain), hagrid),
		Or(hagrid, dumbldore),
		Not(And(hagrid, dumbldore)),
		dumbldore,
	)

	one := Or(rain, hagrid)
	two := Or(rain, hagrid)

	fmt.Println(one.Eq(two))

	fmt.Println(knowledge.Formula())
	fmt.Println(modelCheck.ModelCheck(knowledge, rain))
}
