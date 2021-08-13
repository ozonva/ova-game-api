package game

import (
	"fmt"
	"time"
)

type Hero struct {
	HeroID      uint64
	UserID      uint64
	Type        TypeHero
	Name        string
	Description string
	StartTime   time.Time
	EndTime     time.Time
}

func (self Hero) String() string {
	return fmt.Sprintf(
		"Hero(id=%d,user=%d,type=%s,name=%s)",
		self.HeroID,
		self.UserID,
		self.Type.String(),
		self.Name,
	)
}

func Create(userId uint64, heroId uint64, typeHero TypeHero) Hero {
	return Hero{
		UserID: userId,
		HeroID: heroId,
		Type:   typeHero,
		Name:   fmt.Sprintf("Hero №%d", heroId),
	}
}
