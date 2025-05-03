package main

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
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

	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {

	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {

	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {

	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {

	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {

	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {

	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {

	}
}

func WithType(personType int) func(*GamePerson) {
	if personType != BuilderGamePersonType && personType != BlacksmithGamePersonType && personType != WarriorGamePersonType {
		panic("invalid person type")
	}

	return func(person *GamePerson) {
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

type GamePerson struct {
	// [-2,147,483,648, 2,147,483,647]
	x    int32
	y    int32
	z    int32
	gold uint32
	mana uint16
	name [42]byte
	// [1-4] bits - respect
	// [5-8] bits - strength
	respStr byte
	// [1-4] bits - experience
	// [5-8] bits - level
	expLvl byte
	// [1-2] bits - personType (строитель/кузнец/воин) 0,01,10
	// [3] bit - has house
	// [4] bit - has gun
	// [5] bit - has family
	enumsVars byte
}

func NewGamePerson(options ...Option) GamePerson {
	var gamePerson GamePerson
	for _, opt := range options {
		opt(&gamePerson)
	}

	return gamePerson
}

func (p *GamePerson) Name() string {
	return ""
}

func (p *GamePerson) X() int {
	return int(p.x)
}

func (p *GamePerson) Y() int {
	return int(p.y)
}

func (p *GamePerson) Z() int {
	return int(p.z)
}

func (p *GamePerson) Gold() int {
	return int(p.gold)
}

func (p *GamePerson) Mana() int {
	return int(p.mana)
}

func (p *GamePerson) Health() int {
	return 0
}

func (p *GamePerson) Respect() int {
	return 0
}

func (p *GamePerson) Strength() int {
	return 0
}

func (p *GamePerson) Experience() int {
	return 0
}

func (p *GamePerson) Level() int {
	return 0
}

func (p *GamePerson) HasHouse() bool {
	return true
}

func (p *GamePerson) HasGun() bool {
	return true
}

func (p *GamePerson) HasFamilty() bool {
	return true
}

func (p *GamePerson) Type() int {
	return 0
}
