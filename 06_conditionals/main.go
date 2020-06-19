package main

import "fmt"

func main() {

	// x := 5
	// y := 10

	// if x < y {
	// 	fmt.Printf("%d is less than %d\n", x, y)
	// }

	color := "red"

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is not red or blue")
	}

}
