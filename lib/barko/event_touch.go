package barko

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type TouchListener interface {
	Event
	SetOnTouchDown(f func(x, y float32))
	SetOnTouchUp(f func(x, y float32))
	SetOnTouchPressed(f func(x, y float32))
}

type touchListener struct {
	*BaseEvent
	onTouchDown    func(x, y float32)
	onTouchUp      func(x, y float32)
	onTouchPressed func(x, y float32)
	touches        map[ebiten.TouchID]bool
}

func NewTouchListener() *touchListener {
	return &touchListener{NewBaseEvent(),
		func(x, y float32) {}, func(x, y float32) {}, func(x, y float32) {}, make(map[ebiten.TouchID]bool)}
}

func (t *touchListener) SetOnTouchDown(f func(x, y float32)) {
	t.onTouchDown = f
}

func (t *touchListener) SetOnTouchUp(f func(x, y float32)) {
	t.onTouchUp = f
}

func (t *touchListener) SetOnTouchPressed(f func(x, y float32)) {
	t.onTouchPressed = f
}

func (t *touchListener) Update(dt float32) {
	// release touch
	for k := range t.touches {
		if inpututil.IsTouchJustReleased(k) {
			x, y := ebiten.TouchPosition(k)
			if !t.IsPaused() && !t.IsRemoved() {
				t.onTouchUp(float32(x), float32(y))
			}
			delete(t.touches, k)
			continue
		}

		// on touch pressed
		x, y := ebiten.TouchPosition(k)
		if !t.IsPaused() && !t.IsRemoved() {
			t.onTouchPressed(float32(x), float32(y))
		}
	}

	// new touch
	newTouches := inpututil.AppendJustPressedTouchIDs(make([]ebiten.TouchID, 0))
	log.Printf("%v", t.touches)
	for _, n := range newTouches {
		x, y := ebiten.TouchPosition(n)
		if !t.IsPaused() && !t.IsRemoved() {
			t.onTouchDown(float32(x), float32(y))
		}
		t.touches[n] = true
	}
}
