package utils

import (
	"fmt"
	"github.com/ozonva/ova-game-api/pkg/game"
)

func HeroesSplitToBulks(heroes []game.Hero) (map[uint64]game.Hero, error) {
	out := make(map[uint64]game.Hero)

	for _, hero := range heroes {
		if _, ok := out[hero.HeroID]; ok {
			return nil, fmt.Errorf("Duplicate id %d", hero.HeroID)
		}
		out[hero.HeroID] = hero
	}

	return out, nil
}

func HeroesSplitToBulksForTypes(heroes []game.Hero) (map[game.TypeHero][]game.Hero, error) {
	out := make(map[game.TypeHero][]game.Hero)
	hasMap := make(map[uint64]struct{})

	for _, hero := range heroes {
		if _, ok := hasMap[hero.HeroID]; ok {
			return nil, fmt.Errorf("Duplicate id %d", hero.HeroID)
		}
		hasMap[hero.HeroID] = struct{}{}
		out[hero.Type] = append(out[hero.Type], hero)
	}

	return out, nil
}

func HeroesToChunks(source []game.Hero, chunkSize int) ([][]game.Hero, error) {
	if chunkSize <= 0 {
		return nil, fmt.Errorf("chunkSize cannot be less than or equal to zero! Value chunkSize: %d", chunkSize)
	}

	chunks := make([][]game.Hero, 0)
	sourceLen := len(source)

	for i := 0; i < sourceLen; i += chunkSize {
		end := i + chunkSize

		if end > sourceLen {
			end = sourceLen
		}

		chunks = append(chunks, source[i:end])
	}

	return chunks, nil
}
