package people

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Greet() {
	fmt.Printf("%s says hello!\n", p.Name)
}

type Behavior interface {
	Act()
}
