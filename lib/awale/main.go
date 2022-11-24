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
	manager.PreloadImage("connect", "assets/menu/connect.png")
	manager.PreloadImage("continue", "assets/menu/continue.png")
	manager.PreloadImage("create_pin", "assets/menu/create_pin.png")
	manager.PreloadImage("easy", "assets/menu/easy.png")
	manager.PreloadImage("enter_code", "assets/menu/enter_code.png")
	manager.PreloadImage("hard", "assets/menu/hard.png")
	manager.PreloadImage("medium", "assets/menu/medium.png")
	manager.PreloadImage("new_game", "assets/menu/new_game.png")
	manager.PreloadImage("with_computer", "assets/menu/with_computer.png")
	manager.PreloadImage("with_player", "assets/menu/with_player.png")
}

func setDefaultScene() {
	manager := barko.GetGameManager()
	manager.AddScene("menu", NewMenuScene())
	manager.ChangeScene("menu")
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
