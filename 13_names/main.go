package main

import "fmt"

func callName(name string) string {
	return "Hello " + name
}

func main() {
	fmt.Printf(callName("Swayam\n"))
}
