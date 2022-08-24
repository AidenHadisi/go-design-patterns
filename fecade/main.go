package facade

import "fmt"

//Let's say we have Buffer and Viewport which are complicated classes
type Buffer struct {
	width, height int
	buffer        []rune
}

func NewBuffer(width, height int) *Buffer {
	return &Buffer{width, height,
		make([]rune, width*height)}
}

func (b *Buffer) At(index int) rune {
	return b.buffer[index]
}

type Viewport struct {
	buffer *Buffer
	offset int
}

func NewViewport(buffer *Buffer) *Viewport {
	return &Viewport{buffer: buffer}
}

func (v *Viewport) GetCharacterAt(index int) rune {
	return v.buffer.At(v.offset + index)
}

// a facade over buffers and viewports
type Console struct {
	buffers   []*Buffer
	viewports []*Viewport
	offset    int
}

//creates a default scenario
//we only have a signle buffer and viewport in facade here
func NewConsole() *Console {
	b := NewBuffer(10, 10)
	v := NewViewport(b)
	return &Console{[]*Buffer{b}, []*Viewport{v}, 0}
}

//high level func. Since we know fecade only provides one of each, this is easier
func (c *Console) GetCharacterAt(index int) rune {
	return c.viewports[0].GetCharacterAt(index)
}

func main() {
	//instead of working with low level stuff like buffers and viewports, we use fecade
	//since a single buffer and viewport is common we don't need to work with low level stuff
	//if people want more, then can use types directly
	c := NewConsole()
	u := c.GetCharacterAt(1)

	fmt.Println(u)
}
