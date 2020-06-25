package main

import "fmt"

type Name struct {
	name  string
	email string
}

func (n Name) details() string {
	return "Hello my name is " + n.name + " email is " + n.email

}

func main() {
	name1 := Name{"Swayam", "swayamsamyak@gmail.com"}
	fmt.Println(name1.details())
}
