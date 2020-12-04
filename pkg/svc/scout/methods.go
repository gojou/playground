package scout

import "strconv"

// GetScoutBasics returns basic scout info
func (s Scout) GetScoutBasics() string {
	return s.Lastname + ", " + s.Firstname + " - Den: " + strconv.Itoa(s.Den) + " - Rank: " + s.Rank

}
