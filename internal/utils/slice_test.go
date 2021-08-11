package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceChunksMore(t *testing.T) {
	assert := assert.New(t)

	source := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := SliceChunks(source, 4)
	assert.Len(result, 3)

	expected := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9},
	}

	assert.Equal(expected, result)
}

func TestSliceChunksEquals(t *testing.T) {
	assert := assert.New(t)

	source := []int{1, 2, 3, 4}
	result := SliceChunks(source, 4)
	assert.Len(result, 1)

	expected := [][]int{
		{1, 2, 3, 4},
	}

	assert.Equal(expected, result)
}

func TestSliceChunksLess(t *testing.T) {
	assert := assert.New(t)

	source := []int{1, 2, 3}
	result := SliceChunks(source, 4)
	assert.Len(result, 1)

	expected := [][]int{
		{1, 2, 3},
	}

	assert.Equal(expected, result)
}

func TestSliceChunksEmpty(t *testing.T) {
	assert := assert.New(t)

	source := []int{}
	result := SliceChunks(source, 4)
	assert.Len(result, 0)

	expected := [][]int{}

	assert.Equal(expected, result)
}

func TestSliceChunksZeroChunkInteger(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic message: chunkSize cannot be less than or equal to zero!")
		}
	}()

	source := []int{1, 2, 3}
	SliceChunks(source, 0)
}

func TestSliceChunksLessZeroChunkInteger(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic message: chunkSize cannot be less than or equal to zero!")
		}
	}()

	source := []int{1, 2, 3}
	SliceChunks(source, -1)
}

func TestSliceDifferenceSlices(t *testing.T) {
	assert := assert.New(t)

	source := []int{1, 2, 3, 4, 5}
	comparable := []int{2, 3}
	expected := []int{1, 4, 5}
	result := SliceDifference(source, comparable)

	assert.Equal(expected, result)
}

func TestSliceDifferenceSlicesWithSourceDuplicates(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic message: Duplicate value in slice!")
		}
	}()

	source := []int{1, 2, 3, 4, 5, 5}
	comparable := []int{2, 3}
	SliceDifference(source, comparable)
}

func TestSliceDifferenceSlicesWithComparableCases(t *testing.T) {
	testCases := []struct {
		source     []int
		comparable []int
		expected   []int
	}{
		{
			source:     []int{1, 2, 3, 4, 5},
			comparable: []int{2, 3, 3},
			expected:   []int{1, 4, 5},
		},
		{
			source:     []int{1, 2, 3, 3, 3, 4, 5},
			comparable: []int{2, 3, 3},
			expected:   []int{1, 4, 5},
		},
		{
			source:     []int{2, 3, 3},
			comparable: []int{2, 3, 3},
			expected:   []int{},
		},
		{
			source:     []int{2},
			comparable: []int{2, 3, 3},
			expected:   []int{},
		},
		{
			source:     []int{2},
			comparable: []int{},
			expected:   []int{2},
		},
		{
			source:     []int{},
			comparable: []int{},
			expected:   []int{},
		},
		{
			source:     nil,
			comparable: nil,
			expected:   []int{},
		},
	}

	for _, testCase := range testCases {
		result := SliceDifference(testCase.source, testCase.comparable)

		assert.Equal(t, testCase.expected, result)
	}
}

func TestSliceDifferenceHardcore(t *testing.T) {
	assert := assert.New(t)

	source := []int{1, 2, 3, 4, 5}
	result := SliceDifferenceHardcore(source)
	assert.Equal([]int{1, 4, 5}, result)

	source = []int{1, 4, 5, 6, 7}
	result = SliceDifferenceHardcore(source)
	assert.Equal([]int{1, 4, 5, 6, 7}, result)
}
