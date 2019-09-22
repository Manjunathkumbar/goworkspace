package main

import "fmt"

func main() {

	p1 := struct {
		first string
		last  string
	}{
		first: "james",
		last:  "bond",
	}
	fmt.Println(p1)

}
