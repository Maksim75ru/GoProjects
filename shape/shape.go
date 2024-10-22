package shape

type Figure interface {
	GetSquare() (int, error)
}

type Square struct {
	Width  int
	Length int
}

func NewSquare(width int) Square {
	return Square{
		Width:  width,
		Length: width,
	}
}

func (s Square) GetSquare() int {
	return s.Length * s.Width
}

type Rectangle struct {
	Width  int
	Length int
}

func NewRectangle(width, length int) Rectangle {
	return Rectangle{
		Width:  width,
		Length: length,
	}
}

func (r Rectangle) GetSquare() int {
	return r.Length * r.Width
}
