package main

import (
	"unsafe"
)

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		if len(name) > maxPersonName {
			return
		}

		for i, char := range unsafe.Slice(unsafe.StringData(name), len(name)) {
			person.name[i] = char
		}
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
		person.respStr |= byte(respect)
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.respStr |= byte(strength) << 4
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.expLvl |= byte(experience)
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.expLvl |= byte(level) << 4
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.enumsVars |= 1 << 2
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.enumsVars |= 1 << 3
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.enumsVars |= 1 << 4
	}
}

func WithType(personType int) func(*GamePerson) {
	if personType != BuilderGamePersonType && personType != BlacksmithGamePersonType && personType != WarriorGamePersonType {
		return nil
	}

	return func(person *GamePerson) {
		person.enumsVars |= byte(personType)
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

const maxPersonName = 42

type GamePerson struct {
	// [-2,147,483,648, 2,147,483,647]
	x    int32
	y    int32
	z    int32
	gold uint32
	mana uint16
	name [maxPersonName]byte
	// [1-4] bits - respect
	// [5-8] bits - strength
	respStr byte
	// [1-4] bits - experience
	// [5-8] bits - level
	expLvl byte
	// [1-2] bits - personType (строитель/кузнец/воин)
	// [3] bit - has house
	// [4] bit - has gun
	// [5] bit - has family
	enumsVars byte // 00011110
}

func NewGamePerson(options ...Option) GamePerson {
	var gamePerson GamePerson
	for _, opt := range options {
		opt(&gamePerson)
	}

	return gamePerson
}

func (p *GamePerson) Name() string {
	for i, char := range p.name {
		if char == 0 {
			return unsafe.String(unsafe.SliceData(p.name[:i]), i+1)
		}

		if i == len(p.name)-1 {
			return unsafe.String(unsafe.SliceData(p.name[:i]), i+1)
		}
	}

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
	return int(p.respStr << 4 >> 4)
}

func (p *GamePerson) Strength() int {
	return int(p.respStr >> 4)
}

func (p *GamePerson) Experience() int {
	return int(p.expLvl << 4 >> 4)
}

func (p *GamePerson) Level() int {
	return int(p.expLvl >> 4)
}

func (p *GamePerson) HasHouse() bool {
	return p.enumsVars>>2&1 == 1
}

func (p *GamePerson) HasGun() bool {
	return p.enumsVars>>3&1 == 1
}

func (p *GamePerson) HasFamily() bool {
	return p.enumsVars>>4&1 == 1
}

func (p *GamePerson) Type() int {
	return int(p.enumsVars << 6 >> 6)
}
