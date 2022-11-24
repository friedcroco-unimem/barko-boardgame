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
	_, saveExists := isSaveDataExist()
	newGameMode := barko.NewSprite("new_game")
	newGameMode.SetAnchor(0.5, 0.5)
	if saveExists {
		newGameMode.SetPosition(float32(screenWidth/2), float32(screenHeight/2-40))
	} else {
		newGameMode.SetPosition(float32(screenWidth/2), float32(screenHeight/2))
	}

	s.AddChild(newGameMode)

	clickNewGame := barko.NewMouseOverNodeListener(barko.MouseLeft, newGameMode)
	clickNewGame.SetOnMouseDown(func(float32, float32) {
		manager.AddScene("new_game_menu", NewMenuComputerNewGameScene())
		manager.ChangeScene("new_game_menu")
	})
	s.AddEvent(clickNewGame)

	if saveExists {
		continueMode := barko.NewSprite("continue")
		continueMode.SetAnchor(0.5, 0.5)
		continueMode.SetPosition(float32(screenWidth/2), float32(screenHeight/2+40))
		s.AddChild(continueMode)

		clickContinue := barko.NewMouseOverNodeListener(barko.MouseLeft, continueMode)
		clickContinue.SetOnMouseDown(func(float32, float32) {
			manager.AddScene("continue", NewGameSceneWithSaveData())
			manager.ChangeScene("continue")
		})
		s.AddEvent(clickContinue)
	}
}
