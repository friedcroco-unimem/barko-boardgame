package awale

import "fmt"

type Square interface {
	GetSize() int
	IsEmpty() bool
	IsMovable() bool
	IsPassable() bool
	PopUnit() Unit
	AddUnit(unit Unit)
	PopAllUnits() []Unit
	Clone() Square
}

type baseSquare struct {
	units []Unit
}

func NewBaseSquare() *baseSquare {
	return &baseSquare{make([]Unit, 0)}
}

func NewBaseSquareWithUnits(units []Unit) *baseSquare {
	return &baseSquare{units}
}

func (s *baseSquare) GetSize() int {
	return len(s.units)
}

func (s *baseSquare) IsEmpty() bool {
	return len(s.units) == 0
}

func (s *baseSquare) PopUnit() Unit {
	size := len(s.units)

	if size == 0 {
		panic(fmt.Errorf("pop from empty square"))
	}

	res := s.units[size-1]
	s.units = s.units[:size-1]
	return res
}

func (s *baseSquare) AddUnit(unit Unit) {
	s.units = append(s.units, unit)
}

func (s *baseSquare) PopAllUnits() []Unit {
	res := s.units
	s.units = make([]Unit, 0)
	return res
}

type normalSquare struct {
	*baseSquare
}

func NewNormalSquare() Square {
	return &normalSquare{NewBaseSquare()}
}

func NewNormalSquareDefault() Square {
	units := make([]Unit, 0)
	for i := 0; i < 5; i++ {
		units = append(units, normalUnitSample)
	}
	return &normalSquare{NewBaseSquareWithUnits(units)}
}

func NewNormalSquareWithCount(n int) Square {
	units := make([]Unit, 0)
	for i := 0; i < n; i++ {
		units = append(units, normalUnitSample)
	}
	return &normalSquare{NewBaseSquareWithUnits(units)}
}

func (s *normalSquare) IsMovable() bool {
	return true
}

func (s *normalSquare) IsPassable() bool {
	return true
}

func (s *normalSquare) Clone() Square {
	return &normalSquare{NewBaseSquareWithUnits(s.units)}
}

type bossSquare struct {
	*baseSquare
}

func NewBossSquare() Square {
	return &bossSquare{NewBaseSquare()}
}

func NewBossSquareDefault() Square {
	return &bossSquare{NewBaseSquareWithUnits([]Unit{bossUnitSample})}
}

func NewBossSquareWithCount(n int) Square {
	units := make([]Unit, 0)
	for i := 0; i < n; i++ {
		units = append(units, bossUnitSample)
	}
	return &bossSquare{NewBaseSquareWithUnits(units)}
}

func (s *bossSquare) IsMovable() bool {
	return false
}

func (s *bossSquare) IsPassable() bool {
	return false
}

func (s *bossSquare) Clone() Square {
	return &bossSquare{NewBaseSquareWithUnits(s.units)}
}
