package main

import "fmt"

func main() {
	//Define map
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	//Assign kv
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"

	fmt.Println(emails["Bob"])

	//delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
