package barko

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Node interface {
	draw(screen *ebiten.Image)
	Update(dt float32)
	GetLocalID() int
	SetLocalID(id int)
	GetPosition() (float32, float32)
	SetPosition(x float32, y float32)
	GetOffsetZ() float32
	SetOffsetZ(z float32)
	GetSize() (float32, float32)
	SetSize(width float32, height float32)
	GetAnchor() (float32, float32)
	SetAnchor(x float32, y float32)
	GetScale() (float32, float32)
	SetScale(x float32, y float32)
	GetRotation() float32
	SetRotation(angle float32)
	GetOpacity() float32
	SetOpacity(opacity float32)
	AddAction(action Action)
	RemoveFromScene()
	IsRemoved() bool
	PauseAllActions()
	ResumeAllActions()
	PauseAllAnimations()
	ResumeAllAnimations()
	isAnimated() bool
	setAnimated(state bool)
}

func compareNode(a Node, b Node) bool {
	if a.GetOffsetZ() != b.GetOffsetZ() {
		return a.GetOffsetZ() < b.GetOffsetZ()
	}

	return a.GetLocalID() < b.GetLocalID()
}

var _ Node = &BaseNode{}

type BaseNode struct {
	x, y             float32
	offsetZ          float32
	localID          int
	width, height    float32
	anchorX, anchorY float32
	scaleX, scaleY   float32
	angle            float32
	actions          []Action
	opacity          float32
	animatedState    bool
	removeState      bool
}

func NewBaseNode() *BaseNode {
	return &BaseNode{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, make([]Action, 0), 1, false, false}
}

func (n *BaseNode) draw(screen *ebiten.Image) {
	for _, action := range n.actions {
		action.draw(screen)
	}
}

func (n *BaseNode) isAnimated() bool {
	return n.animatedState
}

func (n *BaseNode) setAnimated(state bool) {
	n.animatedState = state
}

func (n *BaseNode) Update(dt float32) {
	n.animatedState = false

	actions := make([]Action, 0)
	for _, action := range n.actions {
		if !(action.IsDisabled() && action.IsRemovedOnDisabled()) {
			actions = append(actions, action)
		}
	}

	n.actions = actions
	for _, action := range n.actions {
		action.Perform(n, dt)
	}
}

func (n *BaseNode) GetLocalID() int {
	return n.localID
}

func (n *BaseNode) SetLocalID(id int) {
	n.localID = id
}

func (n *BaseNode) GetPosition() (float32, float32) {
	return n.x, n.y
}

func (n *BaseNode) SetPosition(x float32, y float32) {
	n.x = x
	n.y = y
}

func (n *BaseNode) GetOffsetZ() float32 {
	return n.offsetZ
}

func (n *BaseNode) SetOffsetZ(z float32) {
	n.offsetZ = z
}

func (n *BaseNode) GetSize() (float32, float32) {
	return n.width, n.height
}

func (n *BaseNode) SetSize(width float32, height float32) {
	n.width = width
	n.height = height
}

func (n *BaseNode) GetAnchor() (float32, float32) {
	return n.anchorX, n.anchorY
}

func (n *BaseNode) SetAnchor(x float32, y float32) {
	n.anchorX = x
	n.anchorY = y
}

func (n *BaseNode) GetScale() (float32, float32) {
	return n.scaleX, n.scaleY
}

func (n *BaseNode) SetScale(x float32, y float32) {
	n.scaleX = x
	n.scaleY = y
}

func (n *BaseNode) GetRotation() float32 {
	return n.angle
}

func (n *BaseNode) SetRotation(angle float32) {
	n.angle = angle
	if n.angle >= maxAngle {
		n.angle -= float32(int(n.angle/maxAngle)) * maxAngle
	}

	if n.angle < maxAngle {
		n.angle -= float32(int(n.angle/maxAngle+1)) * maxAngle
	}
}

func (n *BaseNode) GetOpacity() float32 {
	return n.opacity
}

func (n *BaseNode) SetOpacity(opacity float32) {
	n.opacity = opacity
}

func (n *BaseNode) AddAction(action Action) {
	n.actions = append(n.actions, action)
}

func (n *BaseNode) RemoveFromScene() {
	n.removeState = true
}

func (n *BaseNode) IsRemoved() bool {
	return n.removeState
}

func (n *BaseNode) PauseAllActions() {
	for _, action := range n.actions {
		action.Pause()
	}
}

func (n *BaseNode) ResumeAllActions() {
	for _, action := range n.actions {
		action.Resume()
	}
}

func (n *BaseNode) PauseAllAnimations() {
	for _, action := range n.actions {
		if _, ok := action.(Animation); ok {
			action.Pause()
		}
	}
}

func (n *BaseNode) ResumeAllAnimations() {
	for _, action := range n.actions {
		if _, ok := action.(Animation); ok {
			action.Resume()
		}
	}
}
