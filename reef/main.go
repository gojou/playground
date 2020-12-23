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

	// myBool := true
	// myInt := 2010
	// myFloat64 := 9.15
	// myString := "Hello World!"
	// myComplex128 := (7 + 7i)
	// myIdentity := identity("br549")
	// myIdentityCast := string(myIdentity)
	// myBytesSized := [5]byte{}
	// myBytesUnsized := []byte{}
	// myMap := map[string]int{"a": 1, "b": 2, "c": 3}
	// myPersonEmpty := *new(person)
	// myPerson := person{firstName: "Jayne", lastName: "Doe", age: 34}
	// myPersonPointer := &myPerson
	// myChan := make(chan int)
	// myNil := interface{}(nil)

	var mySlice []interface{}
	mySlice = append(mySlice,
		true,                      // bool
		2010,                      // int
		98.6,                      // float64
		"Hey there!",              // string
		(7 + 7i),                  // complex128
		identity("br549"),         // main.identity
		string(identity("br549")), // main.identity cast to string
		[5]byte{},                 // slice of bytes, sized to 5
		[]byte{},                  // slice of bytes, unsized
		map[string]int{"a": 1, "b": 2, "c": 3},
		*new(person),
		person{firstName: "Jayne", lastName: "Doe", age: 34},
		&person{firstName: "Jayne", lastName: "Doe", age: 34},
		make(chan int),
		nil,
	)

	for _, elem := range mySlice {
		fmt.Printf("type is: %T -- value is: %v\n", elem, elem)
	}
}
