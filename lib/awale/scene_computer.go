package awale

import "barkdo/lib/barko"

type menuComputerScene struct {
	*barko.BaseScene
}

func NewMenuComputerScene() *menuComputerScene {
	return &menuComputerScene{barko.NewBaseScene()}
}

func (s *menuComputerScene) Init() {
	manager := barko.GetGameManager()
	screenWidth, screenHeight := manager.GetGameLayout()

	// draw background
	background := barko.NewSprite("background")
	s.AddChild(background)

	// draw options
	newGameMode := barko.NewSprite("new_game")
	newGameMode.SetAnchor(0.5, 0.5)
	newGameMode.SetPosition(float32(screenWidth/2), float32(screenHeight/2-40))
	s.AddChild(newGameMode)

	clickNewGame := barko.NewMouseOverNodeListener(barko.MouseLeft, newGameMode)
	clickNewGame.SetOnMouseDown(func(float32, float32) {
		manager.AddScene("new_game_menu", NewMenuComputerNewGameScene())
		manager.ChangeScene("new_game_menu")
	})
	s.AddEvent(clickNewGame)

	continueMode := barko.NewSprite("continue")
	continueMode.SetAnchor(0.5, 0.5)
	continueMode.SetPosition(float32(screenWidth/2), float32(screenHeight/2+40))
	s.AddChild(continueMode)
}
