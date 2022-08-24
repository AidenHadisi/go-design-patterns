package prototype

import "fmt"

type Address struct {
  StreetAddress, City, Country string
}

type Person struct {
  Name string
  Address *Address
}

func main() {
  john := Person{"John",
    &Address{"123 London Rd", "London", "UK"}}

  //jane := john

  // this is a shallow copy
  //jane.Name = "Jane" // ok

  //jane.Address.StreetAddress = "321 Baker St"

  //fmt.Println(john.Name, john.Address)
  //fmt.Println(jane.Name, jane. Address)

  // the above code won't work because we need to copy the actual value not the pointer. 
  // so we need to do it like this
  jane := john
  jane.Address = &Address{
    john.Address.StreetAddress,
    john.Address.City,
    john.Address.Country  }

  jane.Name = "Jane" // ok

  jane.Address.StreetAddress = "321 Baker St"

  fmt.Println(john.Name, john.Address)
  fmt.Println(jane.Name, jane. Address)
}
