package main

import "fmt"

func main() {
	//fmt.Println(4 * 3 * 2 * 1)
	n := factorial(4)
	fmt.Println(n)
	m:=loopfact(5)
	fmt.Println(m)
}

func factorial(n int) int {
	if(n==0) {
		return 1
	}
	return n * factorial(n-1)
}

func loopfact(m int)  int {
	total:=1
	for; m>0; m--{
		total *=m
	}
	return total
}