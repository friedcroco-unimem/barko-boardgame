package barko

type Event interface {
	Update(dt float32)
	RemoveFromScene()
	IsRemoved() bool
	Pause()
	Resume()
}
