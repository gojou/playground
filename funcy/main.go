package main

import "fmt"

type inches float64
type mms float64

func (i inches) String() string {
	return fmt.Sprintf("%.2f\"", i)
}
func (i inches) convert() mms {
	return mms(float64(i) * 2.54)
}

func (m mms) String() string {
	return fmt.Sprintf("%.2fmm", m)
}

func (m mms) convert() inches {
	return inches(float64(m) / 2.54)
}

func main() {
	fmt.Println("Hello world!")
	i := inches(2.0)
	m := mms(5)
	fmt.Println(i)
	fmt.Println(m)
	fmt.Printf("%v is equal to %v\n", i, i.convert())
	fmt.Printf("%v is equal to %v\n", m, m.convert())

}
