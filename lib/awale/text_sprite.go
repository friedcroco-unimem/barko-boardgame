package awale

import "barkdo/lib/barko"

const (
	yourTurn = iota
	oppoturn
	waiting
	win
	lose
	draw
)

func NewTextSprite(text int, s *gameScene, finalFunc barko.FuncCall) *barko.BaseSprite {
	width, height := barko.GetGameManager().GetGameLayout()

	str := "waiting"
	if text == yourTurn {
		str = "yourturn"
	} else if text == oppoturn {
		str = "oppoturn"
	} else if text == win {
		str = "win"
	} else if text == lose {
		str = "lose"
	} else if text == draw {
		str = "draw"
	}

	sprite := barko.NewSprite(str)
	sprite.SetAnchor(0.5, 0.5)
	sprite.SetPosition(float32(width/2), float32(height/2))
	sprite.SetOpacity(0)

	sprite.AddAction(barko.NewSequenceWithActions([]barko.Action{
		barko.NewFadeIn(0.4),
		barko.NewWait(0.5),
		barko.NewFadeOut(0.4),
		finalFunc,
		barko.NewDisappear(),
	}))

	return sprite
}
