package awale

import "barkdo/lib/barko"

type menuComputerNewGameScene struct {
	*barko.BaseScene
}

func NewMenuComputerNewGameScene() *menuComputerNewGameScene {
	return &menuComputerNewGameScene{barko.NewBaseScene()}
}

func (s *menuComputerNewGameScene) Init() {
	manager := barko.GetGameManager()
	screenWidth, screenHeight := manager.GetGameLayout()

	// draw background
	background := barko.NewSprite("background")
	s.AddChild(background)

	// draw options
	easyMode := barko.NewSprite("easy")
	easyMode.SetAnchor(0.5, 0.5)
	easyMode.SetPosition(float32(screenWidth/2), float32(screenHeight/2-40))
	s.AddChild(easyMode)

	clickEasy := barko.NewMouseOverNodeListener(barko.MouseLeft, easyMode)
	clickEasy.SetOnMouseDown(func(float32, float32) {
		manager.AddScene("easy_game", NewGameScene(0, true, easy))
		manager.ChangeScene("easy_game")
	})
	s.AddEvent(clickEasy)

	hardMode := barko.NewSprite("hard")
	hardMode.SetAnchor(0.5, 0.5)
	hardMode.SetPosition(float32(screenWidth/2), float32(screenHeight/2+40))
	s.AddChild(hardMode)

	clickHard := barko.NewMouseOverNodeListener(barko.MouseLeft, hardMode)
	clickHard.SetOnMouseDown(func(float32, float32) {
		manager.AddScene("hard_game", NewGameScene(0, true, hard))
		manager.ChangeScene("hard_game")
	})
	s.AddEvent(clickHard)
}
