package barko

import (
	"github.com/google/btree"
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	draw(screen *ebiten.Image)
	Update(dt float32)
	AddChild(node Node)
	RemoveChild(node Node)
	AddEvent(event Event)
	RemoveEvent(event Event)
	PauseAllEvents()
	ResumeAllEvents()
	PauseAllActions()
	ResumeAllActions()
	PauseAllAnimations()
	ResumeAllAnimations()
}

var _ Scene = &BaseScene{}

type BaseScene struct {
	children *btree.BTreeG[Node]
	events   []Event
	localID  int
}

func NewBaseScene() *BaseScene {
	return &BaseScene{
		btree.NewG(degreeChildTree, compareNode),
		make([]Event, 0),
		0,
	}
}

func (s *BaseScene) getNextLocalID() int {
	s.localID += 1
	return s.localID
}

func (s *BaseScene) draw(screen *ebiten.Image) {
	s.children.Ascend(func(node Node) bool {
		node.draw(screen)
		return true
	})
}

func (s *BaseScene) Update(dt float32) {
	removedChildren := make([]Node, 0)
	s.children.Ascend(func(node Node) bool {
		if node.IsRemoved() {
			removedChildren = append(removedChildren, node)
		}
		return true
	})
	for _, child := range removedChildren {
		s.children.Delete(child)
	}

	events := make([]Event, 0)
	for _, event := range s.events {
		if !event.IsRemoved() {
			events = append(events, event)
		}
	}
	s.events = events

	s.children.Ascend(func(node Node) bool {
		node.Update(dt)
		return true
	})

	for _, event := range s.events {
		event.Update(dt)
	}
}

func (s *BaseScene) AddChild(node Node) {
	node.SetLocalID(s.getNextLocalID())
	s.children.ReplaceOrInsert(node)
}

func (s *BaseScene) RemoveChild(node Node) {
	node.RemoveFromScene()
}

func (s *BaseScene) AddEvent(event Event) {
	s.events = append(s.events, event)
}

func (s *BaseScene) RemoveEvent(event Event) {
	event.RemoveFromScene()
}

func (s *BaseScene) PauseAllEvents() {
	for _, event := range s.events {
		event.Pause()
	}
}

func (s *BaseScene) ResumeAllEvents() {
	for _, event := range s.events {
		event.Resume()
	}
}

func (s *BaseScene) PauseAllActions() {
	s.children.Ascend(func(node Node) bool {
		node.PauseAllActions()
		return true
	})
}

func (s *BaseScene) ResumeAllActions() {
	s.children.Ascend(func(node Node) bool {
		node.ResumeAllActions()
		return true
	})
}

func (s *BaseScene) PauseAllAnimations() {
	s.children.Ascend(func(node Node) bool {
		node.PauseAllAnimations()
		return true
	})
}

func (s *BaseScene) ResumeAllAnimations() {
	s.children.Ascend(func(node Node) bool {
		node.ResumeAllAnimations()
		return true
	})
}
