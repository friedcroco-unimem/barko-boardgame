package barko

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Mouse ebiten.MouseButton

const (
	MouseLeft   Mouse = Mouse(ebiten.MouseButtonLeft)
	MouseRight        = ebiten.MouseButtonRight
	MouseMiddle       = ebiten.MouseButtonMiddle
)

type MouseListener interface {
	Event
	SetOnMouseDown(f func(x, y float32))
	SetOnMouseUp(f func(x, y float32))
	SetOnMousePressed(f func(x, y float32))
	SetOnMouseHover(f func(x, y float32))
}

type mouseListener struct {
	*BaseEvent
	mouse          Mouse
	onMouseDown    func(x, y float32)
	onMouseUp      func(x, y float32)
	onMousePressed func(x, y float32)
	onMouseHover   func(x, y float32)
	isPressed      bool
}

func NewMouseListener(mouse Mouse) MouseListener {
	return &mouseListener{NewBaseEvent(), mouse,
		func(x, y float32) {}, func(x, y float32) {}, func(x, y float32) {}, func(x, y float32) {}, false}
}

func (m *mouseListener) SetOnMouseDown(f func(x, y float32)) {
	m.onMouseDown = f
}

func (m *mouseListener) SetOnMouseUp(f func(x, y float32)) {
	m.onMouseUp = f
}

func (m *mouseListener) SetOnMousePressed(f func(x, y float32)) {
	m.onMousePressed = f
}

func (m *mouseListener) SetOnMouseHover(f func(x, y float32)) {
	m.onMouseHover = f
}

func (m *mouseListener) Update(dt float32) {
	if !m.IsPaused() && !m.IsRemoved() {
		x, y := ebiten.CursorPosition()
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton(m.mouse)) {
			m.onMouseDown(float32(x), float32(y))
			m.isPressed = true
			return
		}

		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton(m.mouse)) {
			m.onMouseUp(float32(x), float32(y))
			m.isPressed = false
			return
		}

		if m.isPressed {
			m.onMousePressed(float32(x), float32(y))
			return
		}

		m.onMouseHover(float32(x), float32(y))
	}
}

type mouseOverNodeListener struct {
	*mouseListener
	node Node
}

func NewMouseOverNodeListener(mouse Mouse, node Node) MouseListener {
	return &mouseOverNodeListener{NewMouseListener(mouse).(*mouseListener), node}
}

func (m *mouseOverNodeListener) Update(dt float32) {
	if !m.IsPaused() && !m.IsRemoved() {
		x, y := ebiten.CursorPosition()
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton(m.mouse)) && m.node.IsPointWithin(float32(x), float32(y)) {
			m.onMouseDown(float32(x), float32(y))
			m.isPressed = true
			return
		}

		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton(m.mouse)) {
			if m.isPressed && m.node.IsPointWithin(float32(x), float32(y)) {
				m.onMouseUp(float32(x), float32(y))
			}
			m.isPressed = false
			return
		}

		if m.isPressed {
			m.onMousePressed(float32(x), float32(y))
			return
		}

		if m.node.IsPointWithin(float32(x), float32(y)) && !m.isPressed {
			m.onMouseHover(float32(x), float32(y))
		}
	}
}
