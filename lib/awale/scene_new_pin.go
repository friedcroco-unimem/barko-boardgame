package awale

import "barkdo/lib/barko"

type menuNewPinScene struct {
	*barko.BaseScene
	pin int
}

func NewMenuNewPinScene(pin int) *menuNewPinScene {
	return &menuNewPinScene{barko.NewBaseScene(), pin}
}

func (s *menuNewPinScene) Init() {
	manager := barko.GetGameManager()
	screenWidth, screenHeight := manager.GetGameLayout()

	// draw background
	background := barko.NewSprite("background")
	s.AddChild(background)

	// draw options
	waiting := barko.NewSprite("waiting")
	waiting.SetAnchor(0.5, 0.5)
	waiting.SetPosition(float32(screenWidth/2), float32(screenHeight/2-40))
	s.AddChild(waiting)

	pinCode := NewDigitGroup(s, float32(screenWidth/2), float32(screenHeight/2+40))
	pinCode.ChangeValue(s.pin)
}

func (s *menuNewPinScene) Update(dt float32) {
	s.BaseScene.Update(dt)

	msg := getNetworkMsg()
	if msg != nil {
		setNetworkMsg(nil)
		manager := barko.GetGameManager()
		manager.AddScene("network_game", NewGameScene(msg.YourTurn, false, 0))
		manager.ChangeScene("network_game")
	}
}
