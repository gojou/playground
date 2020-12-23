package main

import "strconv"

type person struct {
	lastName  string
	firstName string
	age       int
}

func (p person) String() string {
	return p.firstName + " " + p.lastName + ": " + strconv.Itoa(p.age)
}
