package main

import (
	"fmt"
	"strconv"
)

//Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

//Greeting
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " " + strconv.Itoa(p.age)
}

func main() {

	//Init person using struct
	//person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	//Alternative
	person1 := Person{"Samantha", "Smith", "Boston", "f", 25}
	// fmt.Println(person1)
	fmt.Println(person1.greet())

}
