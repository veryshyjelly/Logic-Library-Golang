package logic

func Union(a, b map[string]bool) map[string]bool {
	result := make(map[string]bool)
	for k, v := range a {
		result[k] = v
	}
	for k, v := range b {
		result[k] = v
	}
	return result
}
