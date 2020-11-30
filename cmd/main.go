package main

import (
	"fmt"

	person "github.com/gojou/playground/pkg/svc/person"
)

func main() {
	fmt.Printf("%v\n", "Hello World!")
	var persons []person.Person

	persons = append(persons, person.Person{"Mark", "Poling"})
	persons = append(persons, person.Person{"Liz", "Poling-Hiraldo"})
	persons = append(persons, person.Person{"Aden", "Poling"})
	persons = append(persons, person.Person{"Rhi", "Poling"})

	for _, p := range persons {
		fmt.Printf("%v\n", p.GetLastFirst())
	}
}
