package structural_bridge

import "fmt"

//We might have RasterCircle, RasterSquare, Vector Circle, etc.
//That is bad, we can use bridge method

type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRenderer struct {
}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius", radius)
}

type RasterRenderer struct {
	Dpi int
}

func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels for circle of radius", radius)
}

type Circle struct {
	renderer Renderer
	radius   float32
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

func main() {
	raster := RasterRenderer{}
	vector := VectorRenderer{}
	circle := NewCircle(&vector, 5)
	circle.Draw()

	circle2 := NewCircle(&raster, 4)
	circle2.Draw()
}
