package decorator

import "fmt"

type Shape interface {
  Render() string
}

type Circle struct {
  Radius float32
}

//implement the interface Shape
func (c *Circle) Render() string {
  return fmt.Sprintf("Circle of radius %f",
    c.Radius)
}

//not part of shpae interface
func (c *Circle) Resize(factor float32) {
  c.Radius *= factor
}

type Square struct {
  Side float32
}

func (s *Square) Render() string {
  return fmt.Sprintf("Square with side %f", s.Side)
}

// possible, but not generic enough
//we can do better and make it more generic because we can have many different types of shapes.
type ColoredSquare struct {
  Square
  Color string
}

type ColoredShape struct {
  Shape Shape
  Color string
}

//Notice colored shape is also a shape
func (c *ColoredShape) Render() string {
  return fmt.Sprintf("%s has the color %s",
    c.Shape.Render(), c.Color)
}

type TransparentShape struct {
  Shape Shape
  Transparency float32
}

func (t *TransparentShape) Render() string {
  return fmt.Sprintf("%s has %f%% transparency",
    t.Shape.Render(), t.Transparency * 100.0)
}

func main() {
  circle := Circle{2}
  fmt.Println(circle.Render())

  redCircle := ColoredShape{&circle, "Red"}
  fmt.Println(redCircle.Render())

  //You can apply decorators to decorators
  rhsCircle := TransparentShape{&redCircle, 0.5}
  fmt.Println(rhsCircle.Render())
}
