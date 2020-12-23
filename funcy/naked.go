package main

import "fmt"

type funcy func()

func run(f funcy) {
	f()
}

func hello1() { fmt.Println("Hello friends!") }

func hello2() { fmt.Println("Hola amigos!") }

func hello3() { fmt.Println("Salut mes amis!") }

func greeting() {
	var greetings []func()
	greetings = append(greetings, hello1)
	greetings = append(greetings, hello2)
	greetings = append(greetings, hello3)

	for i, f := range greetings {
		fmt.Printf("%v - %T: ", i, f)
		run(f)
	}

}
