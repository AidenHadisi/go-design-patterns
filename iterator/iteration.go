package main

import "fmt"

type Person struct {
  FirstName, MiddleName, LastName string
}

func (p *Person) Names() []string {
  return []string{p.FirstName, p.MiddleName, p.LastName}
}

func main() {
  p := Person{"Alexander", "Graham", "Bell"}
  for _, name := range p.Names() {
    fmt.Println(name)
  }
}

// or use a channel

// or use a sperate struct that keeps track of current element, and composed of person. It will have Next() and Value() methods.
