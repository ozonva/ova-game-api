package utils

import (
	"github.com/ozonva/ova-game-api/pkg/game"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createHeroes(userId uint64, count int) []game.Hero {
	list := make([]game.Hero, count)

	for index := range list {
		var typeHero game.TypeHero = game.GetTypeHeroesEnums()[index%3]
		list[index] = game.Create(userId, uint64(index+1), typeHero)
	}

	return list
}

func TestHeroesSplitToBulks(t *testing.T) {
	assert := assert.New(t)

	var userId uint64 = 1
	heroes := createHeroes(userId, 5)
	expected := make(map[uint64]game.Hero)
	var typeHero game.TypeHero = game.GetTypeHeroesEnums()[0]
	var typeHero2 game.TypeHero = game.GetTypeHeroesEnums()[1]
	var typeHero3 game.TypeHero = game.GetTypeHeroesEnums()[2]
	expected[1] = game.Create(userId, 1, typeHero)
	expected[2] = game.Create(userId, 2, typeHero2)
	expected[3] = game.Create(userId, 3, typeHero3)
	expected[4] = game.Create(userId, 4, typeHero)
	expected[5] = game.Create(userId, 5, typeHero2)

	result, err := HeroesSplitToBulks(heroes)
	assert.Equal(true, err == nil)
	assert.Equal(expected, result)
}

func TestHeroesSplitToBulksEmpty(t *testing.T) {
	assert := assert.New(t)

	heroes := createHeroes(0, 0)
	expected := make(map[uint64]game.Hero)

	result, err := HeroesSplitToBulks(heroes)
	assert.Nil(err)
	assert.Equal(expected, result)
}

func TestHeroesSplitToBulksDuplicate(t *testing.T) {
	assert := assert.New(t)
	list := make([]game.Hero, 2)

	var iterate uint64
	for iterate = 0; iterate < 2; iterate++ {
		var typeHero game.TypeHero = game.GetTypeHeroesEnums()[0]
		list[iterate] = game.Create(iterate+1, 1, typeHero)
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
	var typeHero game.TypeHero

	typeHero = game.GetTypeHeroesEnums()[0]
	expected[typeHero] = []game.Hero{
		game.Create(userId, 1, typeHero),
		game.Create(userId, 4, typeHero),
	}

	typeHero = game.GetTypeHeroesEnums()[1]
	expected[typeHero] = []game.Hero{
		game.Create(userId, 2, typeHero),
		game.Create(userId, 5, typeHero),
	}

	typeHero = game.GetTypeHeroesEnums()[2]
	expected[typeHero] = []game.Hero{
		game.Create(userId, 3, typeHero),
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
	for iterate = 0; iterate < 2; iterate++ {
		var typeHero game.TypeHero = game.GetTypeHeroesEnums()[0]
		list[iterate] = game.Create(iterate+1, 1, typeHero)
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

	var typeHero game.TypeHero = game.GetTypeHeroesEnums()[0]
	var typeHero2 game.TypeHero = game.GetTypeHeroesEnums()[1]
	var typeHero3 game.TypeHero = game.GetTypeHeroesEnums()[2]
	expected := [][]game.Hero{
		{game.Create(1, uint64(1), typeHero), game.Create(1, uint64(2), typeHero2)},
		{game.Create(1, uint64(3), typeHero3)},
	}

	assert.Equal(expected, result)
}

func TestHeroesToChunksEquals(t *testing.T) {
	assert := assert.New(t)

	source := createHeroes(1, 2)
	result, err := HeroesToChunks(source, 2)
	assert.Nil(err)
	assert.Len(result, 1)

	var typeHero game.TypeHero = game.GetTypeHeroesEnums()[0]
	var typeHero2 game.TypeHero = game.GetTypeHeroesEnums()[1]
	expected := [][]game.Hero{
		{game.Create(1, uint64(1), typeHero), game.Create(1, uint64(2), typeHero2)},
	}

	assert.Equal(expected, result)
}

func TestHeroesToChunksLess(t *testing.T) {
	assert := assert.New(t)

	source := createHeroes(1, 1)
	result, err := HeroesToChunks(source, 2)
	assert.Nil(err)
	assert.Len(result, 1)

	var typeHero game.TypeHero = game.GetTypeHeroesEnums()[0]
	expected := [][]game.Hero{
		{game.Create(1, uint64(1), typeHero)},
	}

	assert.Equal(expected, result)
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
