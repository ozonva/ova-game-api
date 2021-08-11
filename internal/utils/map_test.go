package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapFlipWithDuplicates(t *testing.T) {
	assert := assert.New(t)

	result := MapFlip(map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
		"key4": 4,
		"key5": 4,
		"key6": 4,
		"key7": 7,
	})

	assert.Len(result, 5)

	keys := make([]int, 0)
	for index := range result {
		keys = append(keys, index)
	}

	assert.Contains(keys, 1)
	assert.Contains(keys, 2)
	assert.Contains(keys, 3)
	assert.Contains(keys, 4)
	assert.Contains(keys, 7)
}

func TestMapFlipWithoutDuplicates(t *testing.T) {
	assert := assert.New(t)

	expected := map[int]string{
		1: "key1",
		2: "key2",
		3: "key3",
		4: "key4",
		5: "key5",
		6: "key6",
		7: "key7",
	}
	result := MapFlip(map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
		"key4": 4,
		"key5": 5,
		"key6": 6,
		"key7": 7,
	})

	assert.Equal(expected, result)
}
