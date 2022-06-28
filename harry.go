package main

import (
	. "Knowledge/logic"
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
