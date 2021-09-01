package game

type TypeHero int

const (
	Magician TypeHero = iota
	Fighter
	Shooter
)

func GetTypeHeroesEnums() [3]TypeHero {
	return [...]TypeHero{Magician, Fighter, Shooter}
}

func SearchTypeHeroesEnums(name string) TypeHero {
	for _, value := range GetTypeHeroesEnums() {
		if name == value.String() {
			return value
		}
	}

	return GetTypeHeroesEnums()[0]
}

func (self TypeHero) String() string {
	return [...]string{"Magician", "Fighter", "Shooter"}[self]
}

func (self TypeHero) EnumIndex() int {
	return int(self)
}
