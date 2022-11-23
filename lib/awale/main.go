package awale

import "barkdo/lib/barko"

func setLayout() {
	manager := barko.GetGameManager()
	manager.SetScreenSize(900, 540)
	manager.SetGameLayout(600, 360)
}

func preloadAssets() {
	manager := barko.GetGameManager()
	manager.PreloadImage("background", "assets/background/background.jpg")
	manager.PreloadImage("foreground", "assets/background/foreground.png")
	manager.PreloadImage("normal_unit", "assets/unit/normal.png")
	manager.PreloadImage("boss_unit", "assets/unit/boss.png")
	manager.PreloadImage("arrow", "assets/background/arrow.png")
	manager.PreloadImage("yourturn", "assets/text/yourturn.png")
	manager.PreloadImage("oppoturn", "assets/text/oppoturn.png")
	manager.PreloadImage("waiting", "assets/text/waiting.png")
	manager.PreloadImage("win", "assets/text/win.png")
	manager.PreloadImage("lose", "assets/text/lose.png")
	manager.PreloadImage("draw", "assets/text/draw.png")
	manager.PreloadImage("digit", "assets/text/digit.png")
}

func setDefaultScene() {
	manager := barko.GetGameManager()
	manager.AddScene("game", NewGameScene(0, true))
	manager.ChangeScene("game")
}

func run() {
	manager := barko.GetGameManager()
	manager.RunGame()
}

func Run() {
	setLayout()
	preloadAssets()
	setDefaultScene()
	run()
}
