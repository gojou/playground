// This is modified from the lesson found at:
//		https://www.youtube.com/watch?v=oiX7fAmOYX0
// Let's gets started with Go Reflection, https://golang.org/pkg/reflect/
// Source code: https://github.com/striversity/go-on-the-run

package main

import "fmt"

type identity string

func main() {

	var foo interface{}
	fmt.Println(foo)

	foo = 123
	fmt.Println(foo)

	foo = "So sorry"
	fmt.Println(foo)

	// This works. goo is a copy of foo, which was initialized with an empty interface.
	// goo inherits the structure of foo as well as the value.
	goo := foo
	fmt.Println(goo)
	goo = 12
	fmt.Println(goo)

	var hoo string
	hoo = "Hoo boy"
	fmt.Println(hoo)

	// hoo = 13 // COMPILE ERROR -- hoo initialized with concrete type string, must be string
	// fmt.Println(hoo)

	for _, elem := range getCases() {
		fmt.Printf("type is: %T -- value is: %v\n", elem, elem)
	}
}
