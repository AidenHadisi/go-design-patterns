package main

import "fmt"

//Virtual proxy pretends it is there when it is not really. 
//Proxy is not same as decorator. Proxy provides identical interface, decorator provides enhanced interface.

type Image interface {
  Draw()
}

type Bitmap struct {
  filename string
}

func (b *Bitmap) Draw() {
  fmt.Println("Drawing image", b.filename)
}

//a constructor
func NewBitmap(filename string) *Bitmap {
  fmt.Println("Loading image from", filename)
  return &Bitmap{filename: filename}
}

//We are passing an interface.
func DrawImage(image Image) {
  fmt.Println("About to draw the image")
  image.Draw()
  fmt.Println("Done drawing the image")
}

//The problem is, we are always loading the image even when not calling draw.
//A proxy can help us to lazy load image only when needed.

//A proxy that wraps normal bitmap
type LazyBitmap struct {
  filename string
  bitmap *Bitmap
}

func (l *LazyBitmap) Draw() {
  //construct when calling draw
  if l.bitmap == nil {
    l.bitmap = NewBitmap(l.filename)
  }
  l.bitmap.Draw()
}

func NewLazyBitmap(filename string) *LazyBitmap {
  return &LazyBitmap{filename: filename}
}

func main() {
  //bmp := NewBitmap("demo.png")
  bmp := NewLazyBitmap("demo.png")
  DrawImage(bmp)
}
