package barko

import (
	"errors"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite interface {
	Node
	setRect(rect *Rect)
	setImage(image *ebiten.Image)
}

var _ Sprite = &BaseSprite{}

type BaseSprite struct {
	*BaseNode
	image *ebiten.Image
	rect  *Rect
}

func NewSprite(imageKey string) *BaseSprite {
	image := GetGameManager().getImage(imageKey)
	width, height := image.Size()
	rect := NewRect(0, 0, float32(width), float32(height))
	res := &BaseSprite{NewBaseNode(), image, rect}
	res.width = float32(width)
	res.height = float32(height)
	return res
}

func NewSpriteWithRect(imageKey string, rect *Rect) *BaseSprite {
	if rect == nil {
		panic(errors.New("rect cannot be nil"))
	}

	image := GetGameManager().getImage(imageKey).
		SubImage(image.Rect(int(rect.X), int(rect.Y), int(rect.X+rect.Width), int(rect.Y+rect.Height))).(*ebiten.Image)
	width, height := image.Size()
	res := &BaseSprite{NewBaseNode(), image, rect}
	res.width = float32(width)
	res.height = float32(height)
	return res
}

func (i *BaseSprite) draw(screen *ebiten.Image) {
	if !i.isAnimated() {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(-float64(i.width)*float64(i.anchorX), -float64(i.height)*float64(i.anchorY))
		opts.GeoM.Rotate(2 * math.Pi * float64(i.angle) / float64(maxAngle))
		opts.GeoM.Scale(float64(i.scaleX), float64(i.scaleY))
		opts.GeoM.Translate(float64(i.width)*float64(i.anchorX), float64(i.height)*float64(i.anchorY))
		opts.GeoM.Translate(float64(i.x-i.anchorX*i.width*i.scaleX), float64(i.y-i.anchorY*i.height*i.scaleY))
		opts.ColorM.Reset()
		opts.ColorM.Scale(1, 1, 1, float64(i.opacity))
		screen.DrawImage(i.image, opts)
	}
	i.BaseNode.draw(screen)
}

func (i *BaseSprite) setRect(rect *Rect) {
	i.rect = rect
}

func (i *BaseSprite) setImage(image *ebiten.Image) {
	i.image = image
}

func (i *BaseSprite) Update(dt float32) {
	i.animatedState = false

	actions := make([]Action, 0)
	for _, action := range i.actions {
		if !(action.IsDisabled() && action.IsRemovedOnDisabled()) {
			actions = append(actions, action)
		}
	}

	i.actions = actions
	for _, action := range i.actions {
		action.Perform(i, dt)
	}
}
