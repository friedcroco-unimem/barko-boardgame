package awale

import (
	"barkdo/lib/barko"
)

func NewArrowSprite(squareIndex int, direction int, s *gameScene) *barko.BaseSprite {
	width, height := barko.GetGameManager().GetGameLayout()
	width /= 2
	height /= 2

	squareCenterX, _ := getSquareCenter(squareIndex)
	width += int(squareCenterX)

	var unit *barko.BaseSprite = nil
	if (squareIndex < 6 && direction == 1) || (squareIndex > 6 && direction == -1) {
		unit = barko.NewSpriteWithRect("arrow", barko.NewRect(0, 0, 24, 15))
		width += 20
	} else {
		unit = barko.NewSpriteWithRect("arrow", barko.NewRect(24, 0, 24, 15))
		width -= 20
	}
	unit.SetAnchor(0.5, 0.5)

	if squareIndex < 6 {
		height -= 140
	} else {
		height += 140
	}

	unit.SetPosition(float32(width), float32(height))
	click := barko.NewMouseOverNodeListener(barko.MouseLeft, unit)
	click.SetOnMouseDown(func(float32, float32) {
		if unit.IsRemoved() {
			click.RemoveFromScene()
			return
		}

		s.setUpApplyMove(NewMove(squareIndex, direction))
		for _, arrow := range s.arrows {
			arrow.RemoveFromScene()
		}
	})
	s.AddEvent(click)
	return unit
}
