package main

import "fmt"

//func (r receiver) identifier(parameter(s)) (return(s)) { ....logic code }

func main()  {
	//basic function
	foo()
	//func with parameter
	bar("james")
	//func with return
	s1 := woo("moneypenny")
	fmt.Println(s1)
	//func with multiple return
	x, y := mouse("ian", "fleming")
	fmt.Println(x, y)

	fmt.Println("functions")

	foo1(1,2,3,4,5,6,7)
}
func foo(){
	fmt.Println(" hello from fooo")
}

//everything in go is PASS BY VALUE
func bar(s string){
	fmt.Println("hello ", s)
}

func woo(s string)  string {
	return fmt.Sprint("hello from woo :", s)
}
func mouse(fn string, ln string) (string, bool) {
	a:= fmt.Sprint(fn , ln, "says hello")
	b := true
	return a,b
}

//variadic parameter
func  foo1(x... int)  {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}