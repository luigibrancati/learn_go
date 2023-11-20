package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "Alex", lastName: "Andersen"}
	fmt.Println(alex)
	var alex2 person
	fmt.Println(alex2)
	fmt.Printf("%+v\n", alex2)
	alex2.firstName = "Alex"
	alex2.lastName = "Andersen"
	fmt.Printf("%+v\n", alex2)
}
