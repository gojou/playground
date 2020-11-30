package person

// GetLastFirst returns name in "Last, First" format
func (p Person) GetLastFirst() string {
	return p.Lastname + ", " + p.Firstname
}
