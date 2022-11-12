package barko

import "github.com/hajimehoshi/ebiten/v2"

type Action interface {
	draw(screen *ebiten.Image)
	Perform(node Node, dt float32)
	Reset()
	Disable()
	Enable()
	IsDisabled() bool
	IsRemovedOnDisabled() bool
	Pause()
	IsPaused() bool
	Resume()
	GetInterval() float32
}

var _ Action = &BaseAction{}

type BaseAction struct {
	disableState     bool
	pauseState       bool
	interval         float32
	removeOnDisabled bool
}

func NewBaseAction() *BaseAction {
	return &BaseAction{false, false, 0, true}
}

func NewBaseActionWithInterval(interval float32) *BaseAction {
	return &BaseAction{false, false, interval, true}
}

func (a *BaseAction) draw(screen *ebiten.Image) {}

func (a *BaseAction) Perform(node Node, dt float32) {}

func (a *BaseAction) Reset() {
	a.disableState = false
}

func (a *BaseAction) Disable() {
	a.disableState = true
}

func (a *BaseAction) Enable() {
	a.disableState = false
}

func (a *BaseAction) IsDisabled() bool {
	return a.disableState
}

func (a *BaseAction) IsRemovedOnDisabled() bool {
	return a.removeOnDisabled
}

func (a *BaseAction) Pause() {
	a.pauseState = true
}

func (a *BaseAction) IsPaused() bool {
	return a.pauseState
}

func (a *BaseAction) Resume() {
	a.pauseState = false
}

func (a *BaseAction) GetInterval() float32 {
	return a.interval
}
