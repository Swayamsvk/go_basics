package main

import "fmt"

func greeting(name string) string {
	return "\nHello" + name
}
func main() {
	fmt.Printf(greeting("Swayam "))
}
