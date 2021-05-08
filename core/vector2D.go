package core

type Vector2D interface {
	X() float64
	Y() float64
	SetX(value float64)
	SetY(value float64)
}

type vector2D struct {
	x, y float64
}

func NewVector2D(x, y float64) Vector2D {
	return &vector2D{
		x: x,
		y: y,
	}
}

func (v *vector2D) X() float64 {
	return v.x
}

func (v *vector2D) Y() float64 {
	return v.y
}

func (v *vector2D) SetX(value float64) {
	v.x = value
}

func (v *vector2D) SetY(value float64) {
	v.y = value
}
