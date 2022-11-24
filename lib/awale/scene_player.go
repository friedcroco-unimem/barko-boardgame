package awale

import (
	"barkdo/lib/barko"
	"errors"
)

type menuPlayerScene struct {
	*barko.BaseScene
}

func NewMenuPlayerScene() *menuPlayerScene {
	return &menuPlayerScene{barko.NewBaseScene()}
}

func (s *menuPlayerScene) Init() {
	manager := barko.GetGameManager()
	screenWidth, screenHeight := manager.GetGameLayout()

	// draw background
	background := barko.NewSprite("background")
	s.AddChild(background)

	// draw options
	createMode := barko.NewSprite("create_pin")
	createMode.SetAnchor(0.5, 0.5)
	createMode.SetPosition(float32(screenWidth/2), float32(screenHeight/2-40))
	s.AddChild(createMode)

	clickCreate := barko.NewMouseOverNodeListener(barko.MouseLeft, createMode)
	clickCreate.SetOnMouseDown(func(float32, float32) {
		pin, err := createNewPin()
		if err != nil {
			panic(errors.New("cannot create pin"))
		}

		manager := barko.GetGameManager()
		manager.AddScene("waiting", NewMenuNewPinScene(pin))
		manager.ChangeScene("waiting")
	})
	s.AddEvent(clickCreate)

	joinMode := barko.NewSprite("enter_code")
	joinMode.SetAnchor(0.5, 0.5)
	joinMode.SetPosition(float32(screenWidth/2), float32(screenHeight/2+40))
	s.AddChild(joinMode)

	clickJoin := barko.NewMouseOverNodeListener(barko.MouseLeft, joinMode)
	clickJoin.SetOnMouseDown(func(float32, float32) {
		manager := barko.GetGameManager()
		manager.AddScene("join_pin", NewMenuJoinPinScene())
		manager.ChangeScene("join_pin")
	})
	s.AddEvent(clickJoin)
}
