package awale

import (
	"barkdo/lib/barko"
)

func NewDigitSprite(digit int) *barko.BaseSprite {
	res := barko.NewSpriteWithRect("digit", barko.NewRect(float32(digit*16), 0, 16, 24))
	res.SetAnchor(0.5, 0.5)
	return res
}

type digitGroup struct {
	digits     []barko.Sprite
	value      int
	posX, posY float32
	s          barko.Scene
}

func NewDigitGroup(s barko.Scene, posX, posY float32) *digitGroup {
	res := &digitGroup{make([]barko.Sprite, 0), 0, posX, posY, s}
	res.ChangeValue(0)
	return res
}

func (d *digitGroup) GetValue() int {
	return d.value
}

func (d *digitGroup) ChangeValue(val int) {
	d.value = val

	for _, digit := range d.digits {
		digit.RemoveFromScene()
	}
	d.digits = make([]barko.Sprite, 0)

	t := val
	if t > 0 {
		for ; t > 0; t /= 10 {
			digit := t % 10
			sprite := NewDigitSprite(digit)
			d.s.AddChild(sprite)
			d.digits = append(d.digits, sprite)
		}
	} else {
		sprite := NewDigitSprite(0)
		d.s.AddChild(sprite)
		d.digits = append(d.digits, sprite)
	}

	posX := d.posX - float32(len(d.digits)-1)/2*20
	for i := len(d.digits) - 1; i >= 0; i-- {
		d.digits[i].SetPosition(posX, d.posY)
		posX += 16
	}
}
