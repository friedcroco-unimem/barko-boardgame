package barko

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type spritePool struct {
	sprites map[string]*ebiten.Image
}

func newSpritePool() *spritePool {
	return &spritePool{make(map[string]*ebiten.Image)}
}

func (p *spritePool) preloadImage(key string, filename string) {
	image, _, err := ebitenutil.NewImageFromFile(filename)
	if err != nil {
		panic(err)
	}

	p.sprites[key] = image
}

func (p *spritePool) getImage(key string) *ebiten.Image {
	if found, ok := p.sprites[key]; ok {
		return found
	}

	panic(fmt.Errorf("no %v image preloaded", key))
}
