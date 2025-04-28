package main

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		person.name = name
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.x = int32(x)
		person.y = int32(y)
		person.z = int32(z)
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.gold = uint32(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.mana = uint16(mana)
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.health = uint16(health)
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.respect = uint8(respect)
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.strength = uint8(strength)
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.experience = uint8(experience)
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.level = uint8(level)
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.hasHouse = true
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.hasGun = true
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.hasFamily = true
	}
}

func WithType(personType int) func(*GamePerson) {
	if personType != BuilderGamePersonType && personType != BlacksmithGamePersonType && personType != WarriorGamePersonType {
		panic("invalid person type")
	}
	
	return func(person *GamePerson) {
		person.personType = uint8(personType)
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

type GamePerson struct {
	name       string
	x          int32
	y          int32
	z          int32
	gold       uint32
	mana       uint16
	health     uint16
	respect    uint8
	strength   uint8
	experience uint8
	level      uint8
	personType uint8
	hasHouse   bool
	hasFamily  bool
	hasGun     bool
}

func NewGamePerson(options ...Option) GamePerson {
	var gamePerson GamePerson
	for _, opt := range options {
		opt(&gamePerson)
	}

	return gamePerson
}

func (p *GamePerson) Name() string {
	// need to implement
	return p.name
}

func (p *GamePerson) X() int {
	// need to implement
	return int(p.x)
}

func (p *GamePerson) Y() int {
	// need to implement
	return int(p.y)
}

func (p *GamePerson) Z() int {
	// need to implement
	return int(p.z)
}

func (p *GamePerson) Gold() int {
	// need to implement
	return int(p.gold)
}

func (p *GamePerson) Mana() int {
	// need to implement
	return int(p.mana)
}

func (p *GamePerson) Health() int {
	// need to implement
	return int(p.health)
}

func (p *GamePerson) Respect() int {
	// need to implement
	return int(p.respect)
}

func (p *GamePerson) Strength() int {
	// need to implement
	return int(p.strength)
}

func (p *GamePerson) Experience() int {
	// need to implement
	return int(p.experience)
}

func (p *GamePerson) Level() int {
	// need to implement
	return int(p.level)
}

func (p *GamePerson) HasHouse() bool {
	// need to implement
	return p.hasHouse
}

func (p *GamePerson) HasGun() bool {
	// need to implement
	return p.hasGun
}

func (p *GamePerson) HasFamilty() bool {
	// need to implement
	return p.hasFamily
}

func (p *GamePerson) Type() int {
	// need to implement
	return int(p.personType)
}
