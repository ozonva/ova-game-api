package repo

import (
	"github.com/google/uuid"
	"github.com/ozonva/ova-game-api/pkg/game"
)

// HeroRepo - интерфейс хранилища для сущности Hero
type HeroRepo interface {
	AddHeroes(heroes []game.Hero) error
	ListHeroes(limit, offset uint64) ([]game.Hero, error)
	DescribeHeroes(heroId uuid.UUID) (*game.Hero, error)
}
