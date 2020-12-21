package shapes

type Rect struct {
	height float64
	width  float64
}

func (r rect) Area() float64 {
	return r.height * r.width
}
