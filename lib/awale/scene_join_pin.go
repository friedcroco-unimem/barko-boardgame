package awale

import "barkdo/lib/barko"

type menuJoinPinScene struct {
	*barko.BaseScene
	pin *digitGroup
}

func NewMenuJoinPinScene() *menuJoinPinScene {
	return &menuJoinPinScene{barko.NewBaseScene(), nil}
}

func (s *menuJoinPinScene) Init() {
	manager := barko.GetGameManager()
	screenWidth, screenHeight := manager.GetGameLayout()

	// draw background
	background := barko.NewSprite("background")
	s.AddChild(background)

	// draw options
	waiting := barko.NewSprite("enter_code")
	waiting.SetAnchor(0.5, 0.5)
	waiting.SetPosition(float32(screenWidth/2), float32(screenHeight/2-40))
	s.AddChild(waiting)

	s.pin = NewDigitGroup(s, float32(screenWidth/2), float32(screenHeight/2+40))

	// event
	for i := barko.KeyDigit0; i <= barko.KeyDigit9; i++ {
		digit := int(i - barko.KeyDigit0)
		keyDown := barko.NewKeyboardListener(i)
		keyDown.SetOnKeyDown(func() {
			s.pin.ChangeValue(s.pin.GetValue()*10 + digit)
		})
		s.AddEvent(keyDown)
	}

	keyDown := barko.NewKeyboardListener(barko.KeyBackspace)
	keyDown.SetOnKeyDown(func() {
		s.pin.ChangeValue(s.pin.GetValue() / 10)
	})
	s.AddEvent(keyDown)
}

func (s *menuJoinPinScene) Update(dt float32) {
	s.BaseScene.Update(dt)

	pin := s.pin.GetValue()
	if pin >= 100000 {
		yourTurn, err := joinPin(pin)
		if err != nil {
			panic(err)
		}

		manager := barko.GetGameManager()
		manager.AddScene("network_game", NewGameScene(yourTurn, false, 0))
		manager.ChangeScene("network_game")
	}
}
