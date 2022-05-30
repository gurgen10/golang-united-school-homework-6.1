package main

import (
	"fmt"
	"math"
	"errors"
)

type Shape interface {
	// CalcPerimeter returns calculation result of perimeter
	CalcPerimeter() float64
	// CalcArea returns calculation result of area
	CalcArea() float64
}

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) CalcArea() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) CalcPerimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (r Rectangle) CalcArea() float64 {
	return r.Height * r.Weight
}

func (r Rectangle) CalcPerimeter() float64 {
	return 2 * (r.Height + r.Weight)
}

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

func main()  {
	box := NewBox(4)
	triangle := Triangle{Side: 5}
	circle := Circle{Radius: 12}
	rec := Rectangle{Height: 4, Weight: 8}
	err := box.AddShape(rec)
	fmt.Println(box)
	fmt.Println(err)
	fmt.Println(triangle)
	fmt.Println(circle)
	fmt.Println(rec)
}
// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity < len(b.shapes) {
		return errors.New("Out of range index")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if b.shapes[i] == nil {
		return nil, errors.New("Out of range index")
	}
	s := b.shapes[i]
	return s, nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if b.shapes[i] == nil {
		return nil, errors.New("Out of range index")
	}
	extracted := b.shapes[i]
	RemoveIndex(b.shapes, i)
	return extracted, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if b.shapes[i] == nil {
		return nil, errors.New("Out of range index")
	}

	b.shapes[i] = shape
	return shape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64 = 0.0
	for _, shape := range b.shapes {
		sum += shape.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64 = 0.0
	for _, shape := range b.shapes {
		sum += shape.CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var circleExists bool = false
	for i, shape := range b.shapes {
		shapeType := fmt.Sprintf("%T", shape)
		if shapeType == "Circle" {
			circleExists = true
			RemoveIndex(b.shapes, i)

		}
	}
	if circleExists {
		return nil
	} else {
		return errors.New("Circle not exists")
	}	
}

func RemoveIndex(s []Shape, index int) []Shape {
    return append(s[:index], s[index+1:]...)
}
