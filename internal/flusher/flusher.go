package flusher

import (
	"context"
	"fmt"
	"github.com/ozonva/ova-game-api/internal/repo"
	"github.com/ozonva/ova-game-api/internal/utils"
	"github.com/ozonva/ova-game-api/pkg/game"
)

type flusher struct {
	chunkSize  uint
	repository repo.HeroRepo
}

// Flusher - интерфейс для сброса задач в хранилище
type Flusher interface {
	Flush(ctx context.Context, heroes []game.Hero) []game.Hero
}

// NewFlusher возвращает Flusher с поддержкой батчевого сохранения
func NewFlusher(chunkSize uint, repository repo.HeroRepo) Flusher {
	return &flusher{
		chunkSize:  chunkSize,
		repository: repository,
	}
}

func (f *flusher) Flush(ctx context.Context, heroes []game.Hero) []game.Hero {
	var result []game.Hero

	chunks, err := utils.HeroesToChunks(heroes, int(f.chunkSize))
	if err != nil {
		return heroes
	}

	for _, chunk := range chunks {
		if err := f.repository.AddHeroes(ctx, chunk); err != nil {
			fmt.Printf("flush error: %s\n", err)
			result = append(result, chunk...)
		}
	}

	return result
}
