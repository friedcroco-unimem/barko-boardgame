package awale

import "barkdo/lib/barko"

type menuScene struct {
	*barko.BaseScene
}

func NewMenuScene() *menuScene {
	return &menuScene{barko.NewBaseScene()}
}

func (s *menuScene) Init() {
	manager := barko.GetGameManager()
	screenWidth, screenHeight := manager.GetGameLayout()

	// draw background
	background := barko.NewSprite("background")
	s.AddChild(background)

	// draw options
	withComputer := barko.NewSprite("with_computer")
	withComputer.SetAnchor(0.5, 0.5)
	// withComputer.SetScale(0.5, 0.5)
	withComputer.SetPosition(float32(screenWidth/2), float32(screenHeight/2-40))
	s.AddChild(withComputer)

	clickComputer := barko.NewMouseOverNodeListener(barko.MouseLeft, withComputer)
	clickComputer.SetOnMouseDown(func(float32, float32) {
		manager.AddScene("menu_computer", NewMenuComputerScene())
		manager.ChangeScene("menu_computer")
	})
	s.AddEvent(clickComputer)

	withPlayer := barko.NewSprite("with_player")
	withPlayer.SetAnchor(0.5, 0.5)
	// withPlayer.SetScale(0.5, 0.5)
	withPlayer.SetPosition(float32(screenWidth/2), float32(screenHeight/2+40))
	s.AddChild(withPlayer)

	clickPlayer := barko.NewMouseOverNodeListener(barko.MouseLeft, withPlayer)
	clickPlayer.SetOnMouseDown(func(float32, float32) {
		manager.AddScene("menu_player", NewMenuPlayerScene())
		manager.ChangeScene("menu_player")
	})
	s.AddEvent(clickPlayer)
}
