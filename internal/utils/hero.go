package utils

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/ozonva/ova-game-api/pkg/game"
)

func HeroesSplitToBulks(heroes []game.Hero) (map[uuid.UUID]game.Hero, error) {
	out := make(map[uuid.UUID]game.Hero)

	for _, hero := range heroes {
		if _, ok := out[hero.ID]; ok {
			return nil, fmt.Errorf("duplicate id %d", hero.ID)
		}
		out[hero.ID] = hero
	}

	return out, nil
}

func HeroesSplitToBulksForTypes(heroes []game.Hero) (map[game.TypeHero][]game.Hero, error) {
	out := make(map[game.TypeHero][]game.Hero)
	hasMap := make(map[uuid.UUID]struct{})

	for _, hero := range heroes {
		if _, ok := hasMap[hero.ID]; ok {
			return nil, fmt.Errorf("duplicate id %d", hero.ID)
		}
		hasMap[hero.ID] = struct{}{}
		out[hero.TypeHero] = append(out[hero.TypeHero], hero)
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
