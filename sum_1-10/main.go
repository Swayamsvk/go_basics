package main

import "fmt"

func main() {

	var sum int
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {

			sum += i
		}
	}

	fmt.Println(sum)

}
