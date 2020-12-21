package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	i := inches(2.0)
	m := mms(5)
	fmt.Println(i)
	fmt.Println(m)
	fmt.Printf("%v is equal to %v\n", i, i.convert())
	fmt.Printf("%v is equal to %v\n", m, m.convert())

}
