package barko

import "github.com/hajimehoshi/ebiten/v2"

type RepeatForever interface {
	Action
}

type repeatForever struct {
	*BaseAction
	action Action
}

func NewRepeatForever(action Action) RepeatForever {
	return &repeatForever{NewBaseAction(), action}
}

func (r *repeatForever) Perform(node Node, dt float32) {
	if !r.IsDisabled() && !r.IsPaused() {
		if r.action.IsDisabled() {
			r.action.Reset()
		}
		r.action.Perform(node, dt)
	}
}

func (r *repeatForever) draw(screen *ebiten.Image) {
	r.action.draw(screen)
}

type Wait interface {
	Action
}

type wait struct {
	*BaseAction
	timer float32
}

func NewWait(interval float32) Wait {
	return &wait{NewBaseActionWithInterval(interval), 0}
}

func (w *wait) Perform(node Node, dt float32) {
	if !w.IsDisabled() && !w.IsPaused() {
		w.timer += dt
		if w.timer >= w.interval {
			w.disableState = true
		}
	}
}

func (w *wait) Reset() {
	w.BaseAction.Reset()
	w.timer = 0
}

type Sequence interface {
	Action
	AddAction(action Action)
}

type sequence struct {
	*BaseAction
	index   int
	actions []Action
}

func NewSequence() Sequence {
	return &sequence{NewBaseAction(), 0, make([]Action, 0)}
}

func NewSequenceWithActions(actions []Action) Sequence {
	return &sequence{NewBaseAction(), 0, actions}
}

func (s *sequence) AddAction(action Action) {
	s.actions = append(s.actions, action)
	if action.GetInterval() > 0 {
		s.interval = 1
	}
}

func (s *sequence) Perform(node Node, dt float32) {
	if !s.IsDisabled() && !s.IsPaused() {
		if s.index >= len(s.actions) {
			s.disableState = true
		} else {
			for s.index < len(s.actions) {
				if s.actions[s.index].GetInterval() == 0 {
					s.actions[s.index].Perform(node, dt)
					if s.actions[s.index].IsDisabled() {
						s.index += 1
					} else {
						break
					}
				} else {
					s.actions[s.index].Perform(node, dt)
					break
				}
			}

			if s.index >= len(s.actions) {
				s.disableState = true
				return
			}

			if s.actions[s.index].IsDisabled() {
				s.index += 1
				if s.index >= len(s.actions) {
					s.disableState = true
				}
			}
		}
	}
}

func (s *sequence) draw(screen *ebiten.Image) {
	if !s.IsDisabled() {
		if s.index >= len(s.actions) {
			s.disableState = true
		} else {
			s.actions[s.index].draw(screen)
		}
	}
}

func (s *sequence) Reset() {
	s.BaseAction.Reset()
	s.index = 0
	for _, action := range s.actions {
		action.Reset()
	}
}

type Parallel interface {
	Action
	AddAction(action Action)
}

type parallel struct {
	*BaseAction
	actions []Action
}

func NewParallel() Parallel {
	return &parallel{NewBaseAction(), make([]Action, 0)}
}

func NewParallelWithActions(actions []Action) Parallel {
	return &parallel{NewBaseAction(), actions}
}

func (p *parallel) AddAction(action Action) {
	p.actions = append(p.actions, action)
	if action.GetInterval() > 0 {
		p.interval = 1
	}
}

func (p *parallel) Perform(node Node, dt float32) {
	if !p.IsDisabled() && !p.IsPaused() {
		p.disableState = true
		for _, action := range p.actions {
			action.Perform(node, dt)
			if !action.IsDisabled() {
				p.disableState = false
			}
		}
	}
}

func (p *parallel) draw(screen *ebiten.Image) {
	if !p.IsDisabled() {
		for _, action := range p.actions {
			action.draw(screen)
		}
	}
}

func (p *parallel) Reset() {
	p.BaseAction.Reset()
	for _, action := range p.actions {
		action.Reset()
	}
}
