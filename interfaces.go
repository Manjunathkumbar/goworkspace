package main

import "fmt"

type person1 struct {
	first string
	last  string
	age   int
}

//embedded struct
type secretAgent1 struct {
	person1
	licence bool
}

func (p person1) speak() {
	fmt.Println("I am", p.first, p.last)
}

func (s secretAgent1) speak() {
	fmt.Println("I am", s.first, s.last)
}

// a value can be more than one type
type human interface {
	speak() // if anything has this method its of type human
}

func bar1(h human) {
	switch h.(type) {
	case person1:
		fmt.Println("I passsed into bar", h.(person1).first) //asserting
	case secretAgent1:
		fmt.Println("I passsed into bar", h.(secretAgent1).first)
	}
	fmt.Println("I passsed into bar", h)
}

func main() {

	s1 := secretAgent1{
		person1: person1{first: "james", last: "bond", age: 45},
		licence: true,
	}

	s2 := secretAgent1{
		person1: person1{first: "money", last: "penny", age: 30},
		licence: true,
	}

	p1 := person1{first: "Dr", last: "yes", age: 27}

	fmt.Println(s1)
	fmt.Println(s1.first, s1.last, s1.age)
	s1.speak()
	bar1(s1)
	bar1(s2)
	bar1(p1)
}
