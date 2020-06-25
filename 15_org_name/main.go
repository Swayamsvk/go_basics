package main

import "fmt"

func main() {
	name := "Somu"

	switch name {
	case "mama":
		fmt.Println("mama")
	case "Somu":
		fmt.Println("somu")

	default:
		fmt.Println("no name")
	}
}
