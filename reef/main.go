package main

import "fmt"

type identity string

type guy struct {
	name string
}

func main() {

	var foo interface{}
	println(foo)

	foo = 123
	println(foo)

	foo = "So sorry"
	println(foo)

	// This works. goo is a copy of foo, which was initialized with an empty interface.
	// goo inherits the structure of foo as well as the value.
	goo := foo
	println(goo)
	goo = 12
	println(goo)

	var hoo string
	hoo = "Hoo boy"
	println(hoo)

	// hoo = 13    ERROR -- hoo initialized with concrete type string, must be string
	// println(hoo)

	per := person{
		firstName: "Mark",
		lastName:  "Poling",
		age:       57,
	}
	fmt.Printf("Hello %v\n", per)

	// println(true)
	// println(2010)
	// println(9.15)
	// println(7 + 7i)
	// println("Hello World!")
	// println(identity("19520925"))
	// println([5]byte{})
	// println([]byte{})
	// println(map[string]int{})
	// println(guy{name: "John Doe"})
	// println(&guy{name: "John Doe"})
	// println(make(chan int))
	// println(*new(guy))
	// println(nil)
}

func println(x interface{}) {
	fmt.Printf("type is: %T -- value is: %v\n", x, x)
}
