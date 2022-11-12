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
}

type funcCall struct {
	*BaseAction
	function func()
}

func NewFuncCall() FuncCall {
	return &funcCall{NewBaseAction(), func() {}}
}

func (d *funcCall) Perform(node Node, dt float32) {
	if !d.IsDisabled() && !d.IsPaused() {
		d.disableState = true
		d.function()
	}
}
