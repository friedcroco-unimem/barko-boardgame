package barko

type FadeIn interface {
	Action
}

type fadeIn struct {
	*BaseAction
	timer float32
}

func NewFadeIn(interval float32) FadeIn {
	return &fadeIn{NewBaseActionWithInterval(interval), 0}
}

func (f *fadeIn) Perform(node Node, dt float32) {
	if !f.IsDisabled() && !f.IsPaused() {
		f.timer += dt
		if f.timer >= f.interval {
			f.disableState = true
			node.SetOpacity(1)
		} else {
			node.SetOpacity(f.timer / f.interval)
		}
	}
}

func (f *fadeIn) Reset() {
	f.BaseAction.Reset()
	f.timer = 0
}

type FadeOut interface {
	Action
}

type fadeOut struct {
	*BaseAction
	timer float32
}

func NewFadeOut(interval float32) FadeOut {
	return &fadeOut{NewBaseActionWithInterval(interval), 0}
}

func (f *fadeOut) Perform(node Node, dt float32) {
	if !f.IsDisabled() && !f.IsPaused() {
		f.timer += dt
		if f.timer >= f.interval {
			f.disableState = true
			node.SetOpacity(0)
		} else {
			node.SetOpacity(1 - f.timer/f.interval)
		}
	}
}

func (f *fadeOut) Reset() {
	f.BaseAction.Reset()
	f.timer = 0
}

type FadeTo interface {
	Action
}

type fadeTo struct {
	*BaseAction
	opacity  float32
	finalOpa float32
	deltaOpa float32
	betOpa   float32
	timer    float32
}

func NewFadeTo(opacity float32, interval float32) FadeTo {
	return &fadeTo{NewBaseActionWithInterval(interval), 1, 1, 0, 0, 0}
}

func (f *fadeTo) Perform(node Node, dt float32) {
	if !f.IsDisabled() && !f.IsPaused() {
		if f.timer == 0 {
			f.opacity = node.GetOpacity()
			f.deltaOpa = f.finalOpa - f.opacity
		}

		f.timer += dt
		if f.timer >= f.interval {
			f.disableState = true
			node.SetOpacity(f.finalOpa)
		} else {
			f.opacity = f.opacity - f.betOpa
			f.betOpa = f.deltaOpa * f.timer / f.interval
			f.opacity = f.opacity + f.betOpa
			node.SetOpacity(f.opacity)
		}
	}
}

func (f *fadeTo) Reset() {
	f.BaseAction.Reset()
	f.timer = 0
	f.betOpa = 0
	f.deltaOpa = 0
	f.opacity = 1
}
