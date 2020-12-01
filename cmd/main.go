package main

import (
	"fmt"

	"github.com/gojou/playground/pkg/svc/person"
)

func main() {
	fmt.Printf("%v\n", "Hello guys!")
	var persons []person.Person

	persons = append(persons, person.Person{Firstname: "Mark", Lastname: "Poling"})
	persons = append(persons, person.Person{Firstname: "Liz", Lastname: "Poling-Hiraldo"})
	persons = append(persons, person.Person{Firstname: "Aden", Lastname: "Poling"})
	persons = append(persons, person.Person{Firstname: "Rhi", Lastname: "Poling"})

	for _, p := range persons {
		fmt.Printf("%v\n", p.GetLastFirst())
	}
}
