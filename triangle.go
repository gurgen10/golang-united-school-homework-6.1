package golang_united_school_homework

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (t Triangle) CalcArea() float64 {
	return (t.Side * t.Side) / 2
}

func (t Triangle) CalcPerimeter() float64 {
	return t.Side * 3
}
