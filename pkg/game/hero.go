package game

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Hero struct {
	ID          uuid.UUID `json:"id" db:"id"`
	UserID      uint64    `json:"user_id" db:"user_id"`
	TypeHero    TypeHero  `json:"type_hero" db:"type_hero"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"created_at"`
}

func (h Hero) String() string {
	return fmt.Sprintf(
		"Hero(id=%d,user=%d,type=%s,name=%s)",
		h.ID,
		h.UserID,
		h.TypeHero.String(),
		h.Name,
	)
}

func (h Hero) Id() uuid.UUID {
	return h.ID
}

func (h *Hero) GenerateId() {
	h.ID = uuid.New()
}

func (h *Hero) GenerateName() {
	h.Name = fmt.Sprintf("Hero №%s", h.ID.String())
}

func (h *Hero) SetDescription(description string) {
	h.Description = description
}

func NewHero(userId uint64, typeHero TypeHero, name string) Hero {
	hero := Hero{
		UserID:   userId,
		TypeHero: typeHero,
	}
	hero.GenerateId()
	hero.Name = name

	return hero
}
