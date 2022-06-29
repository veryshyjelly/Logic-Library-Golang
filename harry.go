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
	)

	knowledge.Add(Or(hagrid, dumbldore))
	knowledge.Add(Not(And(hagrid, dumbldore)))
	knowledge.Add(dumbldore)

	//
	//one := Or(rain, hagrid)
	//two := Or(rain, hagrid)
	//
	//fmt.Println(one.Eq(two))

	fmt.Println(knowledge.Formula())
	fmt.Println(modelCheck.ModelCheck(knowledge, rain))
	fmt.Println(modelCheck.ModelCheck(knowledge, hagrid))
	fmt.Println(modelCheck.ModelCheck(knowledge, dumbldore))
}
