package main

import "fmt"

func main() {
	fmt.Printf("%v\n", "Hello World!")
	var persons []Person

	persons = append(persons, Person{"Mark", "Poling"})
	persons = append(persons, Person{"Liz", "Poling-Hiraldo"})
	persons = append(persons, Person{"Aden", "Poling"})
	persons = append(persons, Person{"Rhi", "Poling"})

	for _, p := range persons {
		fmt.Printf("%v\n", p.GetLastFirst())
	}
}
