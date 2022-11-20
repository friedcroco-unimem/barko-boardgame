package barko

type Disappear interface {
	Action
}

type disappear struct {
	*BaseAction
}

func NewDisappear() Disappear {
	return &disappear{NewBaseAction()}
}

func (d *disappear) Perform(node Node, dt float32) {
	if !d.IsDisabled() && !d.IsPaused() {
		d.disableState = true
		node.RemoveFromScene()
	}
}

type FuncCall interface {
	Action
	SetFunc(f func())
}

type funcCall struct {
	*BaseAction
	function func()
}

func NewFuncCall() FuncCall {
	return &funcCall{NewBaseAction(), func() {}}
}

func NewFuncCallWithFunc(f func()) FuncCall {
	return &funcCall{NewBaseAction(), f}
}

func (d *funcCall) SetFunc(f func()) {
	d.function = f
}

func (d *funcCall) Perform(node Node, dt float32) {
	if !d.IsDisabled() && !d.IsPaused() {
		d.disableState = true
		d.function()
	}
}
