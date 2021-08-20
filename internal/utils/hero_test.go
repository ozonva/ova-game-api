package utils

import (
	"github.com/google/uuid"
	"github.com/ozonva/ova-game-api/pkg/game"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createHeroes(userId uint64, count int) []game.Hero {
	list := make([]game.Hero, count)

	for index := range list {
		var typeHero game.TypeHero = game.GetTypeHeroesEnums()[index%3]
		list[index] = game.NewHero(userId, typeHero)
	}

	return list
}

func TestHeroesSplitToBulks(t *testing.T) {
	assert := assert.New(t)

	var userId uint64 = 1
	heroes := createHeroes(userId, 5)
	expected := make(map[uuid.UUID]game.Hero)
	for _, hero := range heroes {
		if _, ok := expected[hero.ID]; ok {
			continue
		}
		expected[hero.ID] = hero
	}

	result, err := HeroesSplitToBulks(heroes)
	assert.Equal(true, err == nil)
	assert.Equal(expected, result)
}

func TestHeroesSplitToBulksEmpty(t *testing.T) {
	assert := assert.New(t)

	heroes := createHeroes(0, 0)
	expected := make(map[uuid.UUID]game.Hero)

	result, err := HeroesSplitToBulks(heroes)
	assert.Nil(err)
	assert.Equal(expected, result)
}

func TestHeroesSplitToBulksDuplicate(t *testing.T) {
	assert := assert.New(t)
	list := make([]game.Hero, 2)

	var iterate uint64
	var typeHero game.TypeHero = game.GetTypeHeroesEnums()[0]
	hero := game.NewHero(1, typeHero)
	for iterate = 0; iterate < 2; iterate++ {
		list[iterate] = hero
	}

	result, err := HeroesSplitToBulks(list)
	assert.NotNil(err)
	assert.Nil(result)
}

func TestHeroesSplitToBulksForTypes(t *testing.T) {
	assert := assert.New(t)

	var userId uint64 = 1
	heroes := createHeroes(userId, 5)
	expected := make(map[game.TypeHero][]game.Hero)

	for _, hero := range heroes {
		expected[hero.Type] = append(expected[hero.Type], hero)
	}

	result, err := HeroesSplitToBulksForTypes(heroes)
	assert.Equal(true, err == nil)
	assert.Equal(expected, result)
}

func TestHeroesSplitToBulksForTypesEmpty(t *testing.T) {
	assert := assert.New(t)

	heroes := createHeroes(0, 0)
	expected := make(map[game.TypeHero][]game.Hero)

	result, err := HeroesSplitToBulksForTypes(heroes)
	assert.Nil(err)
	assert.Equal(expected, result)
}

func TestHeroesSplitToBulksForTypesDuplicate(t *testing.T) {
	assert := assert.New(t)
	list := make([]game.Hero, 2)

	var iterate uint64
	var typeHero game.TypeHero = game.GetTypeHeroesEnums()[0]
	hero := game.NewHero(1, typeHero)
	for iterate = 0; iterate < 2; iterate++ {
		list[iterate] = hero
	}

	result, err := HeroesSplitToBulksForTypes(list)
	assert.NotNil(err)
	assert.Nil(result)
}

func TestHeroesToChunksMore(t *testing.T) {
	assert := assert.New(t)

	source := createHeroes(1, 3)
	result, err := HeroesToChunks(source, 2)
	assert.Nil(err)
	assert.Len(result, 2)

	for index, chunk := range result {
		assert.Equal(2-index, len(chunk))
	}
}

func TestHeroesToChunksEquals(t *testing.T) {
	assert := assert.New(t)

	source := createHeroes(1, 2)
	result, err := HeroesToChunks(source, 2)
	assert.Nil(err)
	assert.Len(result, 1)

	for index, chunk := range result {
		assert.Equal(2-index, len(chunk))
	}
}

func TestHeroesToChunksLess(t *testing.T) {
	assert := assert.New(t)

	source := createHeroes(1, 1)
	result, err := HeroesToChunks(source, 2)
	assert.Nil(err)
	assert.Len(result, 1)

	for _, chunk := range result {
		assert.Equal(1, len(chunk))
	}
}

func TestHeroesToChunksEmpty(t *testing.T) {
	assert := assert.New(t)

	source := make([]game.Hero, 0)
	result, err := HeroesToChunks(source, 2)
	assert.Nil(err)
	assert.Len(result, 0)

	expected := [][]game.Hero{}

	assert.Equal(expected, result)
}

func TestHeroesToChunksZeroChunkInteger(t *testing.T) {
	assert := assert.New(t)

	source := createHeroes(1, 1)
	result, err := HeroesToChunks(source, 0)
	assert.NotNil(err)
	assert.Nil(result)
}

func TestHeroesToChunksLessZeroChunkInteger(t *testing.T) {
	assert := assert.New(t)

	source := createHeroes(1, 1)
	result, err := HeroesToChunks(source, -1)
	assert.NotNil(err)
	assert.Nil(result)
}
