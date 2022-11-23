package awale

import (
	"barkdo/lib/barko"
	"math/rand"
)

func NewNormalUnitSprite() *barko.BaseSprite {
	unit := barko.NewSpriteWithRect("normal_unit", barko.NewRect(0, 0, 15, 15))
	timeStep := rand.Float32()*0.02 + 0.06
	frame := make([]barko.AnimFrame, 0)
	for i := 0; i < 6; i++ {
		frame = append(frame,
			barko.NewAnimFrameWithRect("normal_unit", barko.NewRect(float32((i%4)*15), float32((i/4)*15), 15, 15), timeStep))
	}

	unit.AddAction(barko.NewRepeatForever(barko.NewAnimationWithFrames(frame)))
	return unit
}

func NewBossUnitSprite() *barko.BaseSprite {
	unit := barko.NewSpriteWithRect("boss_unit", barko.NewRect(0, 0, 40, 40))
	timeStep := rand.Float32()*0.04 + 0.08
	frame := make([]barko.AnimFrame, 0)
	for i := 0; i < 6; i++ {
		frame = append(frame,
			barko.NewAnimFrameWithRect("boss_unit", barko.NewRect(float32((i%4)*40), float32((i/4)*40), 40, 40), timeStep))
	}

	unit.AddAction(barko.NewRepeatForever(barko.NewAnimationWithFrames(frame)))
	return unit
}

func getNormalUnitColumn(i int) int {
	if i%3 == 0 {
		return 0
	}

	if i%3 == 1 {
		return -1
	}

	return 1
}

func getNormalUnitRow(i int) int {
	row := i / 3
	return ((row%2)*2 - 1) * ((row + 1) / 2)
}

func getBossUnitColumn(i int) int {
	if i == 0 {
		return 0
	}

	if (i-1)%3 == 0 {
		return 0
	}

	if (i-1)%3 == 1 {
		return -1
	}

	return 1
}

func getBossUnitRow(i int) int {
	row := (i + 2) / 3
	return ((row%2)*2 - 1) * ((row + 1) / 2)
}

func getSquareCenter(i int) (float32, float32) {
	if i == 0 {
		return -240, 0
	}

	if i == 6 {
		return 240, 0
	}

	if i <= 5 {
		return float32(-240 + 80*i), -60
	}

	return float32(240 - 80*(i-6)), 60
}

func getUnitAbsolutePosition(squareIndex int, unitIndex int) (float32, float32) {
	width, height := barko.GetGameManager().GetGameLayout()
	width /= 2
	height /= 2

	squareCenterX, squareCenterY := getSquareCenter(squareIndex)
	width += int(squareCenterX)
	height += int(squareCenterY)

	if squareIndex%6 == 0 {
		width += getBossUnitColumn(unitIndex) * 24
		heightRow := getBossUnitRow(unitIndex)
		if heightRow > 0 {
			height += 13
		} else if heightRow < 0 {
			height -= 13
		}

		height += heightRow * 24
		return float32(width), float32(height)
	}

	width += getNormalUnitColumn(unitIndex) * 24
	height += getNormalUnitRow(unitIndex) * 24
	return float32(width), float32(height)
}
