package barko

type ScaleBy interface {
	Action
}

type scaleBy struct {
	*BaseAction
	deltaX, deltaY float32
	dx, dy         float32
	timer          float32
}

func NewScaleBy(scaleX, scaleY float32, interval float32) ScaleBy {
	return &scaleBy{NewBaseActionWithInterval(interval), scaleX, scaleY, 1, 1, 0}
}

func (s *scaleBy) Perform(node Node, dt float32) {
	if !s.IsDisabled() && !s.IsPaused() {
		s.timer += dt
		if s.timer >= s.interval {
			s.disableState = true
			nodeX, nodeY := node.GetScale()
			node.SetScale(nodeX/s.dx*s.deltaX, nodeY/s.dy*s.deltaY)
		} else {
			nodeX, nodeY := node.GetScale()
			nodeX, nodeY = nodeX/s.dx, nodeY/s.dy
			s.dx, s.dy = s.timer/s.interval*(s.deltaX-1)+1, s.timer/s.interval*(s.deltaY-1)+1
			node.SetScale(nodeX*s.dx, nodeY*s.dy)
		}
	}
}

func (s *scaleBy) Reset() {
	s.BaseAction.Reset()
	s.dx, s.dy = 0, 0
	s.timer = 0
}
