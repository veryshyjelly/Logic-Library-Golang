package modelCheck

import (
	"Knowledge/logic"
)

func ModelCheck(knowledge, query logic.Sentence) bool {
	symbols := make(map[string]bool)
	for k, v := range knowledge.Symbols() {
		symbols[k] = v
	}
	for k, v := range query.Symbols() {
		symbols[k] = v
	}

	return checkAll(knowledge, query, symbols, map[string]bool{})
}

func checkAll(knowledge, query logic.Sentence, symbols, model map[string]bool) bool {
	if symbols == nil || len(symbols) == 0 {
		if knowledge.Evaluate(model) {
			return query.Evaluate(model)
		}
		return true
	} else {
		remaining := copyMap(symbols)
		p := pop(remaining)

		modelTrue := copyMap(model)
		modelTrue[p] = true

		modelFalse := copyMap(model)
		modelFalse[p] = false

		return checkAll(knowledge, query, remaining, modelTrue) && checkAll(knowledge, query, remaining, modelFalse)
	}
}

func copyMap(main map[string]bool) map[string]bool {
	res := make(map[string]bool)
	for k, v := range main {
		res[k] = v
	}
	return res
}

func pop(main map[string]bool) string {
	for k := range main {
		delete(main, k)
		return k
	}
	return ""
}
