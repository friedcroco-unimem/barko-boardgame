package barko

type MoveBy interface {
	Action
}

type moveBy struct {
	*BaseAction
	deltaX, deltaY float32
	dx, dy         float32
	timer          float32
}

func NewMoveBy(deltaX, deltaY float32, interval float32) MoveBy {
	return &moveBy{NewBaseActionWithInterval(interval), deltaX, deltaY, 0, 0, 0}
}

func (m *moveBy) Perform(node Node, dt float32) {
	if !m.IsDisabled() && !m.IsPaused() {
		m.timer += dt
		if m.timer >= m.interval {
			m.disableState = true
			nodeX, nodeY := node.GetPosition()
			node.SetPosition(nodeX-m.dx+m.deltaX, nodeY-m.dy+m.deltaY)
		} else {
			nodeX, nodeY := node.GetPosition()
			nodeX, nodeY = nodeX-m.dx, nodeY-m.dy
			m.dx, m.dy = m.timer/m.interval*m.deltaX, m.timer/m.interval*m.deltaY
			node.SetPosition(nodeX+m.dx, nodeY+m.dy)
		}
	}
}

func (m *moveBy) Reset() {
	m.BaseAction.Reset()
	m.dx, m.dy = 0, 0
	m.timer = 0
}

type MoveAnchorBy interface {
	Action
}

type moveAnchorBy struct {
	*BaseAction
	deltaX, deltaY float32
	dx, dy         float32
	timer          float32
}

func NewAnchorMoveBy(deltaX, deltaY float32, interval float32) MoveAnchorBy {
	return &moveAnchorBy{NewBaseActionWithInterval(interval), deltaX, deltaY, 0, 0, 0}
}

func (m *moveAnchorBy) Perform(node Node, dt float32) {
	if !m.IsDisabled() && !m.IsPaused() {
		m.timer += dt
		if m.timer >= m.interval {
			m.disableState = true
			nodeX, nodeY := node.GetAnchor()
			node.SetAnchor(nodeX-m.dx+m.deltaX, nodeY-m.dy+m.deltaY)
		} else {
			nodeX, nodeY := node.GetAnchor()
			nodeX, nodeY = nodeX-m.dx, nodeY-m.dy
			m.dx, m.dy = m.timer/m.interval*m.deltaX, m.timer/m.interval*m.deltaY
			node.SetAnchor(nodeX+m.dx, nodeY+m.dy)
		}
	}
}

func (m *moveAnchorBy) Reset() {
	m.BaseAction.Reset()
	m.dx, m.dy = 0, 0
	m.timer = 0
}

type JumpBy interface {
	Action
}

type jumpBy struct {
	*BaseAction
	deltaX, deltaY float32
	height         float32
	dx, dy         float32
	h              float32
	timer          float32
}

func NewJumpBy(deltaX, deltaY float32, height float32, interval float32) JumpBy {
	return &jumpBy{NewBaseActionWithInterval(interval), deltaX, deltaY, height, 0, 0, 0, 0}
}

func (j *jumpBy) Perform(node Node, dt float32) {
	if !j.IsDisabled() && !j.IsPaused() {
		j.timer += dt
		if j.timer >= j.interval {
			j.disableState = true
			nodeX, nodeY := node.GetPosition()
			node.SetPosition(nodeX-j.dx+j.deltaX, nodeY-j.dy+j.deltaY+j.h)
		} else {
			nodeX, nodeY := node.GetPosition()
			nodeX, nodeY = nodeX-j.dx, nodeY-j.dy+j.h
			j.dx, j.dy = j.timer/j.interval*j.deltaX, j.timer/j.interval*j.deltaY
			j.h = 4*j.height/j.interval*j.timer - 4*j.height/(j.interval*j.interval)*j.timer*j.timer
			node.SetPosition(nodeX+j.dx, nodeY+j.dy-j.h)
		}
	}
}

func (j *jumpBy) Reset() {
	j.BaseAction.Reset()
	j.dx, j.dy = 0, 0
	j.h = 0
	j.timer = 0
}
