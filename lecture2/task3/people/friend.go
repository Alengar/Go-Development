package people

import "fmt"

type Friend struct {
	Person
}

func NewFriend(name string) *Friend {
	return &Friend{Person{Name: name}}
}

func (f *Friend) Act() {
	f.Greet()
	fmt.Printf("%s is behaving casually with friends.\n", f.Name)
}
