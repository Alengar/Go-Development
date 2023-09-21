package people

import "fmt"

type Senior struct {
	Colleague
}

func NewSenior(name string) *Senior {
	return &Senior{Colleague{Person{Name: name}}}
}

func (s *Senior) Act() {
	s.Greet()
	fmt.Printf("%s is behaving respectfully with seniors.\n", s.Name)
}
