package barko

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameManager interface {
	SetScreenSize(width int, height int)
	SetGameLayout(width int, height int)
	GetScreenSize() (int, int)
	GetGameLayout() (int, int)
	RunGame()
	SetTitle(title string)
	SetFPS(fps int)
	GetTimeStep() float32
	AddScene(key string, scene Scene)
	ChangeScene(key string)
	PreloadImage(key string, filename string)
	getImage(key string) *ebiten.Image
}

var gameManagerInstance GameManager = newGameManagerDefault()

type gameManager struct {
	screenWidth  int
	screenHeight int
	layoutWidth  int
	layoutHeight int
	game         *application
	fps          int
	dt           float32
	spritePool   *spritePool
}

func (g *gameManager) SetScreenSize(width int, height int) {
	ebiten.SetWindowSize(width, height)
	g.screenWidth = width
	g.screenHeight = height
}

func (g *gameManager) SetGameLayout(width int, height int) {
	g.layoutWidth = width
	g.layoutHeight = height
}

func (g *gameManager) GetScreenSize() (int, int) {
	return g.screenWidth, g.screenHeight
}

func (g *gameManager) GetGameLayout() (int, int) {
	return g.layoutWidth, g.layoutHeight
}

func (g *gameManager) RunGame() {
	if err := ebiten.RunGame(g.game); err != nil {
		log.Fatal(err)
	}
}

func (g *gameManager) SetTitle(title string) {
	ebiten.SetWindowTitle(title)
}

func (g *gameManager) SetFPS(fps int) {
	ebiten.SetTPS(fps)
	g.fps = fps
	g.dt = 1.0 / float32(fps)
}

func (g *gameManager) GetTimeStep() float32 {
	return g.dt
}

func (g *gameManager) AddScene(key string, scene Scene) {
	g.game.addScene(key, scene)
}

func (g *gameManager) ChangeScene(key string) {
	g.game.changeScene(key)
}

func (g *gameManager) PreloadImage(key string, filename string) {
	g.spritePool.preloadImage(key, filename)
}

func (g *gameManager) getImage(key string) *ebiten.Image {
	return g.spritePool.getImage(key)
}

func newGameManagerDefault() GameManager {
	manager := &gameManager{game: newApplication(), spritePool: newSpritePool()}
	manager.SetFPS(60)
	return manager
}

func GetGameManager() GameManager {
	return gameManagerInstance
}
