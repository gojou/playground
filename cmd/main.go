package main

import (
	"fmt"

	"github.com/gojou/playground/pkg/svc/person"
	"github.com/gojou/playground/pkg/svc/scout"
)

func main() {
	fmt.Printf("%v\n", "Hello guys!")
	var persons []person.Person

	persons = append(persons, person.Person{Firstname: "Mark", Lastname: "Poling"})
	persons = append(persons, person.Person{Firstname: "Liz", Lastname: "Poling-Hiraldo"})
	persons = append(persons, person.Person{Firstname: "Aden", Lastname: "Poling"})
	persons = append(persons, person.Person{Firstname: "Rhi", Lastname: "Poling"})

	var scouts []scout.Scout
	scouts = append(scouts, scout.Scout{
		Firstname: "Aden",
		Lastname:  "Poling",
		Den:       4,
		Rank:      "Senior Webelo",
	})

	for _, p := range persons {
		fmt.Printf("%v\n", p.GetLastFirst())
	}

	for _, s := range scouts {
		fmt.Printf("%v\n", s.GetScoutBasics())
	}

}
