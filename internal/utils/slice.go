package utils

func SliceChunks(source []int, chunkSize int) [][]int {
	if chunkSize <= 0 {
		panic("chunkSize cannot be less than or equal to zero!")
	}

	chunks := make([][]int, 0)
	sourceLen := len(source)

	for i := 0; i < sourceLen; i += chunkSize {
		end := i + chunkSize

		if end > sourceLen {
			end = sourceLen
		}

		chunks = append(chunks, source[i:end])
	}

	return chunks
}

func SliceDifferenceHardcore(source []int) []int {
	hardcore := []int{2, 3}

	return SliceDifference(source, hardcore)
}

func SliceDifference(source []int, comparable []int) []int {
	list := make(map[int]struct{})
	newSource := make([]int, 0)
	hasNewSource := make(map[int]struct{})

	for _, value := range comparable {
		list[value] = struct{}{}
	}

	for _, value := range source {
		if _, ok := list[value]; ok {
			continue
		}

		if _, ok := hasNewSource[value]; ok {
			panic("Duplicate value in slice!")
		}

		hasNewSource[value] = struct{}{}
		newSource = append(newSource, value)
	}

	return newSource
}
