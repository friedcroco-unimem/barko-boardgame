package barko

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Key ebiten.Key

const (
	KeyA            = Key(ebiten.KeyA)
	KeyB            = Key(ebiten.KeyB)
	KeyC            = Key(ebiten.KeyC)
	KeyD            = Key(ebiten.KeyD)
	KeyE            = Key(ebiten.KeyE)
	KeyF            = Key(ebiten.KeyF)
	KeyG            = Key(ebiten.KeyG)
	KeyH            = Key(ebiten.KeyH)
	KeyI            = Key(ebiten.KeyI)
	KeyJ            = Key(ebiten.KeyJ)
	KeyK            = Key(ebiten.KeyK)
	KeyL            = Key(ebiten.KeyL)
	KeyM            = Key(ebiten.KeyM)
	KeyN            = Key(ebiten.KeyN)
	KeyO            = Key(ebiten.KeyO)
	KeyP            = Key(ebiten.KeyP)
	KeyQ            = Key(ebiten.KeyQ)
	KeyR            = Key(ebiten.KeyR)
	KeyS            = Key(ebiten.KeyS)
	KeyT            = Key(ebiten.KeyT)
	KeyU            = Key(ebiten.KeyU)
	KeyV            = Key(ebiten.KeyV)
	KeyW            = Key(ebiten.KeyW)
	KeyX            = Key(ebiten.KeyX)
	KeyY            = Key(ebiten.KeyY)
	KeyZ            = Key(ebiten.KeyZ)
	KeyArrowDown    = Key(ebiten.KeyArrowDown)
	KeyArrowLeft    = Key(ebiten.KeyArrowLeft)
	KeyArrowRight   = Key(ebiten.KeyArrowRight)
	KeyArrowUp      = Key(ebiten.KeyArrowUp)
	KeyBackquote    = Key(ebiten.KeyBackquote)
	KeyBackslash    = Key(ebiten.KeyBackslash)
	KeyBackspace    = Key(ebiten.KeyBackspace)
	KeyBracketLeft  = Key(ebiten.KeyBracketLeft)
	KeyBracketRight = Key(ebiten.KeyBracketRight)
	KeyComma        = Key(ebiten.KeyComma)
	KeyDigit0       = Key(ebiten.KeyDigit0)
	KeyDigit1       = Key(ebiten.KeyDigit1)
	KeyDigit2       = Key(ebiten.KeyDigit2)
	KeyDigit3       = Key(ebiten.KeyDigit3)
	KeyDigit4       = Key(ebiten.KeyDigit4)
	KeyDigit5       = Key(ebiten.KeyDigit5)
	KeyDigit6       = Key(ebiten.KeyDigit6)
	KeyDigit7       = Key(ebiten.KeyDigit7)
	KeyDigit8       = Key(ebiten.KeyDigit8)
	KeyDigit9       = Key(ebiten.KeyDigit9)
	KeyEnter        = Key(ebiten.KeyEnter)
	KeyEqual        = Key(ebiten.KeyEqual)
	KeyEscape       = Key(ebiten.KeyEscape)
	KeyMinus        = Key(ebiten.KeyMinus)
	KeyPeriod       = Key(ebiten.KeyPeriod)
	KeyQuote        = Key(ebiten.KeyQuote)
	KeySemicolon    = Key(ebiten.KeySemicolon)
	KeySlash        = Key(ebiten.KeySlash)
	KeySpace        = Key(ebiten.KeySpace)
	KeyTab          = Key(ebiten.KeyTab)
	KeyAlt          = Key(ebiten.KeyAlt)
	KeyControl      = Key(ebiten.KeyControl)
	KeyShift        = Key(ebiten.KeyShift)
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
