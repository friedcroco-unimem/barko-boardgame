package barko

type Event interface {
	Update(dt float32)
	RemoveFromScene()
	IsRemoved() bool
	Pause()
	IsPaused() bool
	Resume()
}

type BaseEvent struct {
	removeState bool
	pauseState  bool
}

func NewBaseEvent() *BaseEvent {
	return &BaseEvent{false, false}
}

func (e *BaseEvent) Update(dt float32) {}

func (e *BaseEvent) RemoveFromScene() {
	e.removeState = true
}

func (e *BaseEvent) IsRemoved() bool {
	return e.removeState
}

func (e *BaseEvent) Pause() {
	e.pauseState = true
}

func (e *BaseEvent) IsPaused() bool {
	return e.pauseState
}

func (e *BaseEvent) Resume() {
	e.pauseState = false
}
