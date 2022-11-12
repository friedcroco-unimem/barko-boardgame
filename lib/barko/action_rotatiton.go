package barko

type RotateBy interface {
	Action
}

type rotateBy struct {
	*BaseAction
	deltaR float32
	dr     float32
	timer  float32
}

func NewRotateBy(deltaR float32, interval float32) RotateBy {
	return &rotateBy{NewBaseActionWithInterval(interval), deltaR, 0, 0}
}

func (r *rotateBy) Perform(node Node, dt float32) {
	if !r.IsDisabled() && !r.IsPaused() {
		r.timer += dt
		if r.timer >= r.interval {
			r.disableState = true
			nodeR := node.GetRotation()
			node.SetRotation(nodeR - r.dr + r.deltaR)
		} else {
			nodeR := node.GetRotation()
			nodeR = nodeR - r.dr
			r.dr = r.timer / r.interval * r.deltaR
			node.SetRotation(nodeR + r.dr)
		}
	}
}

func (r *rotateBy) Reset() {
	r.BaseAction.Reset()
	r.dr = 0
	r.timer = 0
}
