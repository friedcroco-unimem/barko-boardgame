package barko

type Rect struct {
	X, Y          float32
	Width, Height float32
}

func NewRect(x, y float32, width, height float32) *Rect {
	return &Rect{x, y, width, height}
}
