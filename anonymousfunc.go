package main

import "fmt"

func main() {
	foo1()

	func() {
		fmt.Println("anonymous func")
	}()

	func(x int) {
		fmt.Println("anonymous func", x)
	}(27)

	//function expression, like defining variable
	f := func() {
		fmt.Println("my first func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("my second func expression", x)
	}
	g(45)

}
func foo1() {
	fmt.Println("fooo1")

}
