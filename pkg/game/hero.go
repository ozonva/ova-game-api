package game

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Hero struct {
	ID          uuid.UUID `json:"id"`
	UserID      uint64    `json:"user_id"`
	Type        TypeHero  `json:"type"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (h Hero) String() string {
	return fmt.Sprintf(
		"Hero(id=%d,user=%d,type=%s,name=%s)",
		h.ID,
		h.UserID,
		h.Type.String(),
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

func NewHero(userId uint64, typeHero TypeHero) Hero {
	hero := Hero{
		UserID: userId,
		Type:   typeHero,
	}
	hero.GenerateId()
	hero.GenerateName()

	return hero
}
