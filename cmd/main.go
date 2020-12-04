package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gojou/playground/pkg/svc/person"
	"github.com/gojou/playground/pkg/svc/scout"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

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

}
