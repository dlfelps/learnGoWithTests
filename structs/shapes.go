package structs

import "math"


type Rectangle struct {
  Width  float64
  Height float64
}
type Circle struct {
  Radius float64
}

func (r Rectangle) Perimeter() float64 {
  return 2 * (r.Width + r.Height)
}

type Shape interface {
  Area() float64
}

func (r Rectangle) Area() float64 {
  return r.Width * r.Height
}

func (c Circle) Area() float64 {
  return math.Pi * c.Radius * c.Radius
}