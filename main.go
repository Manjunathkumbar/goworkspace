package main

import "fmt"

var y = 43
var z = "shaken not stirred"

var a int

type hotdog int

var b hotdog

func main() {

	fmt.Println("Lets Go")

	for i := 0; i <= 20; i++ {
		fmt.Println(i)
		fmt.Println(y)
		fmt.Printf("%T\n", y)
		//z=43
		//fmt.Printf("%T\n", z)
	}

	fmt.Println("#################")
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("a is of type:"+"%T\n", a)
	fmt.Println(b)
	fmt.Printf("b is of type: "+"%T\n", b)
	//conversion from type to int---> its not casting in go
	a = int(b)
	fmt.Println(a)
	fmt.Printf("b is of converted to int: "+"%T\n", a)

}
