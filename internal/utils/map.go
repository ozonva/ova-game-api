package utils

func MapFlip(source map[string]int) map[int]string {
	list := make(map[int]string)
	for key, value := range source {
		list[value] = key
	}

	return list
}
