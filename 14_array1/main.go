package main

import "fmt"

func main() {
	var arr1 = [2]int{1, 2}
	var arr2 = [2]int{3, 4}
	// var arr3 = [2]int{}

	var arr3 = [2]int{arr1[0] + arr2[0], arr1[1] + arr2[1]}

	fmt.Println(arr3)

}
