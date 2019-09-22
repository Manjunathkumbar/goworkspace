package main

import "fmt"

func main() {
	//arrays--->Building blocks for slices---so we wont use arrays in Go only slices
	var A [5]int
	fmt.Println(A)
	A[2] = 42
	fmt.Println(A)
	fmt.Println(len(A))
	//Slices
	// X := []type{values} //composite literal //defining slice
	//Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data.
	// Except for items with explicit dimension such as transformation matrices,
	// most array programming in Go is done with slices rather than simple arrays.
	x := []int{1, 3, 5, 7, 9}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[0])
	// slice allows to group values of same type

	for i, v := range x {
		fmt.Println(i, v)
	}

	//slicing the slices
	fmt.Println(x[:])
	fmt.Println(x[1:3])

	// without using range
	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}

	//Appending to slice slice

	y := []int{1, 2, 3, 4, 5}
	y = append(y, 77, 88, 99)
	fmt.Println(y)

	z := []int{11, 22, 33, 444, 55}
	z = append(z, y...)
	fmt.Println(z) //[11 22 33 444 55 1 2 3 4 5 77 88 99]

	//delete from slice
	z = append(z[:2], z[4:]...) //three dots ... is very imp
	fmt.Println(z)

	//slice make---built in function make----make([]int, 10, 100)
	m := make([]int, 10, 12)
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))
	m[0] = 24
	m[9] = 999
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))

	m = append(m, 3423)

	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))

	m = append(m, 3423)
	m = append(m, 3423) //it will make double the size in run time 24

	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))

	// multi dimensional slice
	jb := []string{"msd", "james", "bond", "martini"}
	fmt.Println(jb)

	mp := []string{"aegon", "james", "anna", "kennedy"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)

}
