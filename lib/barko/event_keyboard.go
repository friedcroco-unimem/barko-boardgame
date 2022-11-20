package barko

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Key ebiten.Key

const (
	KeyA            = ebiten.KeyA
	KeyB            = ebiten.KeyB
	KeyC            = ebiten.KeyC
	KeyD            = ebiten.KeyD
	KeyE            = ebiten.KeyE
	KeyF            = ebiten.KeyF
	KeyG            = ebiten.KeyG
	KeyH            = ebiten.KeyH
	KeyI            = ebiten.KeyI
	KeyJ            = ebiten.KeyJ
	KeyK            = ebiten.KeyK
	KeyL            = ebiten.KeyL
	KeyM            = ebiten.KeyM
	KeyN            = ebiten.KeyN
	KeyO            = ebiten.KeyO
	KeyP            = ebiten.KeyP
	KeyQ            = ebiten.KeyQ
	KeyR            = ebiten.KeyR
	KeyS            = ebiten.KeyS
	KeyT            = ebiten.KeyT
	KeyU            = ebiten.KeyU
	KeyV            = ebiten.KeyV
	KeyW            = ebiten.KeyW
	KeyX            = ebiten.KeyX
	KeyY            = ebiten.KeyY
	KeyZ            = ebiten.KeyZ
	KeyArrowDown    = ebiten.KeyArrowDown
	KeyArrowLeft    = ebiten.KeyArrowLeft
	KeyArrowRight   = ebiten.KeyArrowRight
	KeyArrowUp      = ebiten.KeyArrowUp
	KeyBackquote    = ebiten.KeyBackquote
	KeyBackslash    = ebiten.KeyBackslash
	KeyBackspace    = ebiten.KeyBackspace
	KeyBracketLeft  = ebiten.KeyBracketLeft
	KeyBracketRight = ebiten.KeyBracketRight
	KeyComma        = ebiten.KeyComma
	KeyDigit0       = ebiten.KeyDigit0
	KeyDigit1       = ebiten.KeyDigit1
	KeyDigit2       = ebiten.KeyDigit2
	KeyDigit3       = ebiten.KeyDigit3
	KeyDigit4       = ebiten.KeyDigit4
	KeyDigit5       = ebiten.KeyDigit5
	KeyDigit6       = ebiten.KeyDigit6
	KeyDigit7       = ebiten.KeyDigit7
	KeyDigit8       = ebiten.KeyDigit8
	KeyDigit9       = ebiten.KeyDigit9
	KeyEnter        = ebiten.KeyEnter
	KeyEqual        = ebiten.KeyEqual
	KeyEscape       = ebiten.KeyEscape
	KeyMinus        = ebiten.KeyMinus
	KeyPeriod       = ebiten.KeyPeriod
	KeyQuote        = ebiten.KeyQuote
	KeySemicolon    = ebiten.KeySemicolon
	KeySlash        = ebiten.KeySlash
	KeySpace        = ebiten.KeySpace
	KeyTab          = ebiten.KeyTab
	KeyAlt          = ebiten.KeyAlt
	KeyControl      = ebiten.KeyControl
	KeyShift        = ebiten.KeyShift
)

type KeyboardListener interface {
	Event
	SetOnKeyDown(f func())
	SetOnKeyUp(f func())
	SetOnKeyPressed(f func())
	GetKey() Key
}

type keyboardListener struct {
	*BaseEvent
	key          Key
	onKeyDown    func()
	onKeyUp      func()
	onKeyPressed func()
	isPressed    bool
}

func NewKeyboardListener(key Key) KeyboardListener {
	return &keyboardListener{NewBaseEvent(), key, func() {}, func() {}, func() {}, false}
}

func (k *keyboardListener) GetKey() Key {
	return k.key
}

func (k *keyboardListener) SetOnKeyDown(f func()) {
	k.onKeyDown = f
}

func (k *keyboardListener) SetOnKeyUp(f func()) {
	k.onKeyUp = f
}

func (k *keyboardListener) SetOnKeyPressed(f func()) {
	k.onKeyPressed = f
}

func (k *keyboardListener) Update(dt float32) {
	if !k.IsPaused() && !k.IsRemoved() {
		if inpututil.IsKeyJustPressed(ebiten.Key(k.key)) {
			k.onKeyDown()
			k.isPressed = true
			return
		}

		if inpututil.IsKeyJustReleased(ebiten.Key(k.key)) {
			k.onKeyUp()
			k.isPressed = false
			return
		}

		if k.isPressed {
			k.onKeyPressed()
		}
	}
}
