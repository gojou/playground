package main

func getCases() []interface{} {
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
	return mySlice
}
