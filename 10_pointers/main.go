package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Println(*b)

	//change value of pointer
	*b = 10
	fmt.Println(a)
}
