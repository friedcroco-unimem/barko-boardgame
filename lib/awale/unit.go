package awale

type Unit interface {
	GetValue() int
}

type normalUnit struct{}

func NewNormalUnit() Unit {
	return &normalUnit{}
}

func (u *normalUnit) GetValue() int { return 1 }

type bossUnit struct{}

func NewBossUnit() Unit {
	return &bossUnit{}
}

func (u *bossUnit) GetValue() int { return 5 }

var normalUnitSample Unit = NewNormalUnit()
var bossUnitSample Unit = NewBossUnit()
