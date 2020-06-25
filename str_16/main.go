package main

import "fmt"

func main() {
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	//Loop through ids
	for i, id := range emails {
		fmt.Println(i + " element " + id)
	}
}
