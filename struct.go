package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

//embedded struct
type secretAgent struct {
	person
	licence bool
}

//method of type secret agent
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)

}

func main() {

	p1 := person{
		first: "james",
		last:  "kennedy",
		age:   32,
	}
	p2 := person{
		first: "miss",
		last:  "moneypenny",
		age:   27,
	}

	s1 := secretAgent{
		person:  person{first: "james", last: "bond", age: 45},
		licence: true,
	}
	fmt.Print(p1)
	fmt.Println(p2)

	fmt.Print(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
	fmt.Println(s1)
	fmt.Println(s1.first, s1.last, s1.age)
	s1.speak()

	//anonymous struct

}
