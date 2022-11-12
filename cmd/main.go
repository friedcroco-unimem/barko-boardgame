package main

import (
	"barkdo/lib/barko"
)

func main() {
	barko.GetGameManager().SetScreenSize(640, 480)
	barko.GetGameManager().SetGameLayout(640, 480)

	barko.GetGameManager().PreloadImage("test0", "lmouse_00.png")
	barko.GetGameManager().PreloadImage("test1", "lmouse_01.png")
	barko.GetGameManager().PreloadImage("test2", "lmouse_02.png")
	barko.GetGameManager().PreloadImage("test3", "lmouse_03.png")
	barko.GetGameManager().PreloadImage("test4", "lmouse_04.png")
	barko.GetGameManager().PreloadImage("test5", "lmouse_05.png")

	scene := barko.NewBaseScene()
	barko.GetGameManager().AddScene("scene", scene)
	barko.GetGameManager().ChangeScene("scene")

	image := barko.NewSprite("test0")
	scene.AddChild(image)
	image.SetPosition(200, 200)

	image.AddAction(barko.NewRepeatForever(barko.NewAnimationWithFrames([]barko.AnimFrame{
		barko.NewAnimFrame("test0", 0.2),
		barko.NewAnimFrame("test1", 0.2),
		barko.NewAnimFrame("test2", 0.2),
		barko.NewAnimFrame("test3", 0.2),
		// barko.NewAnimFrame("test4", 0.2),
		// barko.NewAnimFrame("test5", 0.2),
	})))

	barko.GetGameManager().RunGame()
}
