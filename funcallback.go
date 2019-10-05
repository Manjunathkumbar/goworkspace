package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s := sum22(ii...)
	fmt.Println(s)

	s2 := even(sum22, ii...)
	fmt.Println("even numbers sum :", s2)

	s3 := odd(sum22, ii...)
	fmt.Println("odd numbers sum: ", s3)

}

//func(xi ...int) int
func sum22(xi ...int) int {
	//fmt.Printf("%T\n", xi)
	total := 0
	for _, V := range xi {
		total += V
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {

	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odd(f func(xi ...int) int, vi ...int) int {

	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
