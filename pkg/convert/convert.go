package convert

import "strconv"

type StrTo string

func (s StrTo) String() string {
	return string(s)
}
func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}
func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}
