package svc

// Person is potential Soylent Green
type Person struct {
	Firstname string
	Lastname  string
}

// GetLastFirst returns name in "Last, First" format
func (p Person) GetLastFirst() string {
	return p.Lastname + ", " + p.Firstname
}
