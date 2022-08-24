package prototype

import "fmt"

type Address struct {
  StreetAddress, City, Country string
}

//you can create a deep copy method on your structs that clones them.
func (a *Address) DeepCopy() *Address {
  return &Address{
    a.StreetAddress,
    a.City,
    a.Country }
}

type Person struct {
  Name string
  Address *Address
  Friends []string
}

//Same on here
func (p *Person) DeepCopy() *Person {
  q := *p // copies Name
  q.Address = p.Address.DeepCopy()
  
  //deep copy the slice using copy method
  copy(q.Friends, p.Friends)
  return &q
}

func main() {
  john := Person{"John",
    &Address{"123 London Rd", "London", "UK"},
    []string{"Chris", "Matt"}}

  jane := john.DeepCopy()
  jane.Name = "Jane"
  jane.Address.StreetAddress = "321 Baker St"
  jane.Friends = append(jane.Friends, "Angela")

  fmt.Println(john, john.Address)
  fmt.Println(jane, jane.Address)
}


