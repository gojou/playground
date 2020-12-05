package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gojou/playground/cmd/routing"
	"github.com/gojou/playground/pkg/svc/person"
	"github.com/gojou/playground/pkg/svc/scout"
	"github.com/gorilla/mux"
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

	scouts = append(scouts, scout.Scout{
		Firstname: "Hunter",
		Lastname:  "Skdz",
		Den:       4,
		Rank:      "Senior Webelo",
	})

	for _, p := range persons {
		fmt.Printf("%v\n", p.GetLastFirst())
	}

	for _, s := range scouts {
		fmt.Printf("%v\n", s.GetScoutBasics())
	}
	e := run()
	if e != nil {
		log.Fatal(e)
	}

}
func run() (e error) {
	r := mux.NewRouter()
	routing.Routes(r)

	// Critical to work on AppEngine
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
	return e
}
