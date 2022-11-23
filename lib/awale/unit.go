package awale

type Unit interface {
	GetValue() int
}

type unit struct {
	val int
}

func (u *unit) GetValue() int {
	return u.val
}

var normalUnitSample Unit = &unit{1}
var bossUnitSample Unit = &unit{5}
