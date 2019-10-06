package main

import "fmt"

func main() {
	a := 42

	fmt.Println(a)
	fmt.Println(&a)        //ampersand gives address
	fmt.Printf("%T\n", a)  //int
	fmt.Printf("%T\n", &a) //pointe to an int--> *int

	var b *int = &a
	fmt.Println(b)
	c := b
	fmt.Println(*c) //derefrencing an address //*(astrick) gives u the value stored at an address when u have the address
	fmt.Println(*&a)

	*b = 43
	fmt.Println(a)
}
