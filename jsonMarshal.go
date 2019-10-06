package main

import (
	"encoding/json"
	"fmt"
)

type personp struct {
	First string //try with lower case first will get [{},{}]
	Last  string
	Age   int
}

func main() {
	p1 := personp{
		First: "james",
		Last:  "bond",
		Age:   35,
	}
	p2 := personp{
		First: "money",
		Last:  "penny",
		Age:   35,
	}
	people := []personp{p1, p2}
	fmt.Println(people) //go code

	bs, err := json.Marshal(people) //byte slice
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs)) //convert byte slice to int--->json
}
