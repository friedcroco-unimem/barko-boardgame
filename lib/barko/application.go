package barko

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type application struct {
	curScene Scene
	scenes   map[string]Scene
}

func newApplication() *application {
	initScene := NewBaseScene()
	return &application{
		initScene,
		map[string]Scene{"": initScene},
	}
}

func (a *application) Update() error {
	dt := GetGameManager().GetTimeStep()
	a.update(dt)
	return nil
}

func (a *application) Draw(screen *ebiten.Image) {
	a.draw(screen)

	if isDebug {
		msg := fmt.Sprintf(`TPS: %0.2f - FPS: %0.2f`, ebiten.ActualTPS(), ebiten.ActualFPS())
		ebitenutil.DebugPrint(screen, msg)
	}
}

func (a *application) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return GetGameManager().GetGameLayout()
}

func (a *application) update(dt float32) {
	a.curScene.Update(dt)
}

func (a *application) draw(screen *ebiten.Image) {
	a.curScene.draw(screen)
}

func (a *application) addScene(key string, scene Scene) {
	if _, found := a.scenes[key]; found {
		panic(fmt.Errorf("duplicated scene key %v", key))
	}

	a.scenes[key] = scene
}

func (a *application) changeScene(key string) {
	if _, found := a.scenes[key]; !found {
		panic(fmt.Errorf("cannot find %v scene", key))
	}

	a.curScene = a.scenes[key]
	a.curScene.Reset()
	a.curScene.Init()
}
