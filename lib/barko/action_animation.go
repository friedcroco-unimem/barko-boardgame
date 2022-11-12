package barko

import (
	"errors"
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type AnimFrame interface {
	draw(node Node, screen *ebiten.Image)
	getInterval() float32
	getImage() *ebiten.Image
	getRect() *Rect
}

type animFrame struct {
	image    *ebiten.Image
	rect     *Rect
	interval float32
}

func NewAnimFrame(imageKey string, interval float32) AnimFrame {
	image := GetGameManager().getImage(imageKey)
	width, height := image.Size()
	rect := NewRect(0, 0, float32(width), float32(height))
	res := &animFrame{image, rect, interval}
	return res
}

func NewAnimFrameWithRect(imageKey string, rect *Rect, interval float32) AnimFrame {
	if rect == nil {
		panic(errors.New("rect cannot be nil"))
	}

	image := GetGameManager().getImage(imageKey).
		SubImage(image.Rect(int(rect.X), int(rect.Y), int(rect.X+rect.Width), int(rect.Y+rect.Height))).(*ebiten.Image)
	res := &animFrame{image, rect, interval}
	return res
}

func (a *animFrame) getInterval() float32 {
	return a.interval
}

func (a *animFrame) draw(node Node, screen *ebiten.Image) {
	nodeWidth, nodeHeight := node.GetSize()
	nodeAnchorX, nodeAnchorY := node.GetAnchor()
	nodeScaleX, nodeScaleY := node.GetScale()
	nodeX, nodeY := node.GetPosition()
	nodeAngle := node.GetRotation()
	nodeOpacity := node.GetOpacity()

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(-float64(nodeWidth)*float64(nodeAnchorX), -float64(nodeHeight)*float64(nodeAnchorY))
	opts.GeoM.Rotate(2 * math.Pi * float64(nodeAngle) / float64(maxAngle))
	opts.GeoM.Scale(float64(nodeScaleX), float64(nodeScaleY))
	opts.GeoM.Translate(float64(nodeWidth)*float64(nodeAnchorX), float64(nodeHeight)*float64(nodeAnchorY))
	opts.GeoM.Translate(float64(nodeX-nodeAnchorX*nodeWidth*nodeScaleX), float64(nodeY-nodeAnchorY*nodeHeight*nodeScaleY))
	opts.ColorM.Reset()
	opts.ColorM.Scale(1, 1, 1, float64(nodeOpacity))
	screen.DrawImage(a.image, opts)
}

func (a *animFrame) getImage() *ebiten.Image {
	return a.image
}

func (a *animFrame) getRect() *Rect {
	return a.rect
}

type Animation interface {
	Action
	AddFrame(frame AnimFrame)
}

type animation struct {
	*BaseAction
	anims []AnimFrame
	index int
	timer float32
	node  Node
}

func NewAnimation() Animation {
	return &animation{NewBaseAction(), make([]AnimFrame, 0), 0, 0, nil}
}

func NewAnimationWithFrames(frames []AnimFrame) Animation {
	return &animation{NewBaseAction(), frames, 0, 0, nil}
}

func (a *animation) Perform(node Node, dt float32) {
	a.node = node
	if !a.IsDisabled() && !a.IsPaused() {
		if a.index >= len(a.anims) {
			a.disableState = true
			if sprite, ok := node.(Sprite); ok && len(a.anims) > 0 {
				anim := a.anims[len(a.anims)-1]
				sprite.setImage(anim.getImage())
				rect := anim.getRect()
				sprite.setRect(rect)
				sprite.SetSize(rect.Width, rect.Height)
			}
		} else {
			node.setAnimated(true)
			a.timer += dt
			for a.timer >= a.anims[a.index].getInterval() {
				a.timer -= a.anims[a.index].getInterval()
				a.index += 1
				if a.index >= len(a.anims) {
					a.disableState = true
					node.setAnimated(false)
					if sprite, ok := node.(Sprite); ok && len(a.anims) > 0 {
						anim := a.anims[len(a.anims)-1]
						sprite.setImage(anim.getImage())
						rect := anim.getRect()
						sprite.setRect(rect)
						sprite.SetSize(rect.Width, rect.Height)
					}
					break
				}
			}
		}
	}
}

func (a *animation) draw(screen *ebiten.Image) {
	if !a.IsDisabled() {
		if a.index >= len(a.anims) {
			a.disableState = true
		} else {
			if a.node != nil {
				a.anims[a.index].draw(a.node, screen)
			}
		}
	}
}

func (a *animation) Reset() {
	a.BaseAction.Reset()
	a.index = 0
	a.timer = 0
	a.node = nil
}

func (a *animation) AddFrame(frame AnimFrame) {
	a.anims = append(a.anims, frame)
	a.interval = 0
}
