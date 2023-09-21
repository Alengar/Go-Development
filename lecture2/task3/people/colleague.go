package people

import "fmt"

type Colleague struct {
	Person
}

func NewColleague(name string) *Colleague {
	return &Colleague{Person{Name: name}}
}

func (c *Colleague) Act() {
	c.Greet()
	fmt.Printf("%s is behaving professionally with colleagues.\n", c.Name)
}
