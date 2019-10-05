package main

import "fmt"

func main() {
	defer fooo() //used for closing the file or resource
	barrr()      //this will execute first
}

func fooo() {
	fmt.Println("foooo")
}

func barrr() {
	fmt.Println("bar")

}
