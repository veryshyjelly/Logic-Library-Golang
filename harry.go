package main

import (
	. "Knowledge/logic"
	. "Knowledge/modelCheck"
	"fmt"
)

func main() {
	rain := Symbol("rain")
	hagrid := Symbol("hagrid")
	dumbldore := Symbol("dumbledore")

	knowledge := And(
		Implication(Not(rain), hagrid),
		Or(hagrid, dumbldore),
		Not(And(hagrid, dumbldore)),
		dumbldore,
	)

	fmt.Println(ModelCheck(knowledge, rain))
}
