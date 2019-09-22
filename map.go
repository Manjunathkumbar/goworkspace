package main

import "fmt"

func main()  {
	m := map[string]int{
		"james":32,
		"martini":27,
	}
	fmt.Println(m)
	fmt.Println(m["james"])
	fmt.Println(m["msd"])


	//*** imp
	v, ok :=m["msd"]
	fmt.Println(v)
	fmt.Println(ok)

	//if that value exist
	if v, ok :=m["james"]; ok{
		fmt.Println("This is the if print ", v)

	}
	//add to map
	m["todd"]=33
	for k, v :=range m {
		fmt.Println(k, v)
	}
	fmt.Println(m)
	//delete from map
	delete(m, "todd")
	delete(m, "Ian fleming") //if key doesnt exist no error or exception will be thrown
	fmt.Println(m)

}